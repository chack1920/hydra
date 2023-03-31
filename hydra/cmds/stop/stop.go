package stop

import (
	"github.com/lib4dev/cli/cmds"
	"psbnb.com/greatsun/hydra/global"

	"github.com/urfave/cli"
	"psbnb.com/greatsun/hydra/hydra/cmds/pkgs"
)

var isFixed bool

func init() {
	cmds.RegisterFunc(func() cli.Command {
		return cli.Command{
			Name:   "stop",
			Usage:  "停止服务，停止服务器运行",
			Flags:  pkgs.GetFixedFlags(&isFixed),
			Action: doStop,
		}
	})
}

func doStop(c *cli.Context) (err error) {
	//关闭日志显示
	global.Current().Log().Pause()
	//3.创建本地服务
	hydraSrv, err := pkgs.GetService(c, isFixed)
	if err != nil {
		return err
	}

	err = hydraSrv.Stop()
	return pkgs.GetCmdsResult(hydraSrv.DisplayName, "Stop", err)
}
