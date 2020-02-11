package synch_lock

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var (
	pool   *KeyLockPool
	values map[string]int
)

func TestNewKeyLockPool(t *testing.T) {

	pool = NewKeyLockPool(10)
	values = map[string]int{}

	var (
		keys = make([]string, 1000)
	)

	for i := 0; i < 1000; i++ {
		n := rand.Intn(20)
		keys[i] = fmt.Sprint(n)
	}

	fmt.Printf("keys: %v", strings.Join(keys, ","))

	// 统计匹配
	retCheck := map[string]int{}
	for _, key := range keys {
		retCheck[key] += 1
		go runValue(key)
	}

	for {
		time.Sleep(100)
		if pool.IsAllFree() {
			break
		}
	}

	for key, value := range values {
		//fmt.Printf("Key %v 期待值%v， 实际%v/n", key,  retCheck[key], value)
		if ret, exist := retCheck[key]; !exist || ret != value {
			t.Errorf("Error: Key %v 期待值%v， 实际%v", key, ret, value)
		}
	}

}

func runValue(key string) {
	if pool.GetLock(key) {
		if _, exist := values[key]; !exist {
			values[key] = 0
			//fmt.Printf("[INIT] key: %v, value: 0\n", key)
		}
		values[key] += 1
		//fmt.Printf("[Add] key: %v, value: %v\n", key, values[key])
		time.Sleep(100)
		pool.ReleaseLock(key)
		//fmt.Printf("[Release] key: %v, value: %v\n", key, values[key])
	} else {
		//fmt.Printf("[Block] key:%v, keys:%v\n", key, pool.GetKeysStatus())
		time.Sleep(100)
		//fmt.Printf("[Recalll] key:%v, keys:%v\n", key, pool.GetKeysStatus())
		go runValue(key)
	}
}
