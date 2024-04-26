package main

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
