package php2go

import (
	"fmt"
	"reflect"
	"sort"
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
			fmt.Println("unknown field:", col)
			return
		}
		switch value.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.String:
			value := value.Interface().(T)
			retCol = append(retCol, value)
		default:
			fmt.Println("unknown case:", value.Kind())
			return
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
			return
		}
		switch value.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			value := value.Interface().(T)
			sum = sum + value
		default:
			return
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
		return false
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

// Slice2Chunk 将切片按指定大小分块
func Slice2Chunk[T any](ori []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}

	total := len(ori)
	chunkCount := (total + chunkSize - 1) / chunkSize // 计算分块数
	result := make([][]T, 0, chunkCount)              // 预分配结果切片

	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}
		result = append(result, ori[i:end])
	}

	return result
}
