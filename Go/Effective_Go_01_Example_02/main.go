package main

import "fmt"

func main() {
	value := 1
	// if
	if value == 1 {
		fmt.Println("value equals 1!")
	}

	// switch
	for i := 0; i < value; i++ {
		fmt.Printf("value is %d\n", i)
	}

	// switch case
	switch value {
	case 1:
		fmt.Println("value equals 1!")
	case 2:
		fmt.Println("value equals 2!")
	}

	x := 1
	y := 1
	// >> <<
	fmt.Printf("%d", x<<8+y<<16)
}
