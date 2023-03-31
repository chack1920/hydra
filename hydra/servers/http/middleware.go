package http

import (
	"psbnb.com/greatsun/hydra/hydra/servers/pkg/middleware"
	_ "psbnb.com/greatsun/hydra/hydra/servers/pkg/nfs"
)

var middlewares = make(middleware.Handlers, 0, 1)

// Middlewares 用户自定义中间件
var Middlewares middleware.ICustomMiddleware = middlewares
