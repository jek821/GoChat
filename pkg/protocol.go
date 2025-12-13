package pkg


/* 
Here we Define the Transmission Struct, This Struct is used to encapsulate all data sent between Clients, Servers, and ClientHandlers
This provides a standard type that we can decode all bytecode transmissions into before seeing what type the data actually is VIA the DataCode

To Encode a Transmission we first encode the inner data, and then the transmission entirely using the GOB library

There is also a decode method to convert bytecode back to workable struct types, this method will need to be mantained as new DataTypes are added,
each time we decode a transmission we will have to read the DataCode and then use that information to decide how to decode the inner Data.   
*/
type Transmission struct {
	Code DataCode
	Data Data  
}

func CreateTransmission(DC DataCode, D Data) *Transmission {
	return &Transmission{Code: DC, Data: D}
}


func EncodeTransmission(T Transmission) *[]byte


type DataCode uint8 

const (
	TestDataCode DataCode = iota
)

type Data interface{
	isData()
}


type TestData struct{
	Contents int
}

func (TD *TestData) isData(){}



