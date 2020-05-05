package remote_package

import (
	cm "github.com/easierway/concurrent_map" //cm 别名
	"testing"
)

//vendor 依赖管理 godep glide dep
//brew install glide
//cd pwd glide init
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
