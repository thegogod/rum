package gq

import (
	"encoding/json"
)

type String struct{}

func (self String) Key() string {
	return "string"
}

func (self String) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self String) Resolve(params *ResolveParams) Result {
	switch value := params.Value.(type) {
	case string:
		return Result{Data: value}
	}

	return Result{Error: NewError("", "must be a string")}
}

func (self String) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
