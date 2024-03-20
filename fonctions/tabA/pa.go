package tabA

func Pa(stack_A []int, stack_B []int) ([]int, []int) {
	var tab_A []int
	var tab_B []int
	// Ajouter le premier entier de stack_B dans stack_A
	tab_A = append(tab_A, stack_B[0])
	// Copier les éléments restants de stack_A dans tab
	for i := 0; i < len(stack_A); i++ {
		tab_A = append(tab_A, stack_A[i])
	}

	for j := 1; j < len(stack_B); j++ {
		tab_B = append(tab_B, stack_B[j])
	}
	return tab_A, tab_B
}
