package gq

import (
	"encoding/json"
	"reflect"
)

type Int struct{}

func (self Int) Key() string {
	return "int"
}

func (self Int) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self Int) Resolve(params *ResolveParams) Result {
	value := reflect.ValueOf(params.Value)

	if value.IsValid() && value.CanInt() {
		return Result{Data: value.Interface()}
	}

	return Result{Error: NewError("", "must be an integer")}
}

func (self Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
