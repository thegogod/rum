package gq

import (
	"encoding/json"
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
	switch value := params.Value.(type) {
	case int:
		return Result{Data: value}
	case *int:
		return Result{Data: value}
	}

	return Result{Error: NewError("", "must be an integer")}
}

func (self Int) MarshalJSON() ([]byte, error) {
	return json.Marshal("int")
}
