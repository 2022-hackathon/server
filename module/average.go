package module

func Average(num ...int) int {

	res := 0

	for _, v := range num {
		res += v
	}

	return res / len(num)
}
