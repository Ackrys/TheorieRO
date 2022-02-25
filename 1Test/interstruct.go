package main

import "fmt"

type insertion interface {
  nombreElements() int //nomnbre d'elements dans la collection
  enregistreValeur(int) int // mémoriser valeur située à une position
  plusGrand(int, int) bool  // dire si valeur à une position est plus grande que valeur à une autre position
  inverseValeur(int) // inverser une valeur à une position et valeur mémorisée
  ajouteValeur() // ajouter valeur mémorisée en fin de collection
}

func triInsertion(t []int) (tt []int){
  tt = make([]int, 0, lent(t))


  for i :=0; i < t.nombreElements(); i++ {
    var j int
    t.enregistreValeur(i)
    for ; j < t.nombreElements(); j++ {
      if tt.plusGrand(j) {
        break
      }
    }
    for ; j < tt.nombreElements(); j++ {
      tt.inverseValeur(j)
    }
    tt.ajouteValeur(j)
  }
  return tt
}

func main(){

}
