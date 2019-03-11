package utils

//冒泡排序找出数组中最大或者最小的数据  num=1为最大 num=2为最小
func BubbleSort(sort []float64, num int) float64 {
	for i := 0; i < len(sort)-1; i++ {
		for j := 0; j < len(sort)-1-i; j++ {
			if num == 1 {
				if sort[j] < sort[j+1] {
					temp := sort[j]
					sort[j] = sort[j+1]
					sort[j+1] = temp
				}
			}
			if num == 2 {
				if sort[j] > sort[j+1] {
					temp := sort[j]
					sort[j] = sort[j+1]
					sort[j+1] = temp
				}
			}
		}
	}
	return sort[0]
}
