package main

import "fmt"

func fizzbuzz() {
	for i := range 100 {
		i += 1
		// if i%3 == 0 && i%5 == 0 {
		// 	fmt.Println("fizzbuzz")
		// } else if i%3 == 0 {
		// 	fmt.Println("fizz")
		// } else if i%5 == 0 {
		// 	fmt.Println("buzz")
		// } else {
		// 	fmt.Println(i)
		// }
		switch true {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
