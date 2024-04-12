package utils

import (
	"reflect"
)

// interface转Slice
func InterfaceToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

// 切片差集交集
func SliceDifference[T comparable](A, B []T) (notInB, notInA, common []T) {
	setA := make(map[T]bool)
	for _, item := range A {
		setA[item] = true
	}
	setB := make(map[T]bool)
	for _, item := range B {
		setB[item] = true
	}

	for _, item := range A {
		if !setB[item] {
			notInB = append(notInB, item)
		} else {
			common = append(common, item)
		}
	}

	for _, item := range B {
		if !setA[item] {
			notInA = append(notInA, item)
		}
	}
	return notInB, notInA, common
}
