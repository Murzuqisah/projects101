package prime

func IsPrime(n int) bool {
	// 1 and negative numbers are not prime numbers
	if n <= 1 {
		return false
	}

	// loop through the values given
	// a prime number's divisors are 1 and itself
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
