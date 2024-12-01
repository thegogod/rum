package gq

import (
	"encoding/json"
)

type Float struct{}

func (self Float) Key() string {
	return "float"
}

func (self Float) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self Float) Resolve(params *ResolveParams) Result {
	switch value := params.Value.(type) {
	case float32:
		return Result{Data: value}
	case float64:
		return Result{Data: value}
	}

	return Result{Error: NewError(params.Key, "must be a float")}
}

func (self Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
