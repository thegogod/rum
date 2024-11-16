package middleware

import (
	"time"

	"github.com/thegogod/rum/gq"
)

func Elapse(params *gq.ResolveParams, next gq.Resolver) gq.Result {
	now := time.Now()
	res := next(params)
	res.SetMeta("$elapse", time.Now().Sub(now).Milliseconds())
	return res
}
