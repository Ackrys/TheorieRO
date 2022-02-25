package main

import "fmt"

type monType struct {
	valeurTri int
	contenu   string
}

type insertion interface {
	nombreElements() int
	estPlusGrand(int, int) bool
	reordonne([]int)
}

type monTypeTab struct {
	contenu    []monType
	minimum    monType
	minimumPos int
}

func (mT *monTypeTab) nombreElements() int {
	return len(mT.contenu)
}

func (mT *monTypeTab) estPlusGrand(i, j int) bool {
	return mT.contenu[i].valeurTri > mT.contenu[j].valeurTri
}

func (mT *monTypeTab) reordonne(t []int) {
	res := make([]monType, len(mT.contenu))
	for i, v := range t {
		res[i] = mT.contenu[v]
	}
	mT.contenu = res
}

func triInsertion(ins insertion) {
	outtab := make([]int, 0, ins.nombreElements())
	for i := 0; i < ins.nombreElements(); i++ {
		var j int
		for ; j < len(outtab); j++ {
			if ins.estPlusGrand(outtab[j], i) {
				break
			}
		}
		endtab := append([]int{i}, outtab[j:]...)
		outtab = append(outtab[:j], endtab...)
	}
	ins.reordonne(outtab)
}

func main() {

	sel := monTypeTab{
		contenu: []monType{
			monType{valeurTri: 2, contenu: "Deux"},
			monType{valeurTri: 5, contenu: "Cinq"},
			monType{valeurTri: 7, contenu: "Sept"},
			monType{valeurTri: 3, contenu: "Trois"},
			monType{valeurTri: 6, contenu: "Six"},
			monType{valeurTri: 9, contenu: "Neuf"},
			monType{valeurTri: 4, contenu: "Quatre"},
			monType{valeurTri: 1, contenu: "Un"},
			monType{valeurTri: 0, contenu: "Zero"},
			monType{valeurTri: 8, contenu: "Huit"},
		},
	}

	triInsertion(&sel)

	fmt.Println(sel.contenu)

}
