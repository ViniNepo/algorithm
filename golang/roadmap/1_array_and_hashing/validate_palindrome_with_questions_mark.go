package array_and_hashing

func ValidatePalindromeWithQuestions(str string) string {
	strCount := [26]int{}
	markCount := 0

	for _, letter := range str {
		if letter == '?' {
			markCount++
		} else {
			strCount[letter-'a']++
		}
	}

	isOdd := len(str)%2 == 1
	middle := 0

	for i := 25; i >= 0; i-- {
		if strCount[i]%2 == 1 {
			if isOdd {
				isOdd = false
				middle = i
				strCount[i] /= 2
			} else {
				strCount[i] = (strCount[i] + 1) / 2
				markCount--
			}
		} else {
			strCount[i] /= 2
		}
	}

	if markCount < 0 {
		return "-1"
	}

	inOrder := ""
	reverse := ""
	for i, letter := range strCount {
		if letter > 0 {
			inOrder += string(rune(i + 'a'))
			reverse = string(rune(i+'a')) + reverse
		}
	}

	if len(str)%2 == 1 {
		return inOrder + string(rune(middle+'a')) + reverse
	}

	return inOrder + reverse
}
