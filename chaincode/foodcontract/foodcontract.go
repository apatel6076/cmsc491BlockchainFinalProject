package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type FoodContract struct {
}
var ingredientList []ingredient

type food struct {
	OrderId                string
	FoodId                 string
	ConsumerId             string
	ManufactureId          string
	WholesalerId           string
	RetailerId             string
	LogisticsId            string
	Status                 string
	FoodProcessDate     string
	ManufactureProcessDate string
	WholesaleProcessDate   string
	ShippingProcessDate    string
	RetailProcessDate      string
	OrderPrice             int
	ShippingPrice          int
	DeliveryDate           string
	ingredientList []ingredient  //A list of ingredients
}

type ingredient struct{
	IngredientName			string
	ProviderId				string
	Origin					string
	HarvestDate				string
	SellToManufactorDate	string
}


func (t* FoodContract) init(){
	 ingredientList = []ingredient{
		ingredient{
				IngredientName:  "GRAPES",
				ProviderId: "Provider_1",
				Origin: "Farm_1",
				HarvestDate: "2019-11-29",
				SellToManufactorDate: "2019-12-01",
			},

		ingredient{
				IngredientName:  "SUGAR",
				ProviderId: "Provider_1",
				Origin: "Factory_1",
				HarvestDate: "N/A",
				SellToManufactorDate: "2019-12-02",
		},
		ingredient{
				IngredientName:  "PRESERVATIVES",
				ProviderId: "Provider_1",
				Origin: "Factory_1",
				HarvestDate: "N/A",
				SellToManufactorDate: "2019-12-01",	
		},
	}
}

func (t *FoodContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return setupFoodSupplyChainOrder(stub)
}

func (t *FoodContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createRawFood" {

		return t.createRawFood(stub, args)
	} else if function == "manufactureProcessing" {

		return t.manufactureProcessing(stub, args)
	} else if function == "wholesalerDistribute" {

		return t.wholesalerDistribute(stub, args)
	} else if function == "initiateShipment" {

		return t.initiateShipment(stub, args)
	} else if function == "deliverToRetail" {

		return t.deliverToRetail(stub, args)
	} else if function == "completeOrder" {

		return t.completeOrder(stub, args)
	} else if function == "query" {

		return t.query(stub, args)
	}

	return shim.Error("Invalid function name")
}

func setupFoodSupplyChainOrder(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	orderId := args[0]
	consumerId := args[1]
	orderPrice, _ := strconv.Atoi(args[2])
	shippingPrice, _ := strconv.Atoi(args[3])
	//ingredients := createIngredientsList() //call func to get list of ingredients
	foodContract := food{
		OrderId:       orderId,
		ConsumerId:    consumerId,
		OrderPrice:    orderPrice,
		ShippingPrice: shippingPrice,
		Status:        "order initiated"}

	foodBytes, _ := json.Marshal(foodContract)
	stub.PutState(foodContract.OrderId, foodBytes)

	return shim.Success(nil)
}




//func( f *FoodContract) createIngredientsList(stub shim.ChaincodeStubInterface, args, []string) pb.Response{
//	_, args := stub.GetFunctionAndParameters()
//	loopn till 3 {
//	ingredientName := args[n+1]			
//	providerId := args[n+2]			
//	origin := args[n+3]				
//	harvestDate := args[n+4]				
//	sellToManufactorDate := args[n+5]
//	foodContract := ingredientList{
//		IngredientList: ingredientList,
//		ProviderId: providerId,
//		Origin: origin,
//		HarvestDate: harvestDate,
//		SellToManufactorDate: sellToManufactorDate }
//	}
/*	ing := ingredient{}
	ing.ingredientName = "GRAPES"
	ing.providerId = "Provider_1"
	ing.origin = "Farm_1"
	ing.harvestDate = "2019-11-29"
	ing.sellToManufactorDate = "2019-12-01"

	s = append(s, ing)
	ing2 := ingredient{}
	ing2.ingredientName = "SUGAR"
	ing2.providerId = "Provider_2"
	ing2.origin = "Factory_1"
	ing2.harvestDate = "N/A"
	ing2.sellToManufactorDate = "2019-9-01"

	s = append(s, ing2)
	ing3 := ingredient{}
	ing3.ingredientName = "PRESERVATIVES"
	ing3.providerId = "Provider_3"
	ing3.origin = "Factory_2"
	ing3.harvestDate = "N/A"
	ing3.sellToManufactorDate = "2019-10-01"
	s = append(s, ing3)
	
	return s 

}
*/

func (f *FoodContract) createRawFood(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, _ := stub.GetState(orderId)
	fd := food{}
	json.Unmarshal(foodBytes, &fd)

	if fd.Status == "order initiated" {
		fd.FoodId = "JELLY_1"
		currentts := time.Now()
		fd.FoodProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "Food created"
	} else {
		fmt.Printf("Order not initiated yet")
	}

	foodBytes, _ = json.Marshal(fd)
	stub.PutState(orderId, foodBytes)

	return shim.Success(nil)
}

func (f *FoodContract) manufactureProcessing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}

	if fd.Status == "raw food created" {
		fd.ManufactureId = "Manufacture_1"
		currentts := time.Now()
		fd.ManufactureProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "manufacture Process"
	} else {
		fd.Status = "Error"
		fmt.Printf("Raw food not initiated yet")
	}

	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func (f *FoodContract) wholesalerDistribute(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}

	if fd.Status == "manufacture Process" {
		fd.WholesalerId = "Wholesaler_1"
		currentts := time.Now()
		fd.WholesaleProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "wholesaler distribute"
	} else {
		fd.Status = "Error"
		fmt.Printf("Manufacture not initiated yet")
	}

	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (f *FoodContract) initiateShipment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}

	if fd.Status == "wholesaler distribute" {
		fd.LogisticsId = "LogisticsId_1"
		currentts := time.Now()
		fd.ShippingProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "initiated shipment"
	} else {
		fmt.Printf("Wholesaler not initiated yet")
	}

	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) deliverToRetail(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}

	if fd.Status == "initiated shipment" {
		fd.RetailerId = "Retailer_1"
		currentts := time.Now()
		fd.RetailProcessDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "Retailer started"

	} else {
		fmt.Printf("Shipment not initiated yet")
	}

	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) completeOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	orderId := args[0]
	foodBytes, err := stub.GetState(orderId)
	fd := food{}
	err = json.Unmarshal(foodBytes, &fd)
	if err != nil {
		return shim.Error(err.Error())
	}

	if fd.Status == "Retailer started" {
		currentts := time.Now()
		fd.DeliveryDate = currentts.Format("2006-01-02 15:04:05")
		fd.Status = "Consumer received order"
	} else {
		fmt.Printf("Retailer not initiated yet")
	}

	foodBytes0, _ := json.Marshal(fd)
	err = stub.PutState(orderId, foodBytes0)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (f *FoodContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ENIITY string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expected ENIITY Name")
	}

	ENIITY = args[0]
	Avalbytes, err := stub.GetState(ENIITY)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + ENIITY + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil order for " + ENIITY + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(Avalbytes)
}

func main() {

	err := shim.Start(new(FoodContract))
	if err != nil {
		fmt.Printf("Error creating new Food Contract: %s", err)
	}
}
