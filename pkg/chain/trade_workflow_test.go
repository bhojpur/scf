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
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	EXPORTER   = "YunicaRetail"
	EXPBANK    = "YunicaBank"
	EXPBALANCE = 100000
	IMPORTER   = "CornsysTech"
	IMPBANK    = "CornsysBank"
	IMPBALANCE = 200000
	CARRIER    = "JubbalFrieght"
	REGAUTH    = "KrishiDepartment"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkNoState(t *testing.T, stub *shim.MockStub, name string) {
	bytes := stub.State[name]
	if bytes != nil {
		fmt.Println("State", name, "should be absent; found value")
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was", string(bytes), "and not", value, "as expected")
		t.FailNow()
	}
}

func checkBadQuery(t *testing.T, stub *shim.MockStub, function string, name string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(function), []byte(name)})
	if res.Status == shim.OK {
		fmt.Println("Query", name, "unexpectedly succeeded")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, function string, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(function), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	payload := string(res.Payload)
	if payload != value {
		fmt.Println("Query value", name, "was", payload, "and not", value, "as expected")
		t.FailNow()
	}
}

func checkQueryArgs(t *testing.T, stub *shim.MockStub, args [][]byte, value string) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Query", string(args[1]), "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", string(args[1]), "failed to get value")
		t.FailNow()
	}
	payload := string(res.Payload)
	if payload != value {
		fmt.Println("Query value", string(args[1]), "was", payload, "and not", value, "as expected")
		t.FailNow()
	}
}

func checkBadInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status == shim.OK {
		fmt.Println("Invoke", args, "unexpectedly succeeded")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func getInitArguments() [][]byte {
	return [][]byte{[]byte("init"),
		[]byte("YunicaRetail"),
		[]byte("YunicaBank"),
		[]byte("100000"),
		[]byte("CornsysTech"),
		[]byte("CornsysBank"),
		[]byte("200000"),
		[]byte("JubbalFrieght"),
		[]byte("KrishiDepartment")}
}

func TestTradeWorkflow_Init(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	checkState(t, stub, "Exporter", EXPORTER)
	checkState(t, stub, "ExportersBank", EXPBANK)
	checkState(t, stub, "ExportersAccountBalance", strconv.Itoa(EXPBALANCE))
	checkState(t, stub, "Importer", IMPORTER)
	checkState(t, stub, "ImportersBank", IMPBANK)
	checkState(t, stub, "ImportersAccountBalance", strconv.Itoa(IMPBALANCE))
	checkState(t, stub, "Carrier", CARRIER)
	checkState(t, stub, "RegulatoryAuthority", REGAUTH)
}

func TestTradeWorkflow_Agreement(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'requestTrade'
	tradeID := "2ks89j9"
	amount := 50000
	descGoods := "Wood for Toys"
	checkInvoke(t, stub, [][]byte{[]byte("requestTrade"), []byte(tradeID), []byte(strconv.Itoa(amount)), []byte(descGoods)})

	tradeAgreement := &TradeAgreement{amount, descGoods, REQUESTED, 0}
	tradeAgreementBytes, _ := json.Marshal(tradeAgreement)
	tradeKey, _ := stub.CreateCompositeKey("Trade", []string{tradeID})
	checkState(t, stub, tradeKey, string(tradeAgreementBytes))

	expectedResp := "{\"Status\":\"REQUESTED\"}"
	checkQuery(t, stub, "getTradeStatus", tradeID, expectedResp)

	// Invoke bad 'acceptTrade' and verify unchanged state
	checkBadInvoke(t, stub, [][]byte{[]byte("acceptTrade")})
	badTradeID := "abcd"
	checkBadInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(badTradeID)})
	checkState(t, stub, tradeKey, string(tradeAgreementBytes))
	checkQuery(t, stub, "getTradeStatus", tradeID, expectedResp)

	// Invoke 'acceptTrade' and verify state change
	checkInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(tradeID)})
	tradeAgreement.Status = ACCEPTED
	tradeAgreementBytes, _ = json.Marshal(tradeAgreement)
	checkState(t, stub, tradeKey, string(tradeAgreementBytes))

	expectedResp = "{\"Status\":\"ACCEPTED\"}"
	checkQuery(t, stub, "getTradeStatus", tradeID, expectedResp)
}

