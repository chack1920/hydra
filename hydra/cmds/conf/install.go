package conf

import (
	logs "github.com/lib4dev/cli/logger"
	"github.com/urfave/cli"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/global/compatible"
	"psbnb.com/greatsun/hydra/hydra/cmds/pkgs"
	"psbnb.com/greatsun/hydra/registry"
)

func installNow(c *cli.Context) (err error) {
	//1. 绑定应用程序参数
	global.Current().Log().Pause()
	if err := global.Def.Bind(c); err != nil {
		cli.ShowCommandHelp(c, c.Command.Name)
		return err
	}

	//2.检查是否安装注册中心配置
	if registry.GetProto(global.Current().GetRegistryAddr()) != registry.LocalMemory {
		if err := pkgs.Pub2Registry(coverIfExists, importConf); err != nil {
			logs.Log.Error("安装到配置中心:", compatible.FAILED)
			return err
		}
		logs.Log.Info("安装到配置中心:", compatible.SUCCESS)
		return
	}

	return nil
}
