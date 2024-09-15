package main

import (
  "fmt"
  "input"
  "log"
)

func main() {
  fmt.Println("Choose operator (+, -, *, /)")
  fmt.Printf("Choose: ")

  operator, err := input.GetString()
  if err != nil {
    log.Fatal(err)
  } else if operator != "+" &&  operator != "-" && operator != "*" && operator != "/" {
    log.Fatal("Invalid operator: ", operator)
  }
  
  fmt.Printf("Type first number: ")
  fnum, err := input.GetFloat()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Type second number: ")
  snum, err := input.GetFloat()
  if err != nil {
    log.Fatal(err)
  }

  var res float64
  switch operator {
    default:
      fmt.Println("error default")
     
    case "+":
      res = fnum + snum
    case "-":
      res = fnum - snum
    case "*":
      res = fnum * snum
    case "/":
      res = fnum / snum
  }

  fmt.Printf("%0.2f %s %0.2f = %0.2f\n", fnum, operator, snum, res)
  }