func TestTradeWorkflow_LetterOfCredit(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'requestTrade' and 'acceptTrade'
	tradeID := "2ks89j9"
	amount := 50000
	descGoods := "Wood for Toys"
	checkInvoke(t, stub, [][]byte{[]byte("requestTrade"), []byte(tradeID), []byte(strconv.Itoa(amount)), []byte(descGoods)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(tradeID)})

	// Invoke 'requestLC'
	checkInvoke(t, stub, [][]byte{[]byte("requestLC"), []byte(tradeID)})
	letterOfCredit := &LetterOfCredit{"", "", EXPORTER, amount, []string{}, REQUESTED}
	letterOfCreditBytes, _ := json.Marshal(letterOfCredit)
	lcKey, _ := stub.CreateCompositeKey("LetterOfCredit", []string{tradeID})
	checkState(t, stub, lcKey, string(letterOfCreditBytes))

	expectedResp := "{\"Status\":\"REQUESTED\"}"
	checkQuery(t, stub, "getLCStatus", tradeID, expectedResp)

	// Invoke bad 'issueLC' and verify unchanged state
	checkBadInvoke(t, stub, [][]byte{[]byte("issueLC")})
	badTradeID := "abcd"
	checkBadInvoke(t, stub, [][]byte{[]byte("issueLC"), []byte(badTradeID)})
	checkState(t, stub, lcKey, string(letterOfCreditBytes))

	// Invoke 'acceptLC' prematurely and verify failure and unchanged state
	checkBadInvoke(t, stub, [][]byte{[]byte("acceptLC"), []byte(badTradeID)})
	checkState(t, stub, lcKey, string(letterOfCreditBytes))
	checkQuery(t, stub, "getLCStatus", tradeID, expectedResp)

	// Invoke 'issueLC'
	lcID := "lc8349"
	expirationDate := "12/31/2018"
	doc1 := "E/L"
	doc2 := "B/L"
	checkInvoke(t, stub, [][]byte{[]byte("issueLC"), []byte(tradeID), []byte(lcID), []byte(expirationDate), []byte(doc1), []byte(doc2)})
	letterOfCredit = &LetterOfCredit{lcID, expirationDate, EXPORTER, amount, []string{doc1, doc2}, ISSUED}
	letterOfCreditBytes, _ = json.Marshal(letterOfCredit)
	checkState(t, stub, lcKey, string(letterOfCreditBytes))

	expectedResp = "{\"Status\":\"ISSUED\"}"
	checkQuery(t, stub, "getLCStatus", tradeID, expectedResp)

	// Invoke 'acceptLC'
	checkInvoke(t, stub, [][]byte{[]byte("acceptLC"), []byte(tradeID)})
	letterOfCredit = &LetterOfCredit{lcID, expirationDate, EXPORTER, amount, []string{doc1, doc2}, ACCEPTED}
	letterOfCreditBytes, _ = json.Marshal(letterOfCredit)
	checkState(t, stub, lcKey, string(letterOfCreditBytes))

	expectedResp = "{\"Status\":\"ACCEPTED\"}"
	checkQuery(t, stub, "getLCStatus", tradeID, expectedResp)
}

