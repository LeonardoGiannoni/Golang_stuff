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
	if len(result) > 2 { // il numero it is not prime, perchè l'array di divisori del numero avrà lunghezza maggiore di 2
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
	var primes []int                  // dichiarata una slice, che in Go ha l funzione e le caratteristiche di una lista (in Python lista definita cosi: [] -> es. l = [1,3,'a'])
	primes = find_primes(user_number) // passaggio per valore
	fmt.Println(primes)
}
