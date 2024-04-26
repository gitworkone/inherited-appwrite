package modAppWrite

import (
	"github.com/appwrite/sdk-for-go"
)

type CAWDb struct {
}

var g_singleCAWDb *CAWDb = &CAWDb{}

func GetSingleCAWDb() *CAWDb {
	return g_singleCAWDb
}

func (pInst *CAWDb) Initialize() error {
	client := appwrite.NewClient()
	client.SetProject("5df5acd0d48c2")

	return nil
}
