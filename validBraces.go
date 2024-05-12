package kata

func ValidBraces(str string) bool {
	stack := make([]rune, 0)

	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (char == ')' && top != '(') || (char == ']' && top != '[') || (char == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
