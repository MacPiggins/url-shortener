package Base63

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func convertNumber(number uint8) string {
	var base63 string
	if number == 0 {
		return string(alphabet[0])
	}
	for number > 0 {
		base63 = string(alphabet[number%63]) + base63
		number /= 63
	}

	return base63
}

func ConvertToBase63(input []byte) string {
	var base63 string

	for i := 0; i < len(input); i++ {
		number := input[i]
		base63 += convertNumber(number)
	}

	return base63
}
