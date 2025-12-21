package pkg

import "encoding/gob"

func Init() { 
	// Transmission Struct holds Data{} interface
	// Data constructors return pointers so we must register pointer types
	// with gob so that it knows how to encrypt and decrypt transmissions.
	gob.Register(&TestData{}) 
}
