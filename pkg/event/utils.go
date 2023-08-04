package event

import "reflect"

func getStructName(listener interface{}) string {
	return reflect.TypeOf(listener).Name()
}
