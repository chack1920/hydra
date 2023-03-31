package pkgs

import (
	logs "github.com/lib4dev/cli/logger"

	"github.com/micro-plat/lib4go/logger"
	"github.com/urfave/cli"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/hydra/cmds/pkgs/rlog"
	"psbnb.com/greatsun/hydra/hydra/servers"
	"psbnb.com/greatsun/hydra/registry"
)

func (p *ServiceApp) run() (err error) {
	if p.c.Bool("nostd") {
		logger.RemoveStdoutAppender()
	}

	//1. 绑定应用程序参数
	if err := global.Def.Bind(p.c); err != nil {
		cli.ShowCommandHelp(p.c, p.c.Command.Name)
		return err
	}
	if !global.IsDebug {
		logger.AddWriteThread(99)
	}

	//2. 注册远程日志组件
	if err := rlog.Registry(global.Def.PlatName, global.Def.RegistryAddr); err != nil {
		logs.Log.Error(err)
		return err
	}

	globalData := global.Current()

	//3.创建trace性能跟踪
	p.trace, err = startTrace(globalData.GetTrace(), globalData.GetTracePort())
	if err != nil {
		return err
	}

	//4. 处理本地内存作为注册中心的服务发布问题
	if registry.GetProto(globalData.GetRegistryAddr()) == registry.LocalMemory {
		if err := Pub2Registry(true, p.c.String("import")); err != nil {
			return err
		}
	}

	//5. 创建服务器
	p.server = servers.NewRspServers(globalData.GetRegistryAddr(),
		globalData.GetPlatName(), globalData.GetSysName(),
		globalData.GetServerTypes(), globalData.GetClusterName())
	if err := p.server.Start(); err != nil {
		return err
	}
	return nil
}
