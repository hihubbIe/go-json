package json

import (
	"fmt"
	"reflect"
)

func info(o any) (reflect.Value, reflect.Type) {
	return reflect.ValueOf(o), reflect.TypeOf(o)
}

func ToJson(object any) string {

	oVal, oType := info(object)
	oKind := oType.Kind()
	switch oKind {
	case reflect.Int:
		return fmt.Sprintf("\"intValue\": %v", oVal.Int())
	case reflect.Bool:
		return fmt.Sprintf("\"boolValue\": %v", oVal.Bool())
	case reflect.Struct:
		res := "{\n"
		numField := oType.NumField()
		for fi := 0; fi < numField; fi++ {
			field := oType.FieldByIndex([]int{fi})
			//fieldVal := oVal.FieldByIndex([]int{fi})
			res += fmt.Sprintf("\t\"%s\": \"%s\"\n", field.Name /*ToJson(fieldVal)*/, "nested")
		}
		res += "}"
		return res
	default:
		fmt.Println("Unsuported type : " + oKind.String())
	}
	return "err"
}
