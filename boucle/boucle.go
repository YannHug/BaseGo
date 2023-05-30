package boucle

import "fmt"

func Main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum result=%v\n", sum)

	evenCnt := 0
	for evenCnt < 3 {
		fmt.Println("Retrieving events ...")
		evenCnt++
		if evenCnt == 3 {
			fmt.Printf("Got %d notifications\n", evenCnt)
		}
	}

	i := 0
	for {
		i++
		if i%2 != 0 {
			fmt.Println("Odd looping ...")
			continue
		}
		fmt.Println("Looping....")
		if i >= 10 {
			fmt.Println("Loop end")
			break
		}
	}
}
