package booking

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

	"github.com/go-kit/kit/metrics"

	"github.com/bhojpur/scf/pkg/logistic/cargo"
	"github.com/bhojpur/scf/pkg/logistic/location"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) BookNewCargo(origin, destination location.UNLocode, deadline time.Time) (cargo.TrackingID, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.BookNewCargo(origin, destination, deadline)
}

func (s *instrumentingService) LoadCargo(id cargo.TrackingID) (c Cargo, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.LoadCargo(id)
}

func (s *instrumentingService) RequestPossibleRoutesForCargo(id cargo.TrackingID) []cargo.Itinerary {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RequestPossibleRoutesForCargo(id)
}

func (s *instrumentingService) AssignCargoToRoute(id cargo.TrackingID, itinerary cargo.Itinerary) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.AssignCargoToRoute(id, itinerary)
}

func (s *instrumentingService) ChangeDestination(id cargo.TrackingID, l location.UNLocode) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ChangeDestination(id, l)
}

func (s *instrumentingService) Cargos() []Cargo {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_cargos").Add(1)
		s.requestLatency.With("method", "list_cargos").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Cargos()
}

func (s *instrumentingService) Locations() []Location {
	defer func(begin time.Time) {
		s.requestCount.With("method", "list_locations").Add(1)
		s.requestLatency.With("method", "list_locations").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Locations()
}