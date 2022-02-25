package main

import "fmt"

func triSelection(tab []int) {
	for i := 0; i < len(tab)-1; i++ {
		imin := i
		minvalue := tab[i]
		for ii := i + 1; ii < len(tab); ii++ {
			if tab[ii] < minvalue {
				imin, minvalue = ii, tab[ii]
			}
		}
		tab[i], tab[imin] = tab[imin], tab[i]
	}
}

func main() {

	tab := []int{2, 5, 7, 3, 6, 9, 4, 1, 0, 8}

	triSelection(tab)

	fmt.Println(tab)

}
