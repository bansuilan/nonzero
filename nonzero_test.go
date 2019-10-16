package nonzero

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	w := New(3, 3)
	w.Set(1, 1, 1)
	w.Set(1, 2, 2)
	w.Set(2, 1, 0)
	w.Set(2, 2, 2)
	w.Set(2, 3, 2)
	w.Set(3, 1, 1)
	w.Set(3, 2, 1)
	w.Set(3, 3, 1)
	fmt.Printf("Chessboard:\n\n")
	fmt.Println(w.String())
	fmt.Printf("Solution:\n\n")
	solution, err := w.Solve()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range solution {
		fmt.Printf("%4d\t%s\n", i, v)
	}
	fmt.Printf("\n\n\n\n\n")
}
