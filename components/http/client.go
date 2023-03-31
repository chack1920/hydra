package http

import (
	"github.com/micro-plat/lib4go/types"
	"github.com/chack1920/hydra/components/container"
	"github.com/chack1920/hydra/components/pkgs/http"
	"github.com/chack1920/hydra/conf"
	httpconf "github.com/chack1920/hydra/conf/vars/http"
)

// StandardHTTPClient db
type StandardHTTPClient struct {
	c container.IContainer
}

// NewStandardHTTPClient 创建DB
func NewStandardHTTPClient(c container.IContainer) *StandardHTTPClient {
	return &StandardHTTPClient{c: c}
}

// GetRegularClient 获取正式的没有异常数据库实例
func (s *StandardHTTPClient) GetRegularClient(names ...string) (d IClient) {
	d, err := s.GetClient(names...)
	if err != nil {
		panic(err)
	}
	return d
}

// GetClient 获取http请求对象
func (s *StandardHTTPClient) GetClient(names ...string) (d IClient, err error) {
	name := types.GetStringByIndex(names, 0, httpconf.HttpNameNode)
	obj, err := s.c.GetOrCreate(httpconf.HttpTypeNode, name, func(conf *conf.RawConf, keys ...string) (interface{}, error) {
		if conf.IsEmpty() {
			return http.NewClient()
		}
		return http.NewClient(httpconf.WithRaw(conf.GetRaw()))
	})
	if err != nil {
		return nil, err
	}
	return obj.(IClient), nil
}
