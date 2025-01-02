package big_o

func addItems(n int) int {
	return n + n + n
}

func o1(val int) int {
	return addItems(val)
}
