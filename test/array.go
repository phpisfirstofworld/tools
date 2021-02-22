package main

import "github.com/PeterYangs/tools"

func main() {

	array := []string{"1", "2", "3", "4"}

	b := tools.In_array(array, "4")

	println(b)

	array2 := []int{1, 2, 3, 4}

	b2 := tools.In_array(array2, 4)

	println(b2)

}
