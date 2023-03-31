package restart

import (
	"github.com/lib4dev/cli/cmds"
	"github.com/urfave/cli"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/hydra/cmds/pkgs"
)

var isFixed bool

func init() {
	cmds.RegisterFunc(func() cli.Command {
		return cli.Command{
			Name:   "restart",
			Usage:  "重启服务",
			Action: doRestart,
		}
	})
}

func doRestart(c *cli.Context) (err error) {

	//关闭日志显示
	global.Current().Log().Pause()
	//3.创建本地服务
	hydraSrv, err := pkgs.GetService(c, isFixed)
	if err != nil {
		return err
	}
	err = hydraSrv.Restart()
	return pkgs.GetCmdsResult(hydraSrv.DisplayName, "Restart", err)
}
