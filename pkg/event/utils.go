package event

import "reflect"

func GetStructName(data interface{}) string {
	return reflect.TypeOf(data).Name()
}
