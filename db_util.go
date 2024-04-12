package utils

import "fmt"

// 创建golang oracle数据库连接字符串
func NewOracleConnectionString(name, uname, pwd, host string, port int64, dbname string) string {
	return fmt.Sprintf("oracle://%s:%s@%s:%d/%s", uname, pwd, host, port, dbname)
}

// 创建golang oracle数据库连接字符串
func NewMysqlConnectionString(name, uname, pwd, host string, port int64, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", uname, pwd, host, port, dbname)
}
