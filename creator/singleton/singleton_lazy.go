package singleton

import (
	"fmt"
	"sync"
)

// lazySingleton
// @Description: 懒汉实例
type lazySingleton struct {
}

func newLazySingleton() *lazySingleton {
	fmt.Println("lazySingleton init...")
	return &lazySingleton{}
}

var (
	lazySingletonDCLInstance  *lazySingleton
	lazySingletonOnceInstance *lazySingleton
	lock                      sync.Mutex
	once                      sync.Once
)

// GetLazySingletonInstanceWithDCL
// @Description: 使用DCL(双重检测锁)保证并发情况下也只会创建一个实例
// @return *lazySingleton
func GetLazySingletonInstanceWithDCL() *lazySingleton {
	if lazySingletonDCLInstance == nil {
		lock.Lock()
		if lazySingletonDCLInstance == nil {
			lazySingletonDCLInstance = newLazySingleton()
		}
		lock.Unlock()
	}
	return lazySingletonDCLInstance
}

// GetLazySingletonInstanceWithOnce
// @Description: 使用once并发原语实现
// @return *lazySingleton
func GetLazySingletonInstanceWithOnce() *lazySingleton {
	once.Do(func() {
		lazySingletonOnceInstance = newLazySingleton()
	})
	return lazySingletonOnceInstance
}
