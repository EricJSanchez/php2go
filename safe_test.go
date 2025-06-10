package php2go

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSafeSlice(t *testing.T) {
	// 示例结构体类型
	type MyStruct struct {
		Value int
	}

	safeSlice := NewSafeSlice[MyStruct]()
	gt := NewGoTool(4)
	var ms []MyStruct
	// 模拟并发写入
	for i := 0; i < 20; i++ {
		ms = append(ms, MyStruct{Value: i})
	}
	mss := Slice2Chunk[MyStruct](ms, 2)

	for _, item := range mss {
		gt.Add()
		go func(item []MyStruct) {
			defer gt.Done()
			safeSlice.Append(item...)
			fmt.Println(item)
			time.Sleep(1 * time.Second)
		}(item)
	}
	gt.Wait()
	// 获取并打印切片内容
	result := safeSlice.GetSlice()
	for _, v := range result {
		fmt.Println(v.Value)
	}
}

func TestSafeMap(t *testing.T) {
	// 示例结构体类型
	type MyStruct struct {
		Value int
	}
	safeMap := NewSafeMap[int, MyStruct]()
	// 模拟并发写入
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(i+99, MyStruct{Value: i})
		}(i)
	}
	wg.Wait()
	for k, item := range safeMap.GetMap() {
		fmt.Println(k, item.Value)
	}
}
