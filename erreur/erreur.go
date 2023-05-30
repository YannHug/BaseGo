package erreur

import (
	"fmt"
	"os"
)

func readFile(fileName string) (string, error) {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	if len(dat) == 0 {
		//return "", errors.New("Empty Content")
		return "", fmt.Errorf("Empty Content (filename=%v)", fileName)
	}

	return string(dat), nil
}

func Main() {
	dat, err := readFile("Test.txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	fmt.Println(dat)
}
