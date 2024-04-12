package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

// 将字符串加密成 md5(16进制)
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

// 检查md5明文和密文是否一致
func Md5Check(content, encrypted string) bool {
	encrypted = strings.ToLower(encrypted)
	return strings.EqualFold(String2md5(content), encrypted)
}

// RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, fmt.Sprint(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, fmt.Sprint(rand.Intn(26)+65))
		} else {
			result = append(result, fmt.Sprint(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

// 切割由数字组成的字符串到切片
func StringsSplitToSliceInt(s string, sep string) []int64 {
	if s == "" || sep == "" {
		return []int64{}
	}

	p := fmt.Sprintf(`^(\d+[%s]?)+\d|\d$`, sep)
	match, _ := regexp.MatchString(p, s)
	if !match {
		return []int64{}
	}

	strArr := strings.Split(s, sep)
	intArr := make([]int64, len(strArr))

	for k, v := range strArr {
		intArr[k], _ = strconv.ParseInt(v, 10, 64)
	}

	return intArr
}

// 16进制字符串转为表情
func Emejj(ox string) string {
	i, _ := strconv.ParseInt(ox, 0, 0)
	return string(rune(i))
}

func RelaceEmejj(content string) string {
	var resultStr string
	resultStr = content
	rex := regexp.MustCompile(`{U.*?}`)
	if rex != nil {
		result := rex.FindAllStringSubmatch(resultStr, -1)
		for _, v := range result {
			tmp := strings.ReplaceAll(v[0], "U+", "0x")
			tmp2 := Emejj(tmp[1 : len(tmp)-1])
			resultStr = strings.ReplaceAll(resultStr, v[0], tmp2)
		}
	}
	return resultStr
}

// 获取UUID
func GetUidV4() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

// 对字符串切片去除空值
func TrimSlice(values []string) (ret []string) {
	for _, v := range values {
		if v != "" {
			ret = append(ret, v)
		}
	}
	return
}

// 去除string切片中重复的值
func RemoveDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

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

// Substr 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length || end < 0 {
		return ""
	}

	if end > length {
		return string(rs[start:])
	}
	return string(rs[start:end])
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(s ...string) string {
	sort.Strings(s)
	h := sha1.New()
	h.Write([]byte(strings.Join(s, "")))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(s ...string) string {
	sort.Strings(s)
	h := md5.New()
	h.Write([]byte(strings.Join(s, "")))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func FindDifferenceWithInt(A, B []int) (toDelete []int, toAdd []int) {
	setA := make(map[int]bool)
	setB := make(map[int]bool)
	for _, num := range A {
		setA[num] = true
	}
	for _, num := range B {
		setB[num] = true
	}
	for num := range setA {
		if !setB[num] {
			toDelete = append(toDelete, num)
		}
	}
	for num := range setB {
		if !setA[num] {
			toAdd = append(toAdd, num)
		}
	}

	return toDelete, toAdd
}

// A不在B中 返回C
// B不在A中 返回D
func FindDifferenceWithStr(A, B []string) (toDelete []string, toAdd []string) {
	setA := make(map[string]bool)
	setB := make(map[string]bool)
	for _, num := range A {
		setA[num] = true
	}
	for _, num := range B {
		setB[num] = true
	}

	for num := range setA {
		if !setB[num] {
			toDelete = append(toDelete, num)
		}
	}
	for num := range setB {
		if !setA[num] {
			toAdd = append(toAdd, num)
		}
	}

	return toDelete, toAdd
}
