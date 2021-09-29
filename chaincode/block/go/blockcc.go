/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type signature struct {
	Sig     string    `json:sig`
	SigDate time.Time `json:sigDate`
	Lat     float64   `json:lat`
	Lang    float64   `json:lang`
	Notes   string    `json:notesAtManufacturer`
}

type ShipmentOrder struct {
	BookingNumber int    `json:bookingNo`
	BookingDate   string `json:bookingDate`
	Shipper       string `json:shipper`
	Retailer      string `json:retailer`
	Manufacturer  string `json:manufacturer`

	PickupDate       string `json:PickupDate`
	PickupLocation   string `json:PickupLocation`
	DeliveryDate     string `json:deliveryDate`
	DeliveryLocation string `json:deliveryLocation`
	Notes            string `json:notes`

	ManfSign           signature `json:docType`
	RetailerSign       signature `json:docType`
	DriverManfSign     signature `json:docType`
	DriverRetailerSign signature `json:docType`

	Driver        string         `json:driverId`
	VehicleType   string         `json:vehicleType`
	VehicleNumber string         `json:VehicleNumber`
	Status        shipmentStatus `json:shipmentStatus`
	StatusDate    time.Time      `json:shipmentTime`

	Packages []Package
}

type Package struct {
	RWBNumber   string        `json:RWBNo`
	HsnNumber   string        `json:hsnNo`
	ProductName string        `json:productName`
	ProductType string        `json:productType`
	ProductQty  int           `json:qty`
	ProductSize string        `json:ProductSize`
	Status      packageStatus `json:packageStatus`
	StatusDate  time.Time     `json:shipmentStatus`
}

type shipmentStatus int

const (
	ShipmentDeleted     shipmentStatus = 0
	ShipmentOrderPlaced shipmentStatus = 1
	ShipmentAccepted    shipmentStatus = 2
	ShipmentRejected    shipmentStatus = 3
	DriverAssigned      shipmentStatus = 4
	ShipmentPicked      shipmentStatus = 6
	Intransit           shipmentStatus = 7
	Delivered           shipmentStatus = 8
	Verified            shipmentStatus = 10
	Completed           shipmentStatus = 11
)

type packageStatus int

const (
	PackageDeleted   packageStatus = 0
	Created          packageStatus = 1
	ShipperAccepted  packageStatus = 2
	ShipperRejected  packageStatus = 3
	DriverAccepted   packageStatus = 5
	DriverRejected   packageStatus = 6
	RetailerAccepted packageStatus = 7
	RetailerRejected packageStatus = 8
)

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) createShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 5")
	}
	BookingNo, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Invalid Booking Number")
	}
	var ShipmentOrder = ShipmentOrder{BookingNumber: BookingNo, BookingDate: args[1], Shipper: args[2], Retailer: args[3], Manufacturer: args[4], PickupDate: args[5], PickupLocation: args[6], DeliveryDate: args[7], DeliveryLocation: args[8], Notes: args[9], Status: 1, StatusDate: time.Now()}

	ShipmentOrderAsBytes, _ := json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) createPackage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 7 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 7")
	}

	Qty, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Invalid Quantity")
	}
	var Package = Package{RWBNumber: args[1], HsnNumber: args[2], ProductName: args[3], ProductType: args[4], ProductQty: Qty, ProductSize: args[6], Status: 1, StatusDate: time.Now()}
	ShipmentOrder := ShipmentOrder{}
	ShipmentAsBytes, _ := APIstub.GetState(args[0])
	json.Unmarshal(ShipmentAsBytes, &ShipmentOrder)
	ShipmentOrder.Packages = append(ShipmentOrder.Packages, Package)
	ShipmentOrderAsBytes, _ := json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)
	return shim.Success(nil)
}

