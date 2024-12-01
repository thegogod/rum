package gq

import (
	"encoding/json"
	"time"
)

type Date struct{}

func (self Date) Key() string {
	return "Date"
}

func (self Date) Do(params *DoParams) Result {
	return self.Resolve(&ResolveParams{
		Value:   params.Value,
		Context: params.Context,
	})
}

func (self Date) Resolve(params *ResolveParams) Result {
	switch value := params.Value.(type) {
	case time.Time:
		return Result{Data: value}
	case *time.Time:
		return Result{Data: value}
	}

	return Result{Error: NewError("", "must be a Date")}
}

func (self Date) MarshalJSON() ([]byte, error) {
	return json.Marshal("Date")
}
