package cargo

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"time"

	"github.com/bhojpur/scf/pkg/logistic/location"
	"github.com/bhojpur/scf/pkg/logistic/voyage"
)

// Delivery is the actual transportation of the cargo, as opposed to the
// customer requirement (RouteSpecification) and the plan (Itinerary).
type Delivery struct {
	Itinerary               Itinerary
	RouteSpecification      RouteSpecification
	RoutingStatus           RoutingStatus
	TransportStatus         TransportStatus
	NextExpectedActivity    HandlingActivity
	LastEvent               HandlingEvent
	LastKnownLocation       location.UNLocode
	CurrentVoyage           voyage.Number
	ETA                     time.Time
	IsMisdirected           bool
	IsUnloadedAtDestination bool
}

// UpdateOnRouting creates a new delivery snapshot to reflect changes in
// routing, i.e. when the route specification or the itinerary has changed but
// no additional handling of the cargo has been performed.
func (d Delivery) UpdateOnRouting(rs RouteSpecification, itinerary Itinerary) Delivery {
	return newDelivery(d.LastEvent, itinerary, rs)
}

// IsOnTrack checks if the delivery is on track.
func (d Delivery) IsOnTrack() bool {
	return d.RoutingStatus == Routed && !d.IsMisdirected
}

// DeriveDeliveryFrom creates a new delivery snapshot based on the complete
// handling history of a cargo, as well as its route specification and
// itinerary.
func DeriveDeliveryFrom(rs RouteSpecification, itinerary Itinerary, history HandlingHistory) Delivery {
	lastEvent, _ := history.MostRecentlyCompletedEvent()
	return newDelivery(lastEvent, itinerary, rs)
}

// newDelivery creates a up-to-date delivery based on an handling event,
// itinerary and a route specification.
func newDelivery(lastEvent HandlingEvent, itinerary Itinerary, rs RouteSpecification) Delivery {
	var (
		routingStatus           = calculateRoutingStatus(itinerary, rs)
		transportStatus         = calculateTransportStatus(lastEvent)
		lastKnownLocation       = calculateLastKnownLocation(lastEvent)
		isMisdirected           = calculateMisdirectedStatus(lastEvent, itinerary)
		isUnloadedAtDestination = calculateUnloadedAtDestination(lastEvent, rs)
		currentVoyage           = calculateCurrentVoyage(transportStatus, lastEvent)
	)

	d := Delivery{
		LastEvent:               lastEvent,
		Itinerary:               itinerary,
		RouteSpecification:      rs,
		RoutingStatus:           routingStatus,
		TransportStatus:         transportStatus,
		LastKnownLocation:       lastKnownLocation,
		IsMisdirected:           isMisdirected,
		IsUnloadedAtDestination: isUnloadedAtDestination,
		CurrentVoyage:           currentVoyage,
	}

	d.NextExpectedActivity = calculateNextExpectedActivity(d)
	d.ETA = calculateETA(d)

	return d
}

// Below are internal functions used when creating a new delivery.

func calculateRoutingStatus(itinerary Itinerary, rs RouteSpecification) RoutingStatus {
	if itinerary.Legs == nil {
		return NotRouted
	}

	if rs.IsSatisfiedBy(itinerary) {
		return Routed
	}

	return Misrouted
}

func calculateMisdirectedStatus(event HandlingEvent, itinerary Itinerary) bool {
	if event.Activity.Type == NotHandled {
		return false
	}

	return !itinerary.IsExpected(event)
}

func calculateUnloadedAtDestination(event HandlingEvent, rs RouteSpecification) bool {
	if event.Activity.Type == NotHandled {
		return false
	}

	return event.Activity.Type == Unload && rs.Destination == event.Activity.Location
}

func calculateTransportStatus(event HandlingEvent) TransportStatus {
	switch event.Activity.Type {
	case NotHandled:
		return NotReceived
	case Load:
		return OnboardCarrier
	case Unload:
		return InPort
	case Receive:
		return InPort
	case Customs:
		return InPort
	case Claim:
		return Claimed
	}
	return Unknown
}

func calculateLastKnownLocation(event HandlingEvent) location.UNLocode {
	return event.Activity.Location
}

func calculateNextExpectedActivity(d Delivery) HandlingActivity {
	if !d.IsOnTrack() {
		return HandlingActivity{}
	}

	switch d.LastEvent.Activity.Type {
	case NotHandled:
		return HandlingActivity{Type: Receive, Location: d.RouteSpecification.Origin}
	case Receive:
		l := d.Itinerary.Legs[0]
		return HandlingActivity{Type: Load, Location: l.LoadLocation, VoyageNumber: l.VoyageNumber}
	case Load:
		for _, l := range d.Itinerary.Legs {
			if l.LoadLocation == d.LastEvent.Activity.Location {
				return HandlingActivity{Type: Unload, Location: l.UnloadLocation, VoyageNumber: l.VoyageNumber}
			}
		}
	case Unload:
		for i, l := range d.Itinerary.Legs {
			if l.UnloadLocation == d.LastEvent.Activity.Location {
				if i < len(d.Itinerary.Legs)-1 {
					return HandlingActivity{Type: Load, Location: d.Itinerary.Legs[i+1].LoadLocation, VoyageNumber: d.Itinerary.Legs[i+1].VoyageNumber}
				}

				return HandlingActivity{Type: Claim, Location: l.UnloadLocation}
			}
		}
	}

	return HandlingActivity{}
}

func calculateCurrentVoyage(transportStatus TransportStatus, event HandlingEvent) voyage.Number {
	if transportStatus == OnboardCarrier && event.Activity.Type != NotHandled {
		return event.Activity.VoyageNumber
	}

	return voyage.Number("")
}

func calculateETA(d Delivery) time.Time {
	if !d.IsOnTrack() {
		return time.Time{}
	}

	return d.Itinerary.FinalArrivalTime()
}