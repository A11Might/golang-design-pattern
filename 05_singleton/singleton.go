package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	instance     *singleton
	instanceOnce sync.Once
)

func GetInstance() *singleton {
	instanceOnce.Do(func() {
		fmt.Println("生成了一个实例")
		instance = new(singleton)
	})
	return instance
}
