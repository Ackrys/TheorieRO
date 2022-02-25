package main

import "fmt"

func triInsertion(tab []int) []int {
    outtab := make([]int, 0, len(tab)+1)

    for i := 0; i < len(tab); i++ { // Parcours tableau d'origine
        for j := 0; j < len(outtab); j++ { // Parcours tableau de retour
            if outtab[j] > tab[i] { // Valeur courante plus grande
                tab[i], outtab[j] = outtab[j], tab[i] // Inversion des valeurs
            }
        }

        // Ajout de la valeur à la fin du tableau
        outtab = append(outtab, tab[i])
    }

    /*
        for i := 0; i < len(tab); i++ {
            var j int
            for ; j < len(outtab); j++ {
                if outtab[j] > tab[i] {
                    break
                }
            }
            endtab := append([]int{tab[i]}, outtab[j:]...)
            outtab = append(outtab[:j], endtab...)
        }
    */

    return outtab
}

func main() {

    tab := []int{2, 5, 7, 3, 6, 9, 4, 1, 0, 8}

    fmt.Println(triInsertion(tab))

}
