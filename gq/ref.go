package gq

import (
	"encoding/json"
)

type Ref struct {
	resolve func() Schema
}

func (self Ref) Key() string {
	return self.resolve().Key()
}

func (self Ref) Do(params *DoParams) Result {
	return self.resolve().Do(params)
}

func (self Ref) Resolve(params *ResolveParams) Result {
	return self.resolve().Resolve(params)
}

func (self Ref) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.Key())
}
