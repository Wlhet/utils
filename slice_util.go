package utils

import (
	"fmt"
	"reflect"
)

// interface转Slice
func ToSlice(arr interface{}) []interface{} {
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

// 创建golang oracle数据库连接字符串
func NewOracleConnectionString(name, uname, pwd, host string, port int64, dbname string) string {
	return fmt.Sprintf("oracle://%s:%s@%s:%d/%s", uname, pwd, host, port, dbname)
}

// 创建golang oracle数据库连接字符串
func NewMysqlConnectionString(name, uname, pwd, host string, port int64, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", uname, pwd, host, port, dbname)
}
