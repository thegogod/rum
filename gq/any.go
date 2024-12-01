package gq

import (
	"encoding/json"
)

type Any struct{}

func (self Any) Key() string {
	return "any"
}

func (self Any) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self Any) Resolve(params *ResolveParams) Result {
	return Result{Data: params.Value}
}

func (self Any) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
