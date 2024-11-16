package assert

import "reflect"

type Rule struct {
	Key     string `json:"key"`
	Value   any    `json:"value"`
	Message string `json:"message"`
	Resolve RuleFn `json:"-"`
}

type RuleFn func(value reflect.Value) (any, error)
