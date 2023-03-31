package creator

import (
	"fmt"

	"psbnb.com/greatsun/hydra/conf/server/acl/blacklist"
	"psbnb.com/greatsun/hydra/conf/server/acl/limiter"
	"psbnb.com/greatsun/hydra/conf/server/acl/proxy"
	"psbnb.com/greatsun/hydra/conf/server/acl/whitelist"
	"psbnb.com/greatsun/hydra/conf/server/api"
	"psbnb.com/greatsun/hydra/conf/server/auth/apikey"
	"psbnb.com/greatsun/hydra/conf/server/auth/basic"
	"psbnb.com/greatsun/hydra/conf/server/auth/jwt"
	"psbnb.com/greatsun/hydra/conf/server/auth/ras"
	"psbnb.com/greatsun/hydra/conf/server/header"
	"psbnb.com/greatsun/hydra/conf/server/nfs"
	"psbnb.com/greatsun/hydra/conf/server/processor"
	"psbnb.com/greatsun/hydra/conf/server/render"
	"psbnb.com/greatsun/hydra/conf/server/static"
)

type httpBuilder struct {
	BaseBuilder
	tp string
}

// newHTTP 构建http生成器
func newHTTP(tp string, address string, opts ...api.Option) *httpBuilder {
	b := &httpBuilder{tp: tp, BaseBuilder: make(map[string]interface{})}
	b.BaseBuilder[ServerMainNodeName] = api.New(address, opts...)
	return b
}

// Load 加载路由
func (b *httpBuilder) Load() {
}

// Jwt jwt配置
func (b *httpBuilder) Jwt(opts ...jwt.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", jwt.ParNodeName, jwt.SubNodeName)
	b.BaseBuilder[path] = jwt.NewJWT(opts...)
	return b
}

// Fsa fsa静态密钥错误
func (b *httpBuilder) APIKEY(secret string, opts ...apikey.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", apikey.ParNodeName, apikey.SubNodeName)
	b.BaseBuilder[path] = apikey.New(secret, opts...)
	return b
}

// Fsa fsa静态密钥错误
func (b *httpBuilder) Basic(opts ...basic.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", basic.ParNodeName, basic.SubNodeName)
	b.BaseBuilder[path] = basic.NewBasic(opts...)
	return b
}

// WhiteList 设置白名单
func (b *httpBuilder) WhiteList(opts ...whitelist.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", whitelist.ParNodeName, whitelist.SubNodeName)
	b.BaseBuilder[path] = whitelist.New(opts...)
	return b
}

// BlackList 设置黑名单
func (b *httpBuilder) BlackList(opts ...blacklist.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", blacklist.ParNodeName, blacklist.SubNodeName)
	b.BaseBuilder[path] = blacklist.New(opts...)
	return b
}

// Ras 远程认证服务配置
func (b *httpBuilder) Ras(opts ...ras.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", ras.ParNodeName, ras.SubNodeName)
	b.BaseBuilder[path] = ras.NewRASAuth(opts...)
	return b
}

// Header 头配置
func (b *httpBuilder) Header(opts ...header.Option) *httpBuilder {
	b.BaseBuilder[header.TypeNodeName] = header.New(opts...)
	return b
}

// Static 静态文件配置
func (b *httpBuilder) Static(opts ...static.Option) *httpBuilder {
	b.BaseBuilder[static.TypeNodeName] = static.New(b.tp, opts...)
	return b
}

// Limit 服务器限流配置
func (b *httpBuilder) Limit(opts ...limiter.Option) *httpBuilder {
	path := fmt.Sprintf("%s/%s", limiter.ParNodeName, limiter.SubNodeName)
	b.BaseBuilder[path] = limiter.New(opts...)
	return b
}

// Proxy 代理配置
func (b *httpBuilder) Proxy(script string) *httpBuilder {
	path := fmt.Sprintf("%s/%s", proxy.ParNodeName, proxy.SubNodeName)
	b.BaseBuilder[path] = script
	return b
}

// Render 响应渲染配置
func (b *httpBuilder) Render(script string) *httpBuilder {
	b.BaseBuilder[render.TypeNodeName] = script
	return b
}

// Processor 构建Processor配置
func (b *httpBuilder) Processor(opts ...processor.Option) *httpBuilder {
	b.BaseBuilder[processor.TypeNodeName] = processor.New(opts...)
	return b
}

// NFS 网络文件系统配置
func (b *httpBuilder) NFS(local string, opts ...nfs.Option) *httpBuilder {
	b.BaseBuilder[nfs.TypeNodeName] = nfs.New(local, opts...)
	return b
}
