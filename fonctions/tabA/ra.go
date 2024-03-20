package tabA

func Ra(stack []int) []int {
	var tab []int
	for i := 1; i < len(stack); i++ {
		tab = append(tab, stack[i])
	}
	tab = append(tab, stack[0])
	return tab
}
