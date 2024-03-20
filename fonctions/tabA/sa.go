package tabA

func Sa(stack []int) []int {
	buffer := stack[0]
	stack[0] = stack[1]
	stack[1] = buffer
	return stack
}
