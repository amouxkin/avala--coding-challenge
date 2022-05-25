package lcg

import (
	"math/big"
)

// GeneralGenerator Linear Congruential Generator
// Pseudo Random Number
func GeneralGenerator(previous, constant, seed, sampleSpace int) int {
	previous = ((previous * seed) + constant) % sampleSpace
	return previous
}

// LehmerGenerator An implementation of Lehmer random number generator
// Requirements:
//		sampleSpace (M) must be a prime or power of prime.
//		(This is so that modulus works similar to as in an incremental operation)
//
//		seed should be co-prime to sampleSpace
//
//		previous is an element of high multiplicative order
func LehmerGenerator(previous, seed, sampleSpace int) func() int {
	// Check if seed is prime
	if !big.NewInt(int64(seed)).ProbablyPrime(0) {
		panic("seed is not prime")
	}

	// Check is seed is co-prime
	if seed == sampleSpace {
		panic("seed is not a co-prime of sampleSpace")
	}

	// Set
	factors := make(map[int]struct{})
	// Check if sampleSpace is Prime or power of prime
	for i := 1; i < sampleSpace; i++ {
		if sampleSpace%i == 0 {
			factors[i] = struct{}{}
		}
	}

	if len(factors) > 2 {
		panic("sampleSpace is not a prime / power of prime")
	}

	if previous > sampleSpace {
		panic("invalid previous value")
	}

	internalPrevious := previous

	return func() int {
		internalPrevious = GeneralGenerator(internalPrevious, 0, seed, sampleSpace)
		return internalPrevious
	}

}