func (s *SmartContract) addShipmentDRS(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 4")
	}
	Stat, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Invalid Status")
	}

	ShipmentOrderAsBytes, _ := APIstub.GetState(args[0])
	ShipmentOrder := ShipmentOrder{}

	json.Unmarshal(ShipmentOrderAsBytes, &ShipmentOrder)
	if Stat == 1 {
		ShipmentOrder.Status = DriverAssigned
	} else if Stat == 2 {
		//shipment rejected
		ShipmentOrder.Status = ShipmentRejected
	}
	ShipmentOrder.Driver = args[2]
	ShipmentOrder.VehicleType = args[3]
	ShipmentOrder.VehicleNumber = args[4]
	ShipmentOrder.Status = 4

	ShipmentOrderAsBytes, _ = json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) driverAcceptOrReject(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 3")
	}
	RWB := args[1]

	Stat, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid Status")
	}
	status := Created
	if Stat == 1 {
		//driver accepted
		status = DriverAccepted
	} else if Stat == 2 {
		//driver rejected
		status = DriverRejected
	}
	ShipmentOrderAsBytes, _ := APIstub.GetState(args[0])
	ShipmentOrder := ShipmentOrder{}

	json.Unmarshal(ShipmentOrderAsBytes, &ShipmentOrder)
	for i := 1; i < len(ShipmentOrder.Packages); i++ {
		if ShipmentOrder.Packages[i].RWBNumber == RWB {
			ShipmentOrder.Packages[i].Status = status
		}
	}

	ShipmentOrderAsBytes, _ = json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) shipmentPickup(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 7 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 7")
	}

	manfSigt, errt := time.Parse("2006-01-02", args[2])
	if errt != nil {
		return shim.Error("error in parsing date")
	}

	driverSigt, errt := time.Parse("2006-01-02", args[4])
	if errt != nil {
		return shim.Error("error in parsing date")
	}

	LT, err := strconv.ParseFloat(args[5], 64)
	if err != nil {
		return shim.Error("Invalid Lat Value")
	}
	LG, err := strconv.ParseFloat(args[6], 64)
	if err != nil {
		return shim.Error("Invalid Lat Value")
	}

	var ManfSig = signature{Sig: args[1], SigDate: manfSigt, Lat: LT, Lang: LG}
	var DriverSig = signature{Sig: args[3], SigDate: driverSigt, Lat: LT, Lang: LG}

	ShipmentOrderAsBytes, _ := APIstub.GetState(args[0])
	ShipmentOrder := ShipmentOrder{}

	json.Unmarshal(ShipmentOrderAsBytes, &ShipmentOrder)
	ShipmentOrder.Status = ShipmentPicked
	ShipmentOrder.ManfSign = ManfSig
	ShipmentOrder.DriverManfSign = DriverSig
	ShipmentOrder.Notes = args[7]

	ShipmentOrderAsBytes, _ = json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) shipmentDeliver(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 7 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 7")
	}
	retrSigt, errt := time.Parse("2006-01-02", args[2])
	if errt != nil {
		return shim.Error("error in parsing date")
	}

	driverSigt, errt := time.Parse("2006-01-02", args[4])
	if errt != nil {
		return shim.Error("error in parsing date")
	}

	LT, err := strconv.ParseFloat(args[5], 64)
	if err != nil {
		return shim.Error("Invalid Lat Value")
	}
	LG, err := strconv.ParseFloat(args[6], 64)
	if err != nil {
		return shim.Error("Invalid Lat Value")
	}
	var RetailerSig = signature{Sig: args[1], SigDate: retrSigt, Lat: LT, Lang: LG}
	var DriverSig = signature{Sig: args[3], SigDate: driverSigt, Lat: LT, Lang: LG}

	ShipmentOrderAsBytes, _ := APIstub.GetState(args[0])
	ShipmentOrder := ShipmentOrder{}

	json.Unmarshal(ShipmentOrderAsBytes, &ShipmentOrder)
	ShipmentOrder.Status = ShipmentDeleted
	ShipmentOrder.RetailerSign = RetailerSig
	ShipmentOrder.DriverRetailerSign = DriverSig
	ShipmentOrder.Notes = args[7]

	ShipmentOrderAsBytes, _ = json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updatePackageOutbound(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting Minimum 3")
	}

	// Stat, err := strconv.Atoi(args[2])
	// if err != nil {
	// 	return shim.Error("Invalid Status")
	// }
	// status := RetailerAccepted
	// if Stat == 1 {
	// 	//driver accepted
	// 	status = RetailerAccepted
	// } else if Stat == 2 {
	// 	//driver rejected
	// 	status = RetailerRejected
	// }

	ShipmentOrderAsBytes, _ := APIstub.GetState(args[0])
	ShipmentOrder := ShipmentOrder{}

	json.Unmarshal(ShipmentOrderAsBytes, &ShipmentOrder)
	//ShipmentOrder.Packages[0].Status = status
	ShipmentOrder.Status = Delivered

	ShipmentOrderAsBytes, _ = json.Marshal(ShipmentOrder)
	APIstub.PutState(args[0], ShipmentOrderAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "getAllShipments" {
		return s.getAllShipments(APIstub)
	} else if function == "createPackage" {
		return s.createPackage(APIstub, args)
	} else if function == "getAllPackages" {
		return s.getAllPackages(APIstub)
	} else if function == "createShipment" {
		return s.createShipment(APIstub, args)
	} else if function == "queryPackage" {
		return s.queryPackage(APIstub, args)
	} else if function == "queryShipment" {
		return s.queryShipment(APIstub, args)
	} else if function == "addShipmentDRS" {
		return s.addShipmentDRS(APIstub, args)
	} else if function == "driverAcceptOrReject" {
		return s.driverAcceptOrReject(APIstub, args)
	} else if function == "shipmentPickup" {
		return s.shipmentPickup(APIstub, args)
	} else if function == "shipmentDeliver" {
		return s.shipmentDeliver(APIstub, args)
	} else if function == "updatePackageOutbound" {
		return s.updatePackageOutbound(APIstub, args)
	} else if function == "shipmentHistory" {
		return s.shipmentHistory(APIstub, args)
	} else if function == "getShipmentOrdersByShipper" {
		shipperName := args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"Shipper\":\"%s\"},\"sort\": [{\"BookingDate\": \"desc\"}]}", shipperName)
		return s.getQueryResultByStr(APIstub, queryString)
	} else if function == "getShipmentOrdersByManufacturer" {
		manufacturerName := args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"Manufacturer\":\"%s\"},\"sort\": [{\"BookingDate\": \"desc\"}]}", manufacturerName)
		return s.getQueryResultByStr(APIstub, queryString)
	} else if function == "getShipmentOrdersByRetailer" {
		retailerName := args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"Retailer\":\"%s\"},\"sort\": [{\"BookingDate\": \"desc\"}]}", retailerName)
		return s.getQueryResultByStr(APIstub, queryString)
	} else if function == "getShipmentOrdersByDriver" {
		driverId := args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"Driver\":\"%s\"},\"sort\": [{\"BookingDate\": \"desc\"}]}", driverId)
		return s.getQueryResultByStr(APIstub, queryString)
	} else if function == "getShipmentOrdersByVehicle" {
		vehicleNo := args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"Vehicle\":\"%s\"},\"sort\": [{\"BookingDate\": \"desc\"}]}", vehicleNo)
		return s.getQueryResultByStr(APIstub, queryString)
	}

	// return shim.Error("Invalid Smart Contract function name.")
	return shim.Success(nil)
}

