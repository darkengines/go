package composition

import "reflect"
import (
	"github.com/ahmetalpbalkan/go-linq"
)

var exportedTypes = make([]reflect.Type, 0)
var exportedValues = make(map[reflect.Type]interface{})

func Export(exportedType reflect.Type) {
	alreadyExported := !linq.From(exportedTypes).AnyWith(func(existingExportedType reflect.Type) (bool, error) {
		return existingExportedType == exportedType, nil
	})
	if (!alreadyExported) {
		exportedTypes = append(exportedTypes, exportedType)
	}
}

func GetExportedValues(requestedExportedType reflect.Type) (interface{}) {
	matches := linq.From(exportedTypes).Where(func(exportedType reflect.Type) bool {
		return exportedType.AssignableTo(requestedExportedType)
	})
	values := make([]interface{}, 0)
	for _, match := range(matches) {
		exportedValue, found := exportedValues[match]
		if (!found) {
			exportedValue = reflect.New(match)
			exportedValues[match] = exportedValue
		}
		values = append(values, exportedValue)
	}
	return values
}