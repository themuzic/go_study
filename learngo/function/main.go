package main

import "fmt"

// func multiply(a, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (length int, upperCase string) {
// 	defer fmt.Println("I'm done")
// 	length = len(name)
// 	upperCase = strings.ToUpper(name)
// 	return
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func superAdd(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper("nico")
	// fmt.Println(totalLength, upperName)

	// repeatMe("nico", "lynn", "dal", "marl", "flynn")

	// fmt.Println(superAdd(1, 2, 3, 4, 5, 6))

	fmt.Println(canIDrink(14))
}
