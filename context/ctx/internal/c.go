package internal

import (
	"github.com/chack1920/hydra/components"
	"github.com/chack1920/hydra/context"
	"github.com/chack1920/hydra/pkgs"
)

// CallRPC RPC调用
func CallRPC(ctx context.IContext, service string) *pkgs.Rspns {
	response, err := components.Def.RPC().GetRegularRPC().Swap(service, ctx)
	if err != nil {
		ctx.Log().Errorf("调用RPC服务出错:%+v", err)
		return pkgs.NewRspns(err)
	}
	return response
}
