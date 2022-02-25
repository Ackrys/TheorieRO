package main

import "fmt"

type monType struct {
	valeurTri int
	contenu   string
}

type selection interface {
	echangeAvecMinimum(int)
	estMinimum(int) bool
	defMinimum(int)
	nombreElements() int
}

type monTypeTab struct {
	contenu    []monType
	minimum    monType
	minimumPos int
}

func (mT *monTypeTab) echangeAvecMinimum(pos int) {
	mT.contenu[pos], mT.contenu[mT.minimumPos] = mT.contenu[mT.minimumPos], mT.contenu[pos]
}

func (mT *monTypeTab) estMinimum(pos int) bool {
	return mT.contenu[pos].valeurTri < mT.minimum.valeurTri
}

func (mT *monTypeTab) defMinimum(pos int) {
	mT.minimum = mT.contenu[pos]
	mT.minimumPos = pos
}

func (mT *monTypeTab) nombreElements() int {
	return len(mT.contenu)
}

func triSelection(sel selection) {
	for i := 0; i < sel.nombreElements()-1; i++ {
		sel.defMinimum(i)
		for ii := i + 1; ii < sel.nombreElements(); ii++ {
			if sel.estMinimum(ii) {
				sel.defMinimum(ii)
			}
		}
		sel.echangeAvecMinimum(i)
	}
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

	triSelection(&sel)

	fmt.Println(sel.contenu)

}
