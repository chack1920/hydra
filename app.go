package hydra

import (
	"github.com/lib4dev/cli"
	"github.com/micro-plat/lib4go/logger"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/global/compatible"
	"psbnb.com/greatsun/hydra/services"

	_ "psbnb.com/greatsun/hydra/registry/watcher/wchild"
	_ "psbnb.com/greatsun/hydra/registry/watcher/wvalue"

	_ "psbnb.com/greatsun/hydra/hydra/cmds/conf"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/install"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/remove"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/run"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/update"

	_ "psbnb.com/greatsun/hydra/hydra/cmds/start"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/status"
	_ "psbnb.com/greatsun/hydra/hydra/cmds/stop"

	_ "psbnb.com/greatsun/hydra/hydra/cmds/restart"

	_ "psbnb.com/greatsun/hydra/registry/registry/dbr"
	_ "psbnb.com/greatsun/hydra/registry/registry/filesystem"
	_ "psbnb.com/greatsun/hydra/registry/registry/localmemory"
	_ "psbnb.com/greatsun/hydra/registry/registry/redis"
	_ "psbnb.com/greatsun/hydra/registry/registry/zookeeper"
)

// MicroApp  微服务应用
type MicroApp struct {
	app *cli.App
	services.IService
}

// NewApp 创建微服务应用
func NewApp(opts ...Option) (m *MicroApp) {
	m = &MicroApp{
		IService: services.Def,
	}
	for _, opt := range opts {
		opt()
	}
	return m
}

// Start 启动服务器
func (m *MicroApp) Start() {
	defer logger.Close()
	m.app = cli.New(cli.WithVersion(global.Version), cli.WithUsage(global.Usage))
	m.app.Start()
}

// Close 关闭服务器
func (m *MicroApp) Close() {
	Close()
}

// Close 关闭服务器
func Close() {
	compatible.AppClose()
}
