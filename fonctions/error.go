package fonctions

import "fmt"

func Erreur(err error, message string) {
	fmt.Println("!", err, "!"+"\n"+message)
}
