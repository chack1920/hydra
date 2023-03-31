package components

import (
	"fmt"

	"psbnb.com/greatsun/hydra/components/caches"
	"psbnb.com/greatsun/hydra/components/container"
	"psbnb.com/greatsun/hydra/components/dbs"
	"psbnb.com/greatsun/hydra/components/dlock"
	"psbnb.com/greatsun/hydra/components/http"
	"psbnb.com/greatsun/hydra/components/queues"
	"psbnb.com/greatsun/hydra/components/rpcs"
	"psbnb.com/greatsun/hydra/components/uuid"
	"psbnb.com/greatsun/hydra/context"
	"psbnb.com/greatsun/hydra/global"
	"psbnb.com/greatsun/hydra/registry"

	_ "psbnb.com/greatsun/hydra/components/queues/mq/lmq"
	_ "psbnb.com/greatsun/hydra/components/queues/mq/mqtt"
	_ "psbnb.com/greatsun/hydra/components/queues/mq/redis"
	_ "psbnb.com/greatsun/hydra/components/queues/mq/xmq"
)

// IComponent 组件
type IComponent interface {
	Container() container.IContainer
	RPC() rpcs.IComponentRPC
	Queue() queues.IComponentQueue
	Cache() caches.IComponentCache
	HTTP() http.IComponentHTTPClient
	DB() dbs.IComponentDB
	DLock(name string) (dlock.ILock, error)
	UUID() uuid.UUID
}

// Def 默认组件
var Def IComponent = NewComponent()

// Component 组件
type Component struct {
	c          container.IContainer
	rpc        rpcs.IComponentRPC
	queue      queues.IComponentQueue
	cache      caches.IComponentCache
	db         dbs.IComponentDB
	httpClient http.IComponentHTTPClient
}

// NewComponent 创建组件
func NewComponent() *Component {
	c := &Component{
		c: container.NewContainer(),
	}
	c.rpc = rpcs.NewStandardRPC(c.c)
	c.queue = queues.NewStandardQueue(c.c)
	c.cache = caches.NewStandardCache(c.c)
	c.db = dbs.NewStandardDB(c.c)
	c.httpClient = http.NewStandardHTTPClient(c.c)
	return c
}

// Container 获取Container容器
func (c *Component) Container() container.IContainer {
	return c.c
}

// RPC 获取rpc组件
func (c *Component) RPC() rpcs.IComponentRPC {
	return c.rpc
}

// Queue 获取Queue组件
func (c *Component) Queue() queues.IComponentQueue {
	return c.queue
}

// Cache 获取Queue组件
func (c *Component) Cache() caches.IComponentCache {
	return c.cache
}

// DB 获取DB组件
func (c *Component) DB() dbs.IComponentDB {
	return c.db
}

// HTTP 获取HTTP Client组件
func (c *Component) HTTP() http.IComponentHTTPClient {
	return c.httpClient
}

// DLock 获取分布式鍞
func (c *Component) DLock(name string) (dlock.ILock, error) {
	return dlock.NewLock(registry.Join(global.Def.PlatName, "dlock", name), global.Def.RegistryAddr, context.Current().Log())
}

// UUID 获取全局唯一编号
func (c *Component) UUID() uuid.UUID {
	cluster, err := context.Current().APPConf().GetServerConf().GetCluster()
	if err != nil {
		panic(fmt.Errorf("获取集群信息失败:%w", err))
	}
	id := cluster.Current().GetNodeID()
	return uuid.GetSUUID(id).Get()
}
