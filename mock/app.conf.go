package mock

import (
	"github.com/micro-plat/lib4go/types"
	"github.com/chack1920/hydra"
	"github.com/chack1920/hydra/conf/app"
	"github.com/chack1920/hydra/creator"
	"github.com/chack1920/hydra/global"
	"github.com/chack1920/hydra/hydra/servers/http"
)

// NewAPPConf 构建APP配置
func NewAPPConf(opts ...hydra.Option) (app.IAPPConf, error) {
	for _, opt := range opts {
		opt()
	}
	//初始化参数
	global.Def.PlatName = types.GetString(global.Def.PlatName, "mock_plat")
	global.Def.SysName = types.GetString(global.Def.SysName, "tserver")
	global.Def.ClusterName = types.GetString(global.Def.ClusterName, "test")
	global.Def.RegistryAddr = types.GetString(global.Def.RegistryAddr, "lm://.")
	global.Def.ServerTypes = []string{http.API}

	//发布配置
	err := creator.Conf.Pub(global.Current().GetPlatName(),
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
	return app.Cache.GetAPPConf(http.API)
}
