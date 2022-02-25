package main

import "fmt"

type monInterface interface {
  plusx(int) int
  fois2()int
}

type monEntier struct {
  val int
}

func test(mi monInterface) int {
  return mi.plusx(mi.fois2())
}

func (m autreEntier) plusx(x int) int {
  return m.val + x
}

func (m autreEntier) fois2() int {
  if m.info {
    return m.val * 2
  }
  return 0
}

type autreEntier struct {
  val int
  info bool
}

func main(){
  /*
  var i int
  var b bool
  var s string
  var f float64

  fmt.Println(i)
  fmt.Println(b)
  fmt.Println(s)
  fmt.Println(f)
  */
  m := autreEntier{val: 2}
  fmt.Println(test(m))
}
