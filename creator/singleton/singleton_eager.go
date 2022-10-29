package singleton

import "fmt"

// eagerSingleton 饿汉实例,小写开头,保证只能本包访问
type eagerSingleton struct {
}

func newEagerSingleton() *eagerSingleton {
	fmt.Println("eagerSingleton init...")
	return &eagerSingleton{}
}

// 全局变量,加载时初始化,无需考虑并发问题
var eagerSingletonInstance = newEagerSingleton()

// GetEagerSingletonInstance
// @Description: 对外部暴露的接口
// @return *eagerSingleton
func GetEagerSingletonInstance() *eagerSingleton {
	return eagerSingletonInstance
}
