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

import "github.com/bhojpur/scf/pkg/logistic/location"

// A set of sample voyages.
var (
	V100 = New("V100", Schedule{
		[]CarrierMovement{
			{DepartureLocation: location.CNHKG, ArrivalLocation: location.JNTKO},
			{DepartureLocation: location.JNTKO, ArrivalLocation: location.USNYC},
		},
	})

	V300 = New("V300", Schedule{
		[]CarrierMovement{
			{DepartureLocation: location.JNTKO, ArrivalLocation: location.NLRTM},
			{DepartureLocation: location.NLRTM, ArrivalLocation: location.DEHAM},
			{DepartureLocation: location.DEHAM, ArrivalLocation: location.AUMEL},
			{DepartureLocation: location.AUMEL, ArrivalLocation: location.JNTKO},
		},
	})

	V400 = New("V400", Schedule{
		[]CarrierMovement{
			{DepartureLocation: location.DEHAM, ArrivalLocation: location.SESTO},
			{DepartureLocation: location.SESTO, ArrivalLocation: location.FIHEL},
			{DepartureLocation: location.FIHEL, ArrivalLocation: location.DEHAM},
		},
	})
)

// These voyages are hard-coded into the current pathfinder. Make sure
// they exist.
var (
	V0100S = New("0100S", Schedule{[]CarrierMovement{}})
	V0200T = New("0200T", Schedule{[]CarrierMovement{}})
	V0300A = New("0300A", Schedule{[]CarrierMovement{}})
	V0301S = New("0301S", Schedule{[]CarrierMovement{}})
	V0400S = New("0400S", Schedule{[]CarrierMovement{}})
)