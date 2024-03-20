package tabB

func Rrb(stack []int) []int {
	var tab []int
	tab = append(tab, stack[len(stack)-1])
	for i := 0; i < len(stack)-1; i++ {
		tab = append(tab, stack[i])
	}
	return tab
}
