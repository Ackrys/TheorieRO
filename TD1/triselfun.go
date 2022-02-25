package main

import "fmt"

func genTriSelection(compare func(int, int) bool) func([]int) {
	return func(tab []int) {
		for i := 0; i < len(tab)-1; i++ {
			imin := i
			minvalue := tab[i]
			for ii := i + 1; ii < len(tab); ii++ {
				if compare(tab[ii], minvalue) {
					imin, minvalue = ii, tab[ii]
				}
			}
			tab[i], tab[imin] = tab[imin], tab[i]
		}
	}
}

func main() {

	triCroissant := genTriSelection(func(arg1 int, arg2 int) bool {
		return arg1 < arg2
	})

	triDecroissant := genTriSelection(func(arg1 int, arg2 int) bool {
		return arg1 > arg2
	})

	tab := []int{2, 5, 7, 3, 6, 9, 4, 1, 0, 8}
	triCroissant(tab)
	fmt.Println(tab)

	tab = []int{2, 5, 7, 3, 6, 9, 4, 1, 0, 8}
	triDecroissant(tab)
	fmt.Println(tab)

}
