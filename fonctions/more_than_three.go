package fonctions

import (
	"fmt"
	TabA "push_swap/fonctions/tabA"
	TabB "push_swap/fonctions/tabB"
)

func More_than_three(stack_A []int) ([]int, string) {

	var smallest_int_index int
	var stack_B []int
	var action string

	if !Check_sort(stack_A) && len(stack_B) == 0 {

		for len(stack_A) > 3 {
			smallest_int_index = Smallest_int_index(stack_A)
			switch {
			case (smallest_int_index <= len(stack_A)/2) && smallest_int_index != 0:
				if smallest_int_index == 1 {
					stack_A = TabA.Sa(stack_A)
					fmt.Println("sa")
				} else {
					stack_A = TabA.Ra(stack_A)
					fmt.Println("ra")
				}
			case (smallest_int_index > len(stack_A)/2):
				stack_A = TabA.Rra(stack_A)
				fmt.Println("rra")
			default:
				stack_A, stack_B = TabB.Pb(stack_A, stack_B)
				fmt.Println("pb")
			}
		}
	}
	if !Check_sort(stack_A) && len(stack_B) != 0 {

		stack_A, action = Three_int(stack_A)
		fmt.Println(action)

	}

	if Check_sort(stack_A) && len(stack_B) > 0 {
		for len(stack_B) > 0 {
			stack_A, stack_B = TabA.Pa(stack_A, stack_B)
			fmt.Println("pa")
		}
	}
	return stack_A, ""
}
