package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	endpoints, err := GetEndPoints()

	fmt.Println(endpoints)
	fmt.Println(err)
}
