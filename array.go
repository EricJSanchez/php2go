package php2go

import (
	"math"
	"reflect"
	"sort"
	"sync"
)

type Integer interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

type IntegerString interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

// ArrayReverse 数组翻转
func ArrayReverse[T IntegerString](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

// ArrayColumn 获取二维数组中的某一列，组合成一维数组
func ArrayColumn[T IntegerString](s interface{}, col string) (retCol []T) {
	rv := reflect.ValueOf(s)
	ln := rv.Len()
	for i := 0; i < ln; i++ {
		tmpRv := rv.Index(i)
		value := tmpRv.FieldByName(col)
		if !value.IsValid() {
			panic("unknown field: " + col)
		}
		switch value.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.String:
			value := value.Interface().(T)
			retCol = append(retCol, value)
		default:
			panic("unknown field: " + col)
			break
		}
	}
	return
}

// ArraySum 计算数组中某一列的和，一维数组时计算全部
func ArraySum[T Integer](s interface{}, col string) (sum T) {
	rv := reflect.ValueOf(s)
	ln := rv.Len()
	for i := 0; i < ln; i++ {
		tmpRv := rv.Index(i)
		// 一维数组，直接相加即可
		if col == "" {
			sum = sum + tmpRv.Interface().(T)
			continue
		}
		value := tmpRv.FieldByName(col)
		if !value.IsValid() {
			panic("unknown field: " + col)
			break
		}
		switch value.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			value := value.Interface().(T)
			sum = sum + value
		default:
			panic("unknown field: " + col)
			break
		}
	}
	return
}

// ArrayUnique 数组去重，只限一维数组
func ArrayUnique[T IntegerString](s []T) (retData []T) {
	tmpMap := make(map[T]bool, len(s))
	for i := 0; i < len(s); i++ {
		// map去重
		if _, ok := tmpMap[s[i]]; ok != true {
			tmpMap[s[i]] = true
			retData = append(retData, s[i])
		}
	}
	return
}

// InArray 值是否存在
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

// ArrayIntersect 数组交集，只限一维数组
func ArrayIntersect[T IntegerString](s1, s2 []T) (intersectArr []T) {
	mp := make(map[T]bool)
	for _, s := range s1 {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range s2 {
		if _, ok := mp[s]; ok {
			intersectArr = append(intersectArr, s)
		}
	}
	return
}

// ArrayDiff 求差集，s1-s2 = s1去除s2在s1中存在的值,只限一维数组
func ArrayDiff[T IntegerString](s1, s2 []T) (diffArr []T) {
	temp := make(map[T]bool, len(s2))

	for _, val := range s2 {
		if _, ok := temp[val]; !ok {
			temp[val] = true
		}
	}

	for _, val := range s1 {
		if _, ok := temp[val]; !ok {
			diffArr = append(diffArr, val)
		}
	}
	return
}

// Max 求最大值
func Max[T IntegerString](s1 []T) (retVal T) {
	if len(s1) == 0 {
		return
	}
	retVal = s1[0]
	for _, val := range s1 {
		if val > retVal {
			retVal = val
		}
	}
	return
}

// Min 求最小值
func Min[T IntegerString](s1 []T) (retVal T) {
	if len(s1) == 0 {
		return
	}
	retVal = s1[0]
	for _, val := range s1 {
		if val < retVal {
			retVal = val
		}
	}
	return
}

// SliceRemove 根据index删除值
func SliceRemove[T any](ori []T, idx []int, flag ...int) (ret []T) {
	defer func() {
		if r := recover(); r != nil {
			ret = ori
			return
		}
	}()
	if len(idx) == 0 {
		ret = ori
		return
	}
	if len(flag) == 0 {
		sort.IntSlice(idx).Sort()
	} else {
		for ii, _ := range idx {
			idx[ii] = idx[ii] - 1
		}
	}
	return SliceRemove(append(ori[:idx[0]], ori[idx[0]+1:]...), idx[1:], 1)
}

// Slice2Chunk 切片按传入的数量切割成二维切片
func Slice2Chunk[T any](ori []T, chunkSize int) (ret [][]T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	total := len(ori)
	j := math.Ceil(float64(total) / float64(chunkSize))
	var left, right int
	for i := 0; i < int(j); i++ {
		if left+chunkSize > total {
			right = total
		} else {
			right = left + chunkSize
		}
		cutSlice := ori[left:right]
		ret = append(ret, cutSlice)
		left = right
	}
	return
}

type GoTool struct {
	ch chan int
	wg sync.WaitGroup
}

func NewGoTool(num int) *GoTool {
	return &GoTool{
		ch: make(chan int, num),
		wg: sync.WaitGroup{},
	}
}

func (gt *GoTool) Add() {
	gt.ch <- 1
	gt.wg.Add(1)
}

func (gt *GoTool) Done() {
	<-gt.ch
	gt.wg.Done()
}

func (gt *GoTool) Wait() {
	gt.wg.Wait()
	close(gt.ch)
}

// SafeSlice 是一个线程安全的切片封装
type SafeSlice[T any] struct {
	mu    sync.Mutex
	slice []T
}

// NewSafeSlice 创建并返回一个新的 SafeSlice 实例
func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{
		slice: make([]T, 0),
	}
}

// Append 添加元素到切片中，确保线程安全
func (ss *SafeSlice[T]) Append(value ...T) {
	ss.mu.Lock()         // 加锁
	defer ss.mu.Unlock() // 函数结束时解锁
	ss.slice = append(ss.slice, value...)
}

// GetSlice 返回当前切片的副本，确保线程安全
func (ss *SafeSlice[T]) GetSlice() []T {
	ss.mu.Lock()         // 加锁
	defer ss.mu.Unlock() // 函数结束时解锁
	// 返回切片的副本以避免外部修改
	newSlice := make([]T, len(ss.slice))
	copy(newSlice, ss.slice)
	return newSlice
}
