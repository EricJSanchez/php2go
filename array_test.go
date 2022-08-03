package php2go

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
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

	rr := ArrayColumn[int](sl, "Id")
	fmt.Printf("%T \n", rr)
	fmt.Println(rr)
	rr1 := ArrayColumn[string](sl, "Name")
	fmt.Printf("1 %T \n", rr1)
	fmt.Println(rr1)
	rr2 := ArrayColumn[int64](sl, "Age")
	fmt.Printf("2 %T \n", rr2)
	fmt.Println(rr2)

	rr3 := ArrayReverse([]int64{1, 2, 3, 4, 5})
	fmt.Printf("3 %T \n", rr3)
	fmt.Println(rr3)
	rr4 := ArraySum[int](sl, "Id")
	fmt.Printf("4 %T \n", rr4)
	fmt.Println(rr4)
	rr5 := ArraySum[int64]([]int64{1, 2, 3, 4, 5}, "")
	fmt.Printf("5 %T \n", rr5)
	fmt.Println(rr5)
	rr6 := InArray(6, []int64{1, 2, 3, 4, 5})
	fmt.Printf("6 %T \n", rr6)
	fmt.Println(rr6)
	rr7 := ArrayIntersect([]int64{1, 2, 3, 4, 5}, []int64{33, 109})
	fmt.Printf("7 %T \n", rr7)
	fmt.Println(rr7)
	rr8 := ArrayDiff([]int64{1, 2, 3, 4, 5}, []int64{33, 109, 4})
	fmt.Printf("8 %T \n", rr8)
	fmt.Println(rr8)

	rr9 := ArrayUnique([]string{"a", "b", "c", "a", "b", "d"})
	fmt.Printf("9 %T \n", rr9)
	fmt.Println(rr9)
}
