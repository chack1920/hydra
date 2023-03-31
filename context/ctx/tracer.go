package ctx

import (
	"github.com/micro-plat/lib4go/logger"
	"github.com/chack1920/hydra/conf/app"
	"github.com/chack1920/hydra/context"
	"github.com/chack1920/hydra/context/ctx/internal"
)

type tracer struct {
	*internal.Tracer
	l logger.ILogger
}

func newTracer(path string, l logger.ILogger, c app.IAPPConf) *tracer {
	return &tracer{
		Tracer: internal.Empty,
		l:      l,
	}
}

// Root 根节点
func (t *tracer) Root() context.ITraceSpan {
	return t.Tracer.Root()
}
