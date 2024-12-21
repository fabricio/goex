package main

var (
	b bool    = true
	c int     = 10
	d string  = "World"
	e float64 = 1.2
)

func main(){
  var meuArray [3]int
  meuArray[0] = 10
  meuArray[1] = 20
  meuArray[2] = 30

  for i, v := range meuArray {
    println(i, v)
  }
 }
