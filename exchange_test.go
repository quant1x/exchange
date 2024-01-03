package exchange

import (
	"fmt"
	"testing"
)

func TestGetExchange(t *testing.T) {
	sessions := GetExchange(CN)
	fmt.Println(sessions)
	fmt.Println(sessions.Minutes())
}
