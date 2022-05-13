package main

import "fmt"

func find_primes(number int) (result []int) {
	if number == 0 {
		fmt.Println("0 is not prime")
		return nil
	}
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	if len(result) > 2 { 
		fmt.Printf("The number %d IS NOT prime\n", number)
		fmt.Println("This is the list of its divisors: ")
	} else {
		fmt.Printf("The number %d IS prime\n", number)
		fmt.Println("This is the list of its divisors: ")

	}

	return result
}
func main() {
	var user_number int
	fmt.Scanf("%d", &user_number)
	var primes []int                  
	primes = find_primes(user_number) 
	fmt.Println(primes)
}
