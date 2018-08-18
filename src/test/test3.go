package main

import (
	"math"
	"time"
	"fmt"
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func ret_double(x int) (z, y float32){
	z = float32(x) / 10
	y = float32(x) * 10
	return z, y
}


var (
	x, y int = 3, 4
	f float64 = math.Sqrt(float64(x * x + y * y))
	z uint = uint(f)


)

func NeedInt(x int) int {
	return x * 100 + 1
}

func NeedFloat(x float64) float64{
	return x * 0.1
}

//func main() {
//	sum := 10
//	for sum > 1{
//		if s := sum; s == 5{
//			println(s)
//
//		} else {
//			println(s)
//		}
//		sum -= 1
//	}
//}

/*

水电费水电费
/*




func main() {
	sum := 10
	for sum > 1{
		switch s := sum; s {
		case 1 : println(s)
		case 5: println(s)
		case 9: println(s)
		default:
			var j string = "sfd"
			println(j)

		}
		sum -= 1
	}
}

 */

func main() {
	today := time.Now().Weekday()
	println(time.Sunday)
	switch time.Sunday {
	case today - 6:
		fmt.Println("Todaysss", today, time.Sunday, today-4)
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two Day", today)
	default:
		fmt.Println("haha")
	}
}
