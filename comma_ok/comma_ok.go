package main

import "fmt"

func main() {

	if ok, err := isEven(5); ok {
		fmt.Println(ok)
	} else {
		fmt.Println(err)
	}

}

func isEven(n int) (bool, error) {
	if n%2 == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("Number is not even.")

	}
}
