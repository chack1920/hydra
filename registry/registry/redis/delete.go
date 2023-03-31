package redis

import (
	"fmt"

	"psbnb.com/greatsun/hydra/registry/registry/redis/internal"
)

// Delete 删除节点
func (r *Redis) Delete(path string) error {
	key := internal.SwapKey(path)
	_, err := r.client.Del(key).Result()
	if err != nil {
		return fmt.Errorf("%v(%s)", err, path)
	}
	r.notifyParentChange(path, 0)
	return nil
}
