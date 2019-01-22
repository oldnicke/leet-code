package main

import "fmt"

func main() {
	fmt.Println(subarrayBitwiseORs([]int{0}))
	fmt.Println(subarrayBitwiseORs([]int{1, 1, 2}))
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
}

func subarrayBitwiseORs(A []int) int {
	resMap := map[int]bool{}
	curMap := map[int]bool{}

	// 起始位置
	for i := 0; i < len(A); i++ {
		curNew := map[int]bool{}
		for v := range curMap {
			newV := v | A[i]
			if resMap[newV] == false {
				resMap[newV] = true
			}
			curNew[newV] = true
		}
		curNew[A[i]] = true
		if resMap[A[i]] == false {
			resMap[A[i]] = true
		}
		curMap = curNew
	}
	return len(resMap)
}