func (s *SmartContract) getQueryResultByStr(stub shim.ChaincodeStubInterface, queryString string) sc.Response {

	resultsIterator, err := stub.GetQueryResult(queryString)
	defer resultsIterator.Close()
	if err != nil {
		return shim.Error(err.Error())
	}
	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse,
			err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("-  getQueryResultByStr:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	shipmentAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(string(shipmentAsBytes))
	return shim.Success(shipmentAsBytes)
}

func (s *SmartContract) queryPackage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	packageAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(string(packageAsBytes))
	return shim.Success(packageAsBytes)
}

func (s *SmartContract) getAllShipments(APIstub shim.ChaincodeStubInterface) sc.Response {

	var startKey, endKey string

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	ShipmentList := []ShipmentOrder{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		ShipmentOrder := ShipmentOrder{}
		json.Unmarshal(queryResponse.Value, &ShipmentOrder)
		ShipmentList = append(ShipmentList, ShipmentOrder)
	}

	fmt.Printf("- queryAllMeds:\n%+v\n", ShipmentList)
	ShipmentOrderAsBytes, err := json.Marshal(ShipmentList)
	return shim.Success(ShipmentOrderAsBytes)
}

func (s *SmartContract) getAllPackages(APIstub shim.ChaincodeStubInterface) sc.Response {

	var startKey, endKey string

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	PackageList := []Package{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		Package := Package{}
		json.Unmarshal(queryResponse.Value, &Package)
		PackageList = append(PackageList, Package)
	}

	fmt.Printf("- queryAllMeds:\n%+v\n", PackageList)
	PackageAsBytes, err := json.Marshal(PackageList)
	return shim.Success(PackageAsBytes)
}

//get Shipment History ShipmentHistory
func (s *SmartContract) shipmentHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	BookingNumber := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(BookingNumber)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		if queryResponse.IsDelete == true {
			buffer.WriteString("\"\"")
		}
		if queryResponse.IsDelete == false {
			buffer.WriteString(string(queryResponse.Value))
		}
		//buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString(",\"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Timestamp.String())
		buffer.WriteString("\"")

		buffer.WriteString(",\"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(queryResponse.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- History:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	shps := []ShipmentOrder{
		ShipmentOrder{BookingNumber: 1122, BookingDate: time.Now().Format("20060102150405"), Shipper: "shp1"},
		ShipmentOrder{BookingNumber: 1113, BookingDate: time.Now().Format("20060102150405"), Shipper: "shp1"},
	}

	rec1, err := json.Marshal(shps[0])
	if err != nil {
		return shim.Error("error parsing record one")
	}
	APIstub.PutState("1122", rec1)
	rec2, err1 := json.Marshal(shps[1])
	if err1 != nil {
		return shim.Error("error parsing record two")
	}
	APIstub.PutState("1123", rec2)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
