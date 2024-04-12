package utils

import "time"

// 根据时间戳获取时间
func GetDateByTime(cateTime int64) string {
	return time.Unix(cateTime, 0).Format("2006-01-02 15:04:05")
}

// 2020-09-26 15:06:23.000
func GetTimeStringMillisecond() string {
	tmp := time.Now().Format("2006-01-02 15:04:05")
	return tmp + ".000"
}

// 2020-09-26 15:06:23
func GetTimeStringSecond() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 2020-09-26
func GetTimeStringDay() string {
	return time.Now().Format("2006-01-02")
}

// 当月最后一天 2021-03-31
func GetNowMonthOfLastDayString() string {
	da := time.Now()                              //当前时间
	nextMonth := da.AddDate(0, 1, 0)              //月份加一
	LastDay := nextMonth.AddDate(0, 0, -da.Day()) //减去当前的日数,就是本月最后一天
	return LastDay.Format("2006-01-02")
}

// 当月第一一天 2024-01-01
func GetNowMonthOfFirstDayString() string {
	da := time.Now() //当前时间
	return time.Date(da.Year(), da.Month(), 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
}