func TestTradeWorkflow_ExportLicense(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'requestTrade', 'acceptTrade', 'requestLC', 'issueLC', 'acceptLC'
	tradeID := "2ks89j9"
	amount := 50000
	descGoods := "Wood for Toys"
	checkInvoke(t, stub, [][]byte{[]byte("requestTrade"), []byte(tradeID), []byte(strconv.Itoa(amount)), []byte(descGoods)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(tradeID)})
	checkInvoke(t, stub, [][]byte{[]byte("requestLC"), []byte(tradeID)})
	lcID := "lc8349"
	lcExpirationDate := "12/31/2018"
	doc1 := "E/L"
	doc2 := "B/L"
	checkInvoke(t, stub, [][]byte{[]byte("issueLC"), []byte(tradeID), []byte(lcID), []byte(lcExpirationDate), []byte(doc1), []byte(doc2)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptLC"), []byte(tradeID)})

	// Issue 'requestEL'
	checkInvoke(t, stub, [][]byte{[]byte("requestEL"), []byte(tradeID)})
	exportLicense := &ExportLicense{"", "", EXPORTER, CARRIER, descGoods, REGAUTH, REQUESTED}
	exportLicenseBytes, _ := json.Marshal(exportLicense)
	elKey, _ := stub.CreateCompositeKey("ExportLicense", []string{tradeID})
	checkState(t, stub, elKey, string(exportLicenseBytes))

	expectedResp := "{\"Status\":\"REQUESTED\"}"
	checkQuery(t, stub, "getELStatus", tradeID, expectedResp)

	elID := "el979"
	elExpirationDate := "4/30/2019"

	// Invoke bad 'issueEL' and verify unchanged state
	checkBadInvoke(t, stub, [][]byte{[]byte("issueEL")})
	badTradeID := "abcd"
	checkBadInvoke(t, stub, [][]byte{[]byte("issueEL"), []byte(badTradeID), []byte(elID), []byte(elExpirationDate)})
	checkState(t, stub, elKey, string(exportLicenseBytes))
	checkQuery(t, stub, "getELStatus", tradeID, expectedResp)

	// Invoke 'issueEL' and verify state change
	checkInvoke(t, stub, [][]byte{[]byte("issueEL"), []byte(tradeID), []byte(elID), []byte(elExpirationDate)})
	exportLicense = &ExportLicense{elID, elExpirationDate, EXPORTER, CARRIER, descGoods, REGAUTH, ISSUED}
	exportLicenseBytes, _ = json.Marshal(exportLicense)
	checkState(t, stub, elKey, string(exportLicenseBytes))

	expectedResp = "{\"Status\":\"ISSUED\"}"
	checkQuery(t, stub, "getELStatus", tradeID, expectedResp)
}

func TestTradeWorkflow_ShipmentInitiation(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'requestTrade', 'acceptTrade', 'requestLC', 'issueLC', 'acceptLC', 'requestEL', 'issueEL'
	tradeID := "2ks89j9"
	amount := 50000
	descGoods := "Wood for Toys"
	checkInvoke(t, stub, [][]byte{[]byte("requestTrade"), []byte(tradeID), []byte(strconv.Itoa(amount)), []byte(descGoods)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(tradeID)})
	checkInvoke(t, stub, [][]byte{[]byte("requestLC"), []byte(tradeID)})
	lcID := "lc8349"
	lcExpirationDate := "12/31/2018"
	doc1 := "E/L"
	doc2 := "B/L"
	checkInvoke(t, stub, [][]byte{[]byte("issueLC"), []byte(tradeID), []byte(lcID), []byte(lcExpirationDate), []byte(doc1), []byte(doc2)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptLC"), []byte(tradeID)})
	checkInvoke(t, stub, [][]byte{[]byte("requestEL"), []byte(tradeID)})
	elID := "el979"
	elExpirationDate := "4/30/2019"
	checkInvoke(t, stub, [][]byte{[]byte("issueEL"), []byte(tradeID), []byte(elID), []byte(elExpirationDate)})

	// Invoke 'prepareShipment'
	checkInvoke(t, stub, [][]byte{[]byte("prepareShipment"), []byte(tradeID)})
	slKey, _ := stub.CreateCompositeKey("Shipment", []string{"Location", tradeID})
	checkState(t, stub, slKey, SOURCE)

	expectedResp := "{\"Location\":\"SOURCE\"}"
	checkQuery(t, stub, "getShipmentLocation", tradeID, expectedResp)

	// Invoke bad 'acceptShipmentAndIssueBL' and verify unchanged state
	checkBadInvoke(t, stub, [][]byte{[]byte("acceptShipmentAndIssueBL")})
	badTradeID := "abcd"
	blID := "bl06678"
	blExpirationDate := "8/31/2018"
	sourcePort := "Woodlands Port"
	destinationPort := "Market Port"
	checkBadInvoke(t, stub, [][]byte{[]byte("acceptShipmentAndIssueBL"), []byte(badTradeID), []byte(blID), []byte(blExpirationDate), []byte(sourcePort), []byte(destinationPort)})
	blKey, _ := stub.CreateCompositeKey("BillOfLading", []string{tradeID})
	checkNoState(t, stub, blKey)
	checkBadQuery(t, stub, "getBillOfLading", tradeID)

	// Invoke 'acceptShipmentAndIssueBL' and verify state change
	checkInvoke(t, stub, [][]byte{[]byte("acceptShipmentAndIssueBL"), []byte(tradeID), []byte(blID), []byte(blExpirationDate), []byte(sourcePort), []byte(destinationPort)})
	billOfLading := &BillOfLading{blID, blExpirationDate, EXPORTER, CARRIER, descGoods, amount, IMPBANK, sourcePort, destinationPort}
	billOfLadingBytes, _ := json.Marshal(billOfLading)
	checkState(t, stub, blKey, string(billOfLadingBytes))
	checkQuery(t, stub, "getBillOfLading", tradeID, string(billOfLadingBytes))
}

func TestTradeWorkflow_PaymentFulfilment(t *testing.T) {
	scc := new(TradeWorkflow)
	scc.testMode = true
	stub := shim.NewMockStub("Trade Workflow", scc)

	// Init
	checkInit(t, stub, getInitArguments())

	// Invoke 'requestTrade', 'acceptTrade', 'requestLC', 'issueLC', 'acceptLC', 'requestEL', 'issueEL', 'prepareShipment', 'acceptShipmentAndIssueBL'
	tradeID := "2ks89j9"
	amount := 50000
	descGoods := "Wood for Toys"
	checkInvoke(t, stub, [][]byte{[]byte("requestTrade"), []byte(tradeID), []byte(strconv.Itoa(amount)), []byte(descGoods)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptTrade"), []byte(tradeID)})
	checkInvoke(t, stub, [][]byte{[]byte("requestLC"), []byte(tradeID)})
	lcID := "lc8349"
	lcExpirationDate := "12/31/2018"
	doc1 := "E/L"
	doc2 := "B/L"
	checkInvoke(t, stub, [][]byte{[]byte("issueLC"), []byte(tradeID), []byte(lcID), []byte(lcExpirationDate), []byte(doc1), []byte(doc2)})
	checkInvoke(t, stub, [][]byte{[]byte("acceptLC"), []byte(tradeID)})
	checkInvoke(t, stub, [][]byte{[]byte("requestEL"), []byte(tradeID)})
	elID := "el979"
	elExpirationDate := "4/30/2019"
	checkInvoke(t, stub, [][]byte{[]byte("issueEL"), []byte(tradeID), []byte(elID), []byte(elExpirationDate)})
	checkInvoke(t, stub, [][]byte{[]byte("prepareShipment"), []byte(tradeID)})
	blID := "bl06678"
	blExpirationDate := "8/31/2018"
	sourcePort := "Woodlands Port"
	destinationPort := "Market Port"
	checkInvoke(t, stub, [][]byte{[]byte("acceptShipmentAndIssueBL"), []byte(tradeID), []byte(blID), []byte(blExpirationDate), []byte(sourcePort), []byte(destinationPort)})

	// Invoke 'requestPayment'
	checkInvoke(t, stub, [][]byte{[]byte("requestPayment"), []byte(tradeID)})
	paymentKey, _ := stub.CreateCompositeKey("Payment", []string{tradeID})
	checkState(t, stub, paymentKey, REQUESTED)

	// Invoke 'makePayment'
	checkInvoke(t, stub, [][]byte{[]byte("makePayment"), []byte(tradeID)})
	checkNoState(t, stub, paymentKey)
	// Verify account and payment balances
	payment := amount / 2
	expBalanceStr := strconv.Itoa(EXPBALANCE + payment)
	impBalanceStr := strconv.Itoa(IMPBALANCE - payment)
	checkState(t, stub, expBalKey, expBalanceStr)
	checkState(t, stub, impBalKey, impBalanceStr)
	tradeAgreement := &TradeAgreement{amount, descGoods, ACCEPTED, payment}
	tradeAgreementBytes, _ := json.Marshal(tradeAgreement)
	tradeKey, _ := stub.CreateCompositeKey("Trade", []string{tradeID})
	checkState(t, stub, tradeKey, string(tradeAgreementBytes))

	// Check queries
	checkBadQuery(t, stub, "getAccountBalance", tradeID)
	expectedResp := "{\"Balance\":\"" + expBalanceStr + "\"}"
	checkQueryArgs(t, stub, [][]byte{[]byte("getAccountBalance"), []byte(tradeID), []byte("exporter")}, expectedResp)

	expectedResp = "{\"Balance\":\"" + impBalanceStr + "\"}"
	checkQueryArgs(t, stub, [][]byte{[]byte("getAccountBalance"), []byte(tradeID), []byte("importer")}, expectedResp)

	// Deliver shipment to final location
	checkInvoke(t, stub, [][]byte{[]byte("updateShipmentLocation"), []byte(tradeID), []byte(DESTINATION)})
	slKey, _ := stub.CreateCompositeKey("Shipment", []string{"Location", tradeID})
	checkState(t, stub, slKey, DESTINATION)

	// Invoke 'requestPayment' and 'makePayment'
	checkInvoke(t, stub, [][]byte{[]byte("requestPayment"), []byte(tradeID)})
	checkState(t, stub, paymentKey, REQUESTED)
	checkInvoke(t, stub, [][]byte{[]byte("makePayment"), []byte(tradeID)})
	checkNoState(t, stub, paymentKey)

	// Verify account and payment balances, and check queries
	expBalanceStr = strconv.Itoa(EXPBALANCE + amount)
	impBalanceStr = strconv.Itoa(IMPBALANCE - amount)
	checkState(t, stub, expBalKey, expBalanceStr)
	checkState(t, stub, impBalKey, impBalanceStr)
	tradeAgreement = &TradeAgreement{amount, descGoods, ACCEPTED, amount}
	tradeAgreementBytes, _ = json.Marshal(tradeAgreement)
	checkState(t, stub, tradeKey, string(tradeAgreementBytes))

	expectedResp = "{\"Balance\":\"" + expBalanceStr + "\"}"
	checkQueryArgs(t, stub, [][]byte{[]byte("getAccountBalance"), []byte(tradeID), []byte("exporter")}, expectedResp)

	expectedResp = "{\"Balance\":\"" + impBalanceStr + "\"}"
	checkQueryArgs(t, stub, [][]byte{[]byte("getAccountBalance"), []byte(tradeID), []byte("importer")}, expectedResp)
}
