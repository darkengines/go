package main

import _ "user"
import _ "databasee"
import "communication/http"
import _ "communication/http/json"

import (
	"composition"
)

type SchemaUpdater interface {
	UpdateSchema() error
}

func main() {
	
	instance := composition.GetInstance()
	updaters := make([]SchemaUpdater, 0)
	err := instance.Import(&updaters)
	if (err != nil) {
		panic(err)
	}
	
	for _, updater := range updaters {
		err := updater.(SchemaUpdater).UpdateSchema()
		if (err != nil) {
			panic(err)
		}
	}
	
	s := http.NewServer()
	s.Start()
	
}