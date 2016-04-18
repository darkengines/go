package composition

import "reflect"
import "fmt"

func init() {
	instance = new(container)
	instance.exportedTypes = make(map[reflect.Type]interface{})
	instance.exportedValues = make(map[reflect.Type]reflect.Value)
}

var instance *container

type container struct {
	exportedTypes map[reflect.Type]interface{}
	exportedValues map[reflect.Type]reflect.Value	
}

func GetInstance() *container {
	return instance
}

func(this *container) getValues(requestedExportedType reflect.Type) ([]reflect.Value, error) {
	values := make([]reflect.Value, 0)
	for exportedType, constructor := range(this.exportedTypes) {
		if (exportedType.Implements(requestedExportedType)) {
			exportedValue, found:= this.exportedValues[exportedType]
			if (!found) {
				params := make([]reflect.Value, 0)
				result := reflect.ValueOf(constructor).Call(params)
				exportedValue = result[0]
				var err error
				err = nil
				abstractError := result[1].Interface()
				if (abstractError != nil) {
					err = abstractError.(error)
					return nil, err
				}
				this.exportedValues[exportedType] = exportedValue
			}
			values = append(values, exportedValue)
		}
	}
	return values, nil
}

func(this *container) Export(constructor interface{}) {
	constructorType := reflect.TypeOf(constructor)
	if (constructorType.Kind() != reflect.Func) { 
		panic("Parameter must be a function.")
	}
	if (constructorType.NumIn() > 0) {
		panic("Constructor takes no parameters.")
	}
	isOutOk := constructorType.NumOut() == 2 &&
	constructorType.Out(1) == reflect.TypeOf((*error)(nil)).Elem()
	if (!isOutOk) {
		panic("Constructor must return (<interface{}>, error).")
	}
	returnType := constructorType.Out(0)
	_, exists := this.exportedTypes[returnType]
	if (!exists) {
		fmt.Printf("%s\r\n", returnType)
		this.exportedTypes[returnType] = constructor
	}
}

func (this *container) Import(ptr interface{}) (err error) {
	
	ptrValue := reflect.ValueOf(ptr)
	if (ptrValue.Kind() != reflect.Ptr) {
		panic("ptr must be a pointer.")
	}
	addressable := reflect.Indirect(ptrValue)
	
	var values []reflect.Value
	switch (addressable.Kind()) {
		case(reflect.Slice): {
			requestedExportedType := addressable.Type().Elem()
			fmt.Printf("%s\r\n", requestedExportedType.String())
			values, err = this.getValues(requestedExportedType)
			for _, v := range values {
				fmt.Printf("%s\r\n", v.String())
				addressable.Set(reflect.Append(addressable, v))	
			}
		}
		default: {
			requestedExportedType := addressable.Type()
			values, err = this.getValues(requestedExportedType)
			if (len(values) > 1) {
				panic("Many exported types implement ptr interface.")
			}
			if (len(values) < 1) {
				panic(fmt.Sprintf("Not export implementing %s", requestedExportedType.String()))
			}
			addressable.Set(values[0])
		}
	}
	return err;
}