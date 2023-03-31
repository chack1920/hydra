package http

import (
	"github.com/chack1920/hydra/hydra/servers/pkg/middleware"
	_ "github.com/chack1920/hydra/hydra/servers/pkg/nfs"
)

var middlewares = make(middleware.Handlers, 0, 1)

// Middlewares 用户自定义中间件
var Middlewares middleware.ICustomMiddleware = middlewares
