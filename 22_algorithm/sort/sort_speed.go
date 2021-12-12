package algorithm

import (
	"fmt"
	"time"
)

func Sort() {
	start := time.Now().UnixNano()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		bubbleSort(&arr)
	}
	//打印消耗时间
	fmt.Println(time.Now().UnixNano() - start)
}

//排序
func bubbleSort(arr *[]int) {
	for j := 0; j < len(*arr)-1; j++ {
		for k := 0; k < len(*arr)-1-j; k++ {
			if (*arr)[k] < (*arr)[k+1] {
				temp := (*arr)[k]
				(*arr)[k] = (*arr)[k+1]
				(*arr)[k+1] = temp
			}
		}
	}
}
