package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	xx := (x ^ x>>31) - x>>31
	result := []int{}

	for xx != 0 {
		result = append(result, xx%10)
		xx = xx / 10
	}
	fmt.Println(result)

	v := 0
	k := 1

	for i := len(result) - 1; i >= 0; i-- {
		v += result[i] * k
		k *= 10
	}

	if x < 0 {
		v = ^v + 1
	}

	INT_MAX := int(^uint(0) >> 1)

	if v > INT_MAX {
		v = 0
	} else if v < ^INT_MAX {
		v = 0
	}

	return v
}

func main() {
	fmt.Printf("%d\n", reverse(123))
	fmt.Printf("%d\n", reverse(-123))
}
