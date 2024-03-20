package fonctions

func Check_sort(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i-1] > stack[i] {
			return false
		}
	}
	return true
}

// Attention ne pas mettre de return true dans la boucle for où sinon, elle ne va check que les 2 premiers chiffres
