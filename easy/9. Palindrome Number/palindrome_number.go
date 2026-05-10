package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	invertido := 0

	for x > 0 {
		digito := x % 10
		invertido = invertido*10 + digito
		x = x / 10
	}

	return original == invertido
}
