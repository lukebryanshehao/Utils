package utils

import (
	"math/rand"
	"time"
)

/*
kind
KC_RAND_KIND_NUM   = 0	// 纯数字
KC_RAND_KIND_LOWER = 1	// 小写字母
KC_RAND_KIND_UPPER = 2	// 大写字母
KC_RAND_KIND_ALL   = 3 	// 数字、大小写字母
*/

func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)

	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
func GetRand(result int, groupExpertNumber int) (ids []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < groupExpertNumber; i++ {
		id := r.Intn((result - 10) + 10)
		ids = append(ids, id)
		//查重
		exist := false
		for _, v := range ids {
			if v == id {
				exist = true
				break
			}
		}
		if !exist {
			ids = append(ids, id)
		}
	}
	return ids
}
func GenerateRandomNumber(end int, count int) []int {
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end)
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func GenerateRandomNumber2(end int, count int) []int {
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end)
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}