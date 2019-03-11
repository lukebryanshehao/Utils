package utils

import (
	"github.com/agnivade/levenshtein"
)

//Levenshtein算法
func GetSimilarity(str1,str2 string)float64{
	leftNum:=GetEnglishOrSymbolChineseSize(str1)+GetEnglishOrSymbolChineseSize(str2)
	distance := levenshtein.ComputeDistance(str1, str2)
	right:=distance
	similarity:=(float64(leftNum)-float64(right))/float64(leftNum)
	similarity = similarity*100
	return similarity
}
//提取字符串中汉字的长度/3
func GetChineseSize(str string)int{
	r := []rune(str)
	size := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			size = size + string(r[i])
		}
	}
	return len(size)
}
//提取字符串中英汉或符号的长度
func GetEnglishOrSymbol(str string)int{
	size:=len(str)-GetChineseSize(str)//字符串总长度减去汉字长度
	return size
}
//获取提取完后字符串后的长度
func  GetEnglishOrSymbolChineseSize(str string)int{
	chineseSize:=GetChineseSize(str)/3
	return chineseSize+GetEnglishOrSymbol(str)
}
