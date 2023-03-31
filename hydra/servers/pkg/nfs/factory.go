package nfs

import (
	"fmt"
	"strings"

	"psbnb.com/greatsun/hydra/conf/app"
	"psbnb.com/greatsun/hydra/conf/server/nfs"
	"psbnb.com/greatsun/hydra/hydra/servers/pkg/nfs/infs"
	"psbnb.com/greatsun/hydra/hydra/servers/pkg/nfs/lnfs"
	"psbnb.com/greatsun/hydra/hydra/servers/pkg/nfs/obs"
)

func getNFS(app app.IAPPConf, c *nfs.NFS) infs.Infs {
	switch strings.ToUpper(c.CloudNFS) {
	case "OBS":
		return obs.NewOBS(c.AccessKey, c.SecretKey, c.Local, c.Endpoint, c.Excludes, c.Includes...)
	case "":
		return lnfs.NewNFS(app, c)
	default:
		panic(fmt.Sprintf("不支持的NFS类型:%s", c.CloudNFS))
	}

}
