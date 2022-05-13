package chain

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
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getTradeKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	tradeKey, err := stub.CreateCompositeKey("Trade", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return tradeKey, nil
	}
}

func getLCKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	lcKey, err := stub.CreateCompositeKey("LetterOfCredit", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return lcKey, nil
	}
}

func getELKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	elKey, err := stub.CreateCompositeKey("ExportLicense", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return elKey, nil
	}
}

func getShipmentLocationKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	shipmentLocationKey, err := stub.CreateCompositeKey("Shipment", []string{"Location", tradeID})
	if err != nil {
		return "", err
	} else {
		return shipmentLocationKey, nil
	}
}

func getBLKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	blKey, err := stub.CreateCompositeKey("BillOfLading", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return blKey, nil
	}
}

func getPaymentKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	paymentKey, err := stub.CreateCompositeKey("Payment", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return paymentKey, nil
	}
}
