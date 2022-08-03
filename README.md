# php2go

go实现php的一些数组方法，go版本1.18+，需支持泛型

返回类型由调用指定，舍弃interface{}

```
type Test struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
    Age  int64  `json:"age"`
}
var sl []Test
sl = append(sl, Test{
    Id:   1,
    Name: "n1",
    Age:  101,
})
sl = append(sl, Test{
    Id:   2,
    Name: "n2",
    Age:  102,
})

rr1 := ArrayColumn[int](sl, "Id")
rr1 := ArrayColumn[string](sl, "Name")
rr2 := ArrayColumn[int64](sl, "Age")

rr3 := ArrayReverse([]int64{1, 2, 3, 4, 5})

rr4 := ArraySum[int](sl, "Id")
rr5 := ArraySum[int64]([]int64{1, 2, 3, 4, 5}, "")

rr6 := InArray([]int64{1, 2, 3, 4, 5}, 6)

rr7 := ArrayIntersect([]int64{1, 2, 3, 4, 5}, []int64{33, 109})

rr8 := ArrayDiff([]int64{1, 2, 3, 4, 5}, []int64{33, 109, 4})

rr9 := ArrayUnique([]string{"a", "b", "c", "a", "b", "d"})
```