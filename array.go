package php2go

import (
	"reflect"
)

type Integer interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

type IntegerString interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

// ArrayReverse 数组翻转
func ArrayReverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

// ArrayColumn 获取二维数组中的某一列，组合成一维数组
func ArrayColumn[T any](s interface{}, col string) (retCol []T) {
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

// InArray 值是否存在切片中，只限一维数组
func InArray[T IntegerString](val T, s []T) bool {
	for _, item := range s {
		if val == item {
			return true
		}
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
