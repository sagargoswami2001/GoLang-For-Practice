package main
import "fmt"

func main()  {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is Divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are Even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is Negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 Digit")
	} else {
		fmt.Println(num, "has Multiple Digits")
	}
}
