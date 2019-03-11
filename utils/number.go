package utils

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//创建数字字符串(不足前面补0)
//CreateNumberStr(4,0) return 0001
//CreateNumberStr(5,1) return 00002
func CreateNumberStr(figures int, startNumber int) (string) {
	numberStr := strconv.Itoa(startNumber + 1)
	for (true) {
		if len(strconv.Itoa(startNumber)) < figures {
			numberStr = "0" + numberStr
		}
		if len(numberStr) >= figures {
			break
		}
	}
	return numberStr
}

//创建多位数长度数字字符串
//figures位数
//rang随机数范围
func CreateRandNum(figures int, rang int) (numberStr string) {

	rand.Seed(time.Now().Unix())
	for (true) {
		rnd := rand.Intn(rang)
		numberStr += strconv.Itoa(rnd)

		if len(numberStr) >= figures {
			break
		}
	}
	return
}

//Float64转string
func GetFloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func TrimHtml(src string) string {
	//清除转换符
	src = strings.Replace(src, "&nbsp;", "", -1)
	src = strings.Replace(src, "&rdquo;", "", -1)
	src = strings.Replace(src, "&ldquo;", "", -1)
	src = strings.Replace(src, "&mdash;", "", -1)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "<zzy>")
	//去除&ldquo
	////去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "<zzy>")
	return strings.TrimSpace(src)
}
//去除字符串中的样式
func TrimHtml2(src string) string {
	//清除转换符
	src = strings.Replace(src, "&nbsp;", "", -1)
	src = strings.Replace(src, "&rdquo;", "", -1)
	src = strings.Replace(src, "&ldquo;", "", -1)
	src = strings.Replace(src, "&mdash;", "", -1)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除&ldquo
	////去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func RemoveRep(s []uint) []uint {
	start := time.Now()
	result := []uint{}
	m := make(map[uint]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	fmt.Println("花费时间:", fmt.Sprintf("%vms", (time.Now().UnixNano()-start.UnixNano())/1e+6))
	return result
}

//float64转int四舍五入
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}
//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is wrong")
	}
	if end < 0 || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}