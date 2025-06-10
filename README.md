# php2go

go实现php的一些数组方法，go版本1.18+，需支持泛型

返回类型由调用指定，舍弃interface{}

```
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

rr1 := ArrayColumn[int](ps, "Id")
rr1 := ArrayColumn[string](ps, "Name")
rr2 := ArrayColumn[int64](ps, "Age")

rr3 := ArrayReverse([]int64{1, 2, 3, 4, 5})

rr4 := ArraySum[int](ps, "Id")
rr5 := ArraySum[int64]([]int64{1, 2, 3, 4, 5}, "")

rr6 := InArray(6, []int64{1, 2, 3, 4, 5})

rr7 := ArrayIntersect([]int64{1, 2, 3, 4, 5}, []int64{33, 109})

rr8 := ArrayDiff([]int64{1, 2, 3, 4, 5}, []int64{33, 109, 4})

rr9 := ArrayUnique([]string{"a", "b", "c", "a", "b", "d"})

rr10 := SliceRemove[Person](ps, []int{3, 2, 4, 0})

rr11 := Slice2Chunk[Person](ps, 2)

gt := NewGoTool(4) //控制并发数量
gt.Add()
gt.Done()
gt.Wait()

safeSlice := NewSafeSlice[MyStruct]()
safeSlice.Append(item...)
safeSlice.GetSlice()

safeMap := NewSafeMap[int, MyStruct]()
safeMap.Set(i, MyStruct{Value: i})
safeMap.GetMap()
```