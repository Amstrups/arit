package random

import (
	"fmt"
	"testing"
)

func TestCapitilization(t *testing.T) {
	s := "?notcapitilized#CAPITILIZED!"

	s_ := Capitilization2(s)
  fmt.Println(s_)
}
