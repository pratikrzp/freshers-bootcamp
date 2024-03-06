package examples

import "fmt"

func ArraySlice() {
	dummyArr := []int{}

	fmt.Println(dummyArr)

	dummyArr = append(dummyArr, 1)
	fmt.Println(dummyArr)

	for i := 0; i < 5; i++ {
		dummyArr = append(dummyArr, i)
	}

	fmt.Println(dummyArr)

	for _, num := range dummyArr {
		fmt.Println(num)
	}
}
