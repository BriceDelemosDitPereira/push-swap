package fonctions

import (
	TabA "push_swap/fonctions/tabA"
)

func Two_int(stack_A []int) ([]int, string) {
	if !Check_sort(stack_A) {
		stack_A = TabA.Sa(stack_A)
		return stack_A, "sa"
	}
	return stack_A, ""
}
