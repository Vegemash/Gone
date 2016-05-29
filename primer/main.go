//Prime finds the first 1000000 primes and prints them
package main

import "fmt"


func main() {
	fmt.Println(primesTo(1000000));
}

//Returns the first 'goal' primes
//
func primesTo(goal int) []int{
	primes := make([]int, 0)

	//Defencive guards
	if goal == 0{
		return []int(primes)
	}

	primes = append (primes, 2)

	counter := 3
	//Loop the counter until we have the goal amount of primes
	for len(primes) < goal	{
		primesCounter := 0
		prime := true
		//while the prime we are checking the counter against is 
		//less than the square root of the counter
		for counter > (primes[primesCounter] * primes[primesCounter]) {
			//if the prime cleanly divides the counter the counter is not prime
			if(counter % primes[primesCounter] == 0){
				prime = false
				break
			}else{
				//check against the n ext prime
				primesCounter++
			}
			if primesCounter > len(primes) {
				break
			}
		}
		//WOO we found a prime, add it
		if prime {
			primes = append(primes, counter)
		}
		//increment the counter, ignoring even numbers
		counter += 2
	}
	return []int(primes)
}