package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	Fonctions "push_swap/fonctions"
	TabA "push_swap/fonctions/tabA"
	TabB "push_swap/fonctions/tabB"
)

func main() {

	if len(os.Args) != 2 {
		return
	} else {
		var stack_A []int
		var stack_B []int
		arg1 := os.Args[1]
		arg_split := strings.Split(arg1, " ")

		for i := 0; i < len(arg_split); i++ {
			c, err := strconv.Atoi(arg_split[i])
			if err != nil {
				Fonctions.Erreur(err, "Vérifier que vous ne renseignez que des chiffres / nombres en argument sans doublons.")
				os.Exit(1)
				return
			}
			stack_A = append(stack_A, c)
		}

		scanner := bufio.NewScanner(os.Stdin)
		var echo []string

		res_expected := make([]int, len(stack_A))
		copy(res_expected, stack_A)
		sort.Ints(res_expected)

		for scanner.Scan() {
			line := scanner.Text()
			echo = append(echo, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Erreur de lecture de l'entrée standard :", err)
		}
		// Utilisez le tableau args comme argument pour le prochain exécutable

		for _, k := range echo {
			switch k {
			case "pa":
				stack_A, stack_B = TabA.Pa(stack_A, stack_B)
			case "ra":
				stack_A = TabA.Ra(stack_A)
			case "rra":
				stack_A = TabA.Rra(stack_A)
			case "sa":
				stack_A = TabA.Sa(stack_A)
			case "pb":
				stack_A, stack_B = TabB.Pb(stack_A, stack_B)
			case "rb":
				stack_B = TabB.Rb(stack_B)
			case "rrb":
				stack_B = TabB.Rrb(stack_B)
			case "sb":
				stack_B = TabB.Sb(stack_B)
			case "ss":
				stack_A = TabA.Sa(stack_A)
				stack_B = TabB.Sb(stack_B)
			case "rr":
				stack_A = TabA.Ra(stack_A)
				stack_B = TabB.Rb(stack_B)
			case "rrr":
				stack_A = TabA.Rra(stack_A)
				stack_B = TabB.Rrb(stack_B)
			}
		}
		if Fonctions.Compare(stack_A, res_expected) {
			fmt.Println("OK")
		} else {
			fmt.Println("KO")
		}
	}
}

/*
Lorsque vous exécutez votre programme avec la commande echo -e,
vous pouvez rediriger la sortie de echo -e vers l'entrée de votre programme en utilisant le pipe (|).
Par exemple : echo -e "sa\npb\nrrr\n" | ./checker
*/
