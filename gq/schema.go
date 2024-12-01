package gq

import (
	"context"
)

type DoParams struct {
	Query   string          `json:"query"`
	Value   any             `json:"value,omitempty"`
	Context context.Context `json:"context,omitempty"`
}

type ResolveParams struct {
	Query   Query           `json:"query"`
	Parent  any             `json:"parent,omitempty"`
	Key     string          `json:"key,omitempty"`
	Value   any             `json:"value,omitempty"`
	Context context.Context `json:"context,omitempty"`
}

type Schema interface {
	Key() string
	Do(params *DoParams) Result
	Resolve(params *ResolveParams) Result
}
