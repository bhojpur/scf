package inspection

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

// It provides means to inspect cargos.

import (
	"github.com/bhojpur/scf/pkg/logistic/cargo"
)

// EventHandler provides means of subscribing to inspection events.
type EventHandler interface {
	CargoWasMisdirected(*cargo.Cargo)
	CargoHasArrived(*cargo.Cargo)
}

// Service provides cargo inspection operations.
type Service interface {
	// InspectCargo inspects cargo and send relevant notifications to
	// interested parties, for example if a cargo has been misdirected, or
	// unloaded at the final destination.
	InspectCargo(id cargo.TrackingID)
}

type service struct {
	cargos  cargo.Repository
	events  cargo.HandlingEventRepository
	handler EventHandler
}

// TODO: Should be transactional
func (s *service) InspectCargo(id cargo.TrackingID) {
	c, err := s.cargos.Find(id)
	if err != nil {
		return
	}

	h := s.events.QueryHandlingHistory(id)

	c.DeriveDeliveryProgress(h)

	if c.Delivery.IsMisdirected {
		s.handler.CargoWasMisdirected(c)
	}

	if c.Delivery.IsUnloadedAtDestination {
		s.handler.CargoHasArrived(c)
	}

	s.cargos.Store(c)
}

// NewService creates a inspection service with necessary dependencies.
func NewService(cargos cargo.Repository, events cargo.HandlingEventRepository, handler EventHandler) Service {
	return &service{cargos, events, handler}
}