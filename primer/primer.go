package main

import "fmt"
import "math"

func main() {
	primes := FirstNPrimes(10000)
	var countingPrimes []int
	countingPrimesCounter := 1
	counter := 0
	for countingPrimesCounter <= 100 && counter < len(primes) {
		if intNSigFigs(primes[counter], intLog10(countingPrimesCounter)+1) == countingPrimesCounter {
			countingPrimes = append(countingPrimes, primes[counter])
			countingPrimesCounter++
		}
		counter++
	}
	fmt.Println(countingPrimes)
	fmt.Println(primes[len(primes)-1])
}

//FirstNPrimes : returns the first 'goal' primes
func FirstNPrimes(goal int) []int {
	var primes []int

	//Defencive guards
	if goal == 0 {
		return []int(primes)
	}

	primes = append(primes, 2)

	counter := 3
	//Loop the counter until we have the goal amount of primes
	for len(primes) < goal {
		primesCounter := 0
		prime := true
		//while the prime we are checking the counter against is
		//less than the square root of the counter
		for counter >= (primes[primesCounter] * primes[primesCounter]) {
			//if the prime cleanly divides the counter the counter is not prime
			if counter%primes[primesCounter] == 0 {
				prime = false
				break
			} else {
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

func intNSigFigs(in int, n int) int {
	digits := intLog10(in) + 1
	if digits <= n {
		return in
	}

	return in / intPow(10, digits-n)
}

func intLog10(in int) int {
	return int(math.Log10(float64(in)))
}

func intPow(in int, pow int) int {
	return int(math.Pow(float64(in), float64(pow)))
}
