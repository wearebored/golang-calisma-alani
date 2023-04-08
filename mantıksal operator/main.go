package main
import "fmt"



func main()  {
	// x:=10
	// y:=20

	// if x<y {
	// 	fmt.Printf("%v büyüktür %v",y,x)
	// }
	

	// fmt.Println("a")
	// fmt.Println(sum(10,20))
	// z:=1
	// for z<10{
	// 	z++
	// 	fmt.Println(z)
	// }

	// type datalar struct{
	// 	name string
	// 	age int
	// 	isMarried bool
	// }

	// var es datalar
	// es.name="enes"
	// es.age=26
	// es.isMarried=false

	// fmt.Println(es)
		x:=5
		fmt.Println(&x)
		duble(&x)
		fmt.Println(x)


}

func duble(sum *int){
	*sum = *sum*2

}

func sum (x,y int ) int{
return x+y
} 