package rpc

import (
	"github.com/chack1920/hydra/hydra/servers/pkg/middleware"
)

var middlewares = make(middleware.Handlers, 0, 1)

// Middlewares 用户自定义中间件
var Middlewares middleware.ICustomMiddleware = middlewares
