package main

import (
	"math"
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

 /*
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
 */
/*
func main() {
	var i, j int = 1, 12
	p := &i
	f := &i
	println(f)
	*f = 2
	println(*f)
	println(i, j, p, *p)


}
*/

type adict struct {
	key string
	value int
}

func main()  {
	/*
	v := adict{"a",4}
	p := &v
	p.key = "ERE"
	fmt.Println(v)
	 */
	// var (
	// 	v1 = adict{}
	// 	v2 = adict{"2", 5}
	// 	v3 = adict{key: "rer"}
	// 	p = &adict{"#@#", 3}
	// )
	//fmt.Println(v1, v2, v3, p)
	/*

	 var a [10] int
	  s := a[7: ]
	  b := a[6:]
	  p := &a[8]
	  *p = 3
	  s[0] = 1
	 fmt.Println(b)
	 */

	 /*

	  a := []int{1, 2, 3}
	  fmt.Println(a)
	  b := []struct {
		  k int
		  v bool
	  }{
	  	{1, true},
	  	{1, true},
	  }
	  fmt.Println(b)

	  c := []int{}
	  fmt.Println(c)
	 	  */
	 	  /*

	var ere []int
	er := []int{}
	fmt.Println(er, ere)
	fmt.Println(len(ere), cap(ere), ere)
	if er != nil{
		fmt.Println("%#H#I#")
	}
	 	   */

	 	   /*

	era := make([]int, 0, 10)
	fmt.Println(era)
	 	    */
	 	    /*

	 	    e := [][]int{[]int{1, 2}}
	 	    fmt.Println(e)
	 	    	 	     */


	borad := [][] string{
		[]string{"_", "_", "#"},
		[]string{"_", "ew", "#"},
	}

	//fmt.Println(borad)

	borad[0][0] = "T"
		//fmt.Println(borad)

	borad[1][2] = "Y"
			//fmt.Println(borad)

	//for i := 0; i < len(borad); i++{
	//	fmt.Println(borad[i])
	//}

	for i := 0; i < 10; i++ {
		//println(i)
	}

	var s []int
	//s = append(s, 9)
	//fmt.Println(s)
	s = append(s, 6, 7, 8, 3)
	//	fmt.Println(s)


	for _, i := range s{
		fmt.Printf("%d\n", i)
	}

}