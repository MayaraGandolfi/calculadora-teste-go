package main

import "fmt"

func main(){
	x:= soma(1,2,3)
	fmt.Println(x)

	x = subtrai(30,20)
	fmt.Println(x)

	x = multiplica(4,4,4)
	fmt.Println(x)

	x = divide(10,5)
	fmt.Println(x)


}

func soma (i ... int) int{
	total :=0
 	
	for _, v := range i{
		total += v
	}
	return total
}

func subtrai (a int,b int) int{
	total := a - b
	return total
}

func multiplica (i ... int) int{
	total :=1
 	
	for _, v := range i{
		total *= v
	}
	return total
}

func divide (a int,b int) int{
	total := a / b
	return total
}
