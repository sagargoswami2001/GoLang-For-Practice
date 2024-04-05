package main
import (
	"fmt"
	"time"
)

func main()  {
	
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend")
	default:
		fmt.Println("It's a Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's Before Noon")
	default:
		fmt.Println("It's After Noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Bool")
		case int:
			fmt.Println("I'm an Int")
		default:
			fmt.Printf("Don't Know Type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
