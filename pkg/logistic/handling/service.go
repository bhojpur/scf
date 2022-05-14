package handling

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

// It provides the use-case for registering incidents. Used by views facing the
// people handling the cargo along its route.

import (
	"errors"
	"time"

	"github.com/bhojpur/scf/pkg/logistic/cargo"
	"github.com/bhojpur/scf/pkg/logistic/inspection"
	"github.com/bhojpur/scf/pkg/logistic/location"
	"github.com/bhojpur/scf/pkg/logistic/voyage"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// EventHandler provides a means of subscribing to registered handling events.
type EventHandler interface {
	CargoWasHandled(cargo.HandlingEvent)
}

// Service provides handling operations.
type Service interface {
	// RegisterHandlingEvent registers a handling event in the system, and
	// notifies interested parties that a cargo has been handled.
	RegisterHandlingEvent(completed time.Time, id cargo.TrackingID, voyageNumber voyage.Number,
		unLocode location.UNLocode, eventType cargo.HandlingEventType) error
}

type service struct {
	handlingEventRepository cargo.HandlingEventRepository
	handlingEventFactory    cargo.HandlingEventFactory
	handlingEventHandler    EventHandler
}

func (s *service) RegisterHandlingEvent(completed time.Time, id cargo.TrackingID, voyageNumber voyage.Number,
	loc location.UNLocode, eventType cargo.HandlingEventType) error {
	if completed.IsZero() || id == "" || loc == "" || eventType == cargo.NotHandled {
		return ErrInvalidArgument
	}

	e, err := s.handlingEventFactory.CreateHandlingEvent(time.Now(), completed, id, voyageNumber, loc, eventType)
	if err != nil {
		return err
	}

	s.handlingEventRepository.Store(e)
	s.handlingEventHandler.CargoWasHandled(e)

	return nil
}

// NewService creates a handling event service with necessary dependencies.
func NewService(r cargo.HandlingEventRepository, f cargo.HandlingEventFactory, h EventHandler) Service {
	return &service{
		handlingEventRepository: r,
		handlingEventFactory:    f,
		handlingEventHandler:    h,
	}
}

type handlingEventHandler struct {
	InspectionService inspection.Service
}

func (h *handlingEventHandler) CargoWasHandled(event cargo.HandlingEvent) {
	h.InspectionService.InspectCargo(event.TrackingID)
}

// NewEventHandler returns a new instance of a EventHandler.
func NewEventHandler(s inspection.Service) EventHandler {
	return &handlingEventHandler{
		InspectionService: s,
	}
}