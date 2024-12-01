package gq

import (
	"encoding/json"
)

type Bool struct{}

func (self Bool) Key() string {
	return "bool"
}

func (self Bool) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self Bool) Resolve(params *ResolveParams) Result {
	switch value := params.Value.(type) {
	case bool:
		return Result{Data: value}
	}

	return Result{Error: NewError(params.Key, "must be a boolean")}
}

func (self Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
