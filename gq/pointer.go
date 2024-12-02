package gq

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Pointer struct {
	Type Schema
}

func (self Pointer) Key() string {
	return fmt.Sprintf("Pointer[%s]", self.Type.Key())
}

func (self Pointer) Do(params *DoParams) Result {
	value := reflect.ValueOf(params.Value)

	if value.Kind() == reflect.Pointer && value.IsNil() {
		return Result{}
	}

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
		params.Value = nil

		if value.IsValid() {
			params.Value = value.Interface()
		}
	}

	res := self.Type.Do(params)
	data := reflect.ValueOf(res.Data)

	if data.IsValid() && data.Kind() != reflect.Pointer {
		pointer := reflect.New(data.Type())
		pointer.Elem().Set(data)
		res.Data = pointer.Interface()
	}

	return res
}

func (self Pointer) Resolve(params *ResolveParams) Result {
	value := reflect.ValueOf(params.Value)

	if value.Kind() == reflect.Pointer && value.IsNil() {
		return Result{}
	}

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
		params.Value = nil

		if value.IsValid() {
			params.Value = value.Interface()
		}
	}

	res := self.Type.Resolve(params)
	data := reflect.ValueOf(res.Data)

	if data.IsValid() && data.Kind() != reflect.Pointer {
		pointer := reflect.New(data.Type())
		pointer.Elem().Set(data)
		res.Data = pointer.Interface()
	}

	return res
}

func (self Pointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
