package mock

import (
	"github.com/micro-plat/lib4go/types"
	"psbnb.com/greatsun/hydra/conf/app"
	"psbnb.com/greatsun/hydra/context"
	"psbnb.com/greatsun/hydra/context/ctx"
	"psbnb.com/greatsun/hydra/creator"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/hydra/servers/http"
	_ "psbnb.com/greatsun/hydra/registry/registry/localmemory"
	"psbnb.com/greatsun/hydra/services"
)

// NewContext 创建mock类型的Context包
func NewContext(content string, opts ...Option) context.IContext {

	//构建mock
	mk := newMock(content, opts...)

	//初始化参数
	global.Def.PlatName = types.GetString(global.Def.PlatName, "mock_plat")
	global.Def.SysName = types.GetString(global.Def.SysName, "tserver")
	global.Def.ClusterName = types.GetString(global.Def.ClusterName, "test")
	global.Def.RegistryAddr = types.GetString(global.Def.RegistryAddr, "lm://.")
	global.Def.ServerTypes = []string{http.API}

	services.GetRouter(http.API).BuildRouters("")

	if mk.Conf == nil {
		mk.Conf = creator.Conf
	}

	//发布配置
	err := mk.Conf.Pub(global.Current().GetPlatName(),
		global.Current().GetSysName(),
		global.Current().GetClusterName(),
		global.Def.RegistryAddr, nil)
	if err != nil {
		panic(err)
	}

	//初始化缓存
	err = app.PullAndSave()
	if err != nil {
		panic(err)
	}

	//构建Context
	return ctx.NewCtx(mk, global.Def.ServerTypes[0])
}
