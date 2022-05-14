package voyage

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

// It provides the Voyage aggregate.

import (
	"errors"
	"time"

	"github.com/bhojpur/scf/pkg/logistic/location"
)

// Number uniquely identifies a particular Voyage.
type Number string

// Voyage is a uniquely identifiable series of carrier movements.
type Voyage struct {
	Number   Number
	Schedule Schedule
}

// New creates a voyage with a voyage number and a provided schedule.
func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}

// Schedule describes a voyage schedule.
type Schedule struct {
	CarrierMovements []CarrierMovement
}

// CarrierMovement is a vessel voyage from one location to another.
type CarrierMovement struct {
	DepartureLocation location.UNLocode
	ArrivalLocation   location.UNLocode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

// ErrUnknown is used when a voyage could not be found.
var ErrUnknown = errors.New("unknown voyage")

// Repository provides access a voyage store.
type Repository interface {
	Find(Number) (*Voyage, error)
}