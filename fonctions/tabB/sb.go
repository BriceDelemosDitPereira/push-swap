package tabB

func Sb(stack []int) []int {
	buffer := stack[0]
	stack[0] = stack[1]
	stack[1] = buffer
	return stack
}
