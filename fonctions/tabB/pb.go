package tabB

func Pb(stack_A []int, stack_B []int) ([]int, []int) {
	var tab_A []int
	var tab_B []int
	// Vérifier si stack_B n'est pas vide
	if len(stack_A) > 0 {
		// Ajouter le premier entier de stack_B dans stack_A
		tab_B = append(tab_B, stack_A[0])
		// Copier les éléments restants de stack_A dans tab
		for i := 0; i < len(stack_B); i++ {
			tab_B = append(tab_B, stack_B[i])
		}
		tab_A = append(stack_A[1:])
		// for j := 1; j < len(stack_A); j++ {
		// 	tab_A = append(tab_A, stack_A[j])
		// }
		return tab_A, tab_B
	}

	return tab_A, tab_B
}
