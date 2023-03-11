package main
import "fmt"

func main(){
	var name string ="enes"
	var name2 string
	var age int
	age =22
	name2="göktaş"
	fmt.Println("hello ", name,name2, "yaşı:",age)


	// var kullanımı
	var(
		name3="en"
		age2 int =25
		height float32=1.75
	)
	fmt.Println(name3,age2,height)

	// tek satır var

	var names1,ages1,marrieds1 = "enes", 25,true
	fmt.Println(names1,ages1,marrieds1)

	// tek satırda değer atama := ile
	names2,ages2,marrieds2 := "enes", 25,true
	fmt.Println(names2,ages2,marrieds2,":= kullanarak")

}