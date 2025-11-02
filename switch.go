package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("its a bool")
		case int:
			fmt.Println("its a int")
		default:
			fmt.Printf("unknown type: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
