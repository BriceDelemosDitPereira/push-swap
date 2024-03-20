package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	Fonctions "push_swap/fonctions"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		var stack_A []int                     // t'initie ton tableau A
		var action string                     // permet de print l'action effectuée
		arg1 := os.Args[1]                    // récup arg 1
		arg_split := strings.Split(arg1, " ") // split sur les espace

		for i := 0; i < len(arg_split); i++ { // range le split
			c, err := strconv.Atoi(arg_split[i]) // tu fais un Atoi sur les string "1" "2"...
			if err != nil {                      // gestion d'erreur
				Fonctions.Erreur(err, "Vérifier que vous ne renseignez que des chiffres / nombres en argument sans doublons.")
				os.Exit(1)
				/*
					- L'exit status 0 est généralement utilisé pour indiquer une exécution réussie sans erreur.
					- L'exit status 1 est souvent utilisé pour indiquer une erreur générique ou une exécution anormale.
					- L'exit status 2 (et les valeurs supérieures) peuvent être utilisés pour indiquer des erreurs spécifiques ou des conditions d'erreur particulières.
				*/
				return // en cas de gestion d'erreur dans un nil, il est important de placer un retrun pour sortir du main et mettre fin au programme
			}
			stack_A = append(stack_A, c) // tu rentre les premier chiffre dans le tableau A
		}

		if !Fonctions.Check_doublon(stack_A) {
			fmt.Println("Vérifier qu'il n'y a pas de doublons ! Chacal...")
		}
		if Fonctions.Check_sort(stack_A) {
			return
		} else {
			//fmt.Println(stack_A) // tab pré-modif
			switch len(stack_A) {
			case 1:
				fmt.Println("nombre d'argument insuffisant, renseigner des numéros")
			case 2:
				stack_A, action = Fonctions.Two_int(stack_A)
			case 3:
				stack_A, action = Fonctions.Three_int(stack_A)
			default:
				stack_A, action = Fonctions.More_than_three(stack_A)
			}
			fmt.Println(action)
			//fmt.Println(stack_A) // tableau après modif
		}
	}
}

/* Chaque information pour comprendre l'énoncé a été trouvé sur ce site : https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a

On a 2 stacks (A et B), A est une liste de nombres non organisés et B est vide.
B permet de faire certaines modifications pour ranger les nombres de manière croissante

Après avoir généré les chiffres dans A aléatoirement, l y a plusieurs commandes :
	- sa/sb : permet d'échanger les 2 premiers chiffres au sein d'une stack (sa : dans la stack A, sb dans l'autre)
	- ss : fait sa et sb en même temps

	- ra/rb : permet de mettre le premier chiffre tout en bas de la stack (soit A ou B)
	- rr : fait ra et rb en même temps

	- rra/rrb : permet de mettre le dernier chiffre tout en haut de la stack
	- rrr : fait rra et rrb en même temps

	- pa : permet de mettre le premier chiffre de B tout en haut de la stack A (donc en premier chiffre)
	- pb : permet de mettre le premier chiffre de A tout en haut de la stack B

On va faire une fonction pour chaque direction, nécessaire pour le test Checker

Maitenant que nous avons toute les actions possible, on va voir comment les utiliser
L'algorithme que l'on va utiliser dépend du nombre de chiffres aléatoires appliqués à la stack A

Il y a 4 cas principaux, avec 3 chiffres, avec 5, avec 100 et avec 500

deepai => permet de casser un code quand on veut tester notre code après l'avoir fait
(moins poussé que chatgpt) */

/*
Pour ajouter un exit status à votre code, vous pouvez utiliser la fonction os.Exit() en passant le code de sortie souhaité en argument.
Par exemple, pour définir un exit status de 1, vous pouvez utiliser os.Exit(1).
*/

/*
Pour Créer un fichier executable d'un main.go, se mettre a sa racine et utiliser la commande "go build -o <nom voulu pour le fichier exe> main.go"
Le symbole $ est utilisé pour accéder à la valeur d'une variable d'environnement en shell.
*/
