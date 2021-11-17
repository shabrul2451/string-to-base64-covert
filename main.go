package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	str := base64.StdEncoding.EncodeToString([]byte("shabrulislam2451:86cGLaTxhdas4qnvN3ws"))
	fmt.Println(str)

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", data)
}
