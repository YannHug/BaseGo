package forEach

import "fmt"

func Main() {
	nums := []int{10, -2}
	for i, num := range nums {
		fmt.Printf("nums[=%d] = %d\n", i, num)
	}

	for _, c := range "golang" {
		fmt.Printf("%v\n", string(c))
	}
}
