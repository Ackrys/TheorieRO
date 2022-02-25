package main2;

import ("fmt"
  "time"
)

func ping(c chan int){
  for i := 0; i < 100; i++{
    fmt.Println("ping", i)
  }
  c <- 0
}

func pong(c chan int){
  for i := 0; i < 100; i++{
    fmt.Println("pong", i)
  }
  c <- 1
}


func qui(id int) string{
    if id == 0 {
      return "ping"
    }else{
      return "pong"
    }
}

func main(){

  var c chan int = make(chan int, 2)

  var compte int = 0

  go ping(c)
  go pong(c)
  for compte < 3{
    select {
    case fini := <-c :
      compte++
      fmt.Println(qui(fini), "a terminé")
    case <- time.After(time.Second) :
      compte++
      fmt.Println("Personne n'a terminé")
    }
  }
}
