package pkg

import "encoding/gob"

func Init() { 
	gob.Register(&TestData{})
}
