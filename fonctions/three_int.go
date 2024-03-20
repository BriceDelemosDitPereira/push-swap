package fonctions

import (
	TabA "push_swap/fonctions/tabA"
)

// Il n'y a que 5 cas possible quand on a uniquement 3 chiffres dans la stack (site internet dans les com en bas)
func Three_int(stack_A []int) ([]int, string) {
	if !Check_sort(stack_A) {
		switch {
		case (stack_A[0] > stack_A[1]) && (stack_A[0] < stack_A[2]) && (stack_A[1] < stack_A[2]):
			stack_A = TabA.Sa(stack_A)
			return stack_A, "sa"
		case (stack_A[0] > stack_A[1]) && (stack_A[0] > stack_A[2]) && (stack_A[1] > stack_A[2]):
			stack_A = TabA.Rra(TabA.Sa(stack_A))
			return stack_A, "sa" + "\n" + "rra"
		case (stack_A[0] > stack_A[1]) && (stack_A[0] > stack_A[2]) && (stack_A[1] < stack_A[2]):
			stack_A = TabA.Ra(stack_A)
			return stack_A, "ra"
		case (stack_A[0] < stack_A[1]) && (stack_A[0] < stack_A[2]) && (stack_A[1] > stack_A[2]):
			stack_A = TabA.Ra(TabA.Sa(stack_A))
			return stack_A, "sa" + "\n" + "ra"
		default:
			stack_A = TabA.Rra(stack_A)
			return stack_A, "rra"
		}
	}
	return stack_A, ""
}
