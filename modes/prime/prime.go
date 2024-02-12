package prime

import (
	"fmt"
)

func IsPrime(p int) {
  mid := "IS"
  if (!isPrime(p)) { mid += " NOT" } 
  fmt.Printf("%d %s prime\n", p, mid)
}

func isPrime(p int) bool {
  if p > 1 && p <= 3 {
    return true 
  } 

  if p <= 1 || p%2 == 0 || p%3 == 0 {
    return false 
  }  
    for i := 5; i*i <= p; i += 6 {
      if p%i == 0 || p%(i+2) == 0 {
        return false
      }
    }
    return true 
}
