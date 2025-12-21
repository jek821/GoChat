package main

import (
	Server "GoChat/internal/server"
	"GoChat/pkg"
	"fmt"
)


var testIp = "127.0.0.1"
var testPort = ":8080"

func main() {
	//EncodeDecodeTestData()
	StartServerTest()
}

func EncodeDecodeTestData() {
	pkg.Init()
	dataForTest := 5
	testData :=	pkg.CreateTestData(dataForTest)
	testTransmission :=	pkg.CreateTransmission(testData.DataCode, testData)
	encodedTransmission, err := testTransmission.EncodeTransmission()
	if err != nil {
		fmt.Println("Failed Encoding Test")
	}
	decodedTransmission, err := pkg.DecodeTransmission(encodedTransmission)
	if err != nil {
		fmt.Println("Failed Decoding Test")
	}
	switch decodedTransmission.Code {
	case (pkg.TestDataCode):
		TransmissionData := decodedTransmission.Data.(*pkg.TestData)
		if (TransmissionData.Contents == dataForTest){
		fmt.Println("Encode Decode Test Passed!")
	}
	default:
	fmt.Println("Unknown Data type used in Encode Decode Testing")
	}
}



func StartServerTest() {
	server :=	Server.CreateServer(testIp, testPort)
	err := server.StartServer()
	if err != nil {
		fmt.Println("ERROR IN START SERVER TEST")
	}
}

func StarterServerClientConnectTest() {
	server := Server.CreateServer(testIp, testPort)
	err := server.StartServer()
	if err != nil {
		fmt.Println("ERROR IN START SERVER CLIENT CONNECT TEST") } }
