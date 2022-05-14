package inmem

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

// It provides in-memory implementations of all the domain repositories.

import (
	"sync"

	"github.com/bhojpur/scf/pkg/logistic/cargo"
	"github.com/bhojpur/scf/pkg/logistic/location"
	"github.com/bhojpur/scf/pkg/logistic/voyage"
)

type cargoRepository struct {
	mtx    sync.RWMutex
	cargos map[cargo.TrackingID]*cargo.Cargo
}

func (r *cargoRepository) Store(c *cargo.Cargo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.cargos[c.TrackingID] = c
	return nil
}

func (r *cargoRepository) Find(id cargo.TrackingID) (*cargo.Cargo, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.cargos[id]; ok {
		return val, nil
	}
	return nil, cargo.ErrUnknown
}

func (r *cargoRepository) FindAll() []*cargo.Cargo {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*cargo.Cargo, 0, len(r.cargos))
	for _, val := range r.cargos {
		c = append(c, val)
	}
	return c
}

// NewCargoRepository returns a new instance of a in-memory cargo repository.
func NewCargoRepository() cargo.Repository {
	return &cargoRepository{
		cargos: make(map[cargo.TrackingID]*cargo.Cargo),
	}
}

type locationRepository struct {
	locations map[location.UNLocode]*location.Location
}

func (r *locationRepository) Find(locode location.UNLocode) (*location.Location, error) {
	if l, ok := r.locations[locode]; ok {
		return l, nil
	}
	return nil, location.ErrUnknown
}

func (r *locationRepository) FindAll() []*location.Location {
	l := make([]*location.Location, 0, len(r.locations))
	for _, val := range r.locations {
		l = append(l, val)
	}
	return l
}

// NewLocationRepository returns a new instance of a in-memory location repository.
func NewLocationRepository() location.Repository {
	r := &locationRepository{
		locations: make(map[location.UNLocode]*location.Location),
	}

	r.locations[location.SESTO] = location.Stockholm
	r.locations[location.AUMEL] = location.Melbourne
	r.locations[location.CNHKG] = location.Hongkong
	r.locations[location.JNTKO] = location.Tokyo
	r.locations[location.NLRTM] = location.Rotterdam
	r.locations[location.DEHAM] = location.Hamburg

	return r
}

type voyageRepository struct {
	voyages map[voyage.Number]*voyage.Voyage
}

func (r *voyageRepository) Find(voyageNumber voyage.Number) (*voyage.Voyage, error) {
	if v, ok := r.voyages[voyageNumber]; ok {
		return v, nil
	}

	return nil, voyage.ErrUnknown
}

// NewVoyageRepository returns a new instance of a in-memory voyage repository.
func NewVoyageRepository() voyage.Repository {
	r := &voyageRepository{
		voyages: make(map[voyage.Number]*voyage.Voyage),
	}

	r.voyages[voyage.V100.Number] = voyage.V100
	r.voyages[voyage.V300.Number] = voyage.V300
	r.voyages[voyage.V400.Number] = voyage.V400

	r.voyages[voyage.V0100S.Number] = voyage.V0100S
	r.voyages[voyage.V0200T.Number] = voyage.V0200T
	r.voyages[voyage.V0300A.Number] = voyage.V0300A
	r.voyages[voyage.V0301S.Number] = voyage.V0301S
	r.voyages[voyage.V0400S.Number] = voyage.V0400S

	return r
}

type handlingEventRepository struct {
	mtx    sync.RWMutex
	events map[cargo.TrackingID][]cargo.HandlingEvent
}

func (r *handlingEventRepository) Store(e cargo.HandlingEvent) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	// Make array if it's the first event with this tracking ID.
	if _, ok := r.events[e.TrackingID]; !ok {
		r.events[e.TrackingID] = make([]cargo.HandlingEvent, 0)
	}
	r.events[e.TrackingID] = append(r.events[e.TrackingID], e)
}

func (r *handlingEventRepository) QueryHandlingHistory(id cargo.TrackingID) cargo.HandlingHistory {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return cargo.HandlingHistory{HandlingEvents: r.events[id]}
}

// NewHandlingEventRepository returns a new instance of a in-memory handling event repository.
func NewHandlingEventRepository() cargo.HandlingEventRepository {
	return &handlingEventRepository{
		events: make(map[cargo.TrackingID][]cargo.HandlingEvent),
	}
}