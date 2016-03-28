package main

import (
	"github.com/darkengines/composition"
	"reflect"
)

type SchemaUpdater interface {
	UpdateSchema()
}

func main() {
	updaters := composition.GetExportedValues(reflect.TypeOf((*SchemaUpdater)(nil)).Elem())
	for _, updater := range updaters {
		updater.(SchemaUpdater).UpdateSchema()
	}
}