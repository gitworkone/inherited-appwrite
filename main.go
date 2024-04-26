package main

//-- not works. build error:(2024.0426)
import (
	"fmt"
	"testappwrite/modAppWrite"
)

func main() {
	err := modAppWrite.GetSingleCAWDb().Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
}
