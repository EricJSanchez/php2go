package php2go

import (
	"fmt"
	"sync"
	"testing"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var ps = []Person{
	{
		Id:   1,
		Name: "n1",
		Age:  101,
	}, {
		Id:   2,
		Name: "n2",
		Age:  102,
	}, {
		Id:   3,
		Name: "n3",
		Age:  103,
	},
}

func TestArray(t *testing.T) {
	rr := ArrayColumn[int](ps, "Id")
	fmt.Printf("%T \n", rr)
	fmt.Println(rr)
	rr1 := ArrayColumn[string](ps, "Name")
	fmt.Printf("1 %T:", rr1)
	fmt.Println(rr1)
	rr2 := ArrayColumn[int64](ps, "Age")
	fmt.Printf("2 %T:", rr2)
	fmt.Println(rr2)

	rr3 := ArrayReverse([]int64{1, 2, 3, 4, 5})
	fmt.Printf("3 %T:", rr3)
	fmt.Println(rr3)
	rr4 := ArraySum[int](ps, "Id")
	fmt.Printf("4 %T:", rr4)
	fmt.Println(rr4)
	rr5 := ArraySum[int64]([]int64{1, 2, 3, 4, 5}, "")
	fmt.Printf("5 %T:", rr5)
	fmt.Println(rr5)
	//rr6 := InArray(6, []int{1, 2, 3, 4, 5, 6})
	rr6 := InArray(int64(6), []int64{1, 2, 3, 4, 5, 6})
	fmt.Printf("6 %T:", rr6)
	fmt.Println(rr6)
	rr7 := ArrayIntersect([]int64{1, 2, 3, 4, 5}, []int64{33, 109})
	fmt.Printf("7 %T:", rr7)
	fmt.Println(rr7)
	rr8 := ArrayDiff([]int64{1, 2, 3, 4, 5}, []int64{33, 109, 4})
	fmt.Printf("8 %T:", rr8)
	fmt.Println(rr8)

	rr9 := ArrayUnique([]string{"a", "b", "c", "a", "b", "d"})
	fmt.Printf("9 %T:", rr9)
	fmt.Println(rr9)

	//rr10 := Min([]string{"a", "b", "c", "a", "b", "d", "d3"})
	rr10 := Min([]int{1, 2, 3, 4, 5, 0})
	fmt.Printf("10 %T:", rr10)
	fmt.Println(rr10)

	//rr11 := Max([]string{"a", "b", "c", "a", "b", "d", "d3"})
	rr11 := Min([]int64{})
	fmt.Printf("11 %T:", rr11)
	fmt.Println(rr11)
}

func TestSliceRemove(t *testing.T) {
	rr := SliceRemove[Person](ps, []int{3, 2, 4, 0})
	fmt.Printf("type:%T,val:%v \n", rr, rr)
}

func TestSafeSlice(t *testing.T) {
	// 示例结构体类型
	type MyStruct struct {
		Value int
	}

	safeSlice := NewSafeSlice[MyStruct]()
	var wg sync.WaitGroup
	var ms []MyStruct
	// 模拟并发写入
	for i := 0; i < 20; i++ {
		ms = append(ms, MyStruct{Value: i})
	}
	mss := Slice2Chunk[MyStruct](ms, 2)

	for _, item := range mss {
		wg.Add(1)
		go func(item []MyStruct) {
			defer wg.Done()
			safeSlice.Append(item...)
		}(item)
	}
	wg.Wait()
	// 获取并打印切片内容
	result := safeSlice.GetSlice()
	for _, v := range result {
		fmt.Println(v.Value)
	}
}

func TestSliceCut(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(Slice2Chunk[int](arr, 3))

	arrStr := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(Slice2Chunk[string](arrStr, 2))

	// 示例结构体类型
	type MyStruct struct {
		Value int
	}
	var ms []MyStruct
	for i := 0; i < 10; i++ {
		ms = append(ms, MyStruct{
			Value: i,
		})
	}
	fmt.Println(Slice2Chunk[MyStruct](ms, 2))
}
