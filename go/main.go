package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	var i, j int
	for n > 1 {
		i, j = 0, 0
		for j < n {
			nums[i] = nums[j]
			i++
			j += 2
		}
		n = i
		if n < 2 {
			break
		}
		i, j = 0, 1
		for j < n {
			nums[i] = nums[j]
			i++
			j += 2
		}
		n = i
		if n < 2 {
			break
		}
		i, j = 0, 0
		for j < n {
			if (j+1)%3 == 0 {
				j++
				continue
			}
			nums[i] = nums[j]
			i++
			j++
		}
		n = i
	}
	fmt.Printf("%d\n", nums[0])
}
