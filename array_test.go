package php2go

import (
	"fmt"
	"testing"
	"time"
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

func TestSliceCut(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(Slice2Chunk[int](arr, 20))

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

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Profile struct {
	User    User     `json:"user"`
	Address *Address `json:"address"`
	Tags    []string `json:"tags"`
}
type Score struct {
	Subject string `json:"subject"`
	Points  int    `json:"points"`
}

type Student struct {
	Name   string  `json:"name"`
	Scores []Score `json:"scores"` // 结构体包含切片字段
}

func TestStructToMap(t *testing.T) {
	// 测试基本结构体
	t.Run("BasicStruct", func(t *testing.T) {
		user := User{
			Name:     "张三",
			Age:      25,
			Birthday: time.Date(1998, 5, 10, 0, 0, 0, 0, time.UTC),
		}

		result, err := Struct2Map(user)
		if err != nil {
			t.Fatalf("转换失败: %v", err)
		}
		Pr(result)
		if result["name"] != "张三" || result["age"] != 25 {
			t.Errorf("基本结构体转换错误")
		}
	})

	// 测试嵌套结构体
	t.Run("NestedStruct", func(t *testing.T) {
		profile := Profile{
			User: User{
				Name: "李四",
				Age:  30,
			},
			Address: &Address{
				City:    "北京",
				Country: "中国",
			},
			Tags: []string{"golang", "backend"},
		}

		result, err := Struct2Map(profile)
		if err != nil {
			t.Fatalf("转换失败: %v", err)
		}
		Pr(result)
		if userMap, ok := result["user"].(map[string]interface{}); !ok || userMap["name"] != "李四" {
			t.Errorf("嵌套结构体转换错误")
		}
	})

	t.Run("BasicSliceStruct", func(t *testing.T) {
		student := Student{
			Name: "王五",
			Scores: []Score{
				{Subject: "数学", Points: 90},
				{Subject: "语文", Points: 85},
			},
		}

		result, err := Struct2Map(student)
		if err != nil {
			t.Fatalf("转换失败: %v", err)
		}
		Pr(result)
		scores, ok := result["scores"].([]interface{})
		if !ok || len(scores) != 2 {
			t.Fatal("切片结构体转换失败")
		}
		Pr(scores)
		firstScore := scores[0].(map[string]interface{})
		if firstScore["subject"] != "数学" {
			t.Error("切片元素字段解析错误")
		}
	})

	t.Run("EmptySlice", func(t *testing.T) {
		student := Student{
			Name:   "赵六",
			Scores: []Score{}, // 空切片测试
		}

		result, err := Struct2Map(student)
		if err != nil {
			t.Fatal(err)
		}
		Pr(result)
		if scores, ok := result["scores"].([]interface{}); !ok || len(scores) != 0 {
			t.Error("空切片处理错误")
		}
	})

	// 测试指针结构体
	t.Run("PointerStruct", func(t *testing.T) {
		addr := &Address{
			City:    "上海",
			Country: "中国",
		}

		result, err := Struct2Map(addr)
		if err != nil {
			t.Fatalf("转换失败: %v", err)
		}

		fmt.Println(result)
		if result["city"] != "上海" {
			t.Errorf("指针结构体转换错误")
		}
	})

	// 测试非结构体输入
	t.Run("NonStructInput", func(t *testing.T) {
		result, err := Struct2Map("not a struct")
		Pr(result)
		if err == nil {
			t.Errorf("非结构体输入应该返回错误")
		}
	})
}
