package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type RequestParam interface {
	// Validate() bool
}

type HandlerFunc func(r *ghttp.Request)
