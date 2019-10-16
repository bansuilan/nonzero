package nonzero

import (
	"bytes"
	"fmt"
)

// Step ...
type Step struct {
	x, y int64
}

func (s *Step) String() string {
	return fmt.Sprintf("(%d, %d)", s.x+1, s.y+1)
}

// NonZero ...
type NonZero struct {
	chessboard []int64
	xMax       int64
	yMax       int64
}

// New ...
func New(x, y int64) NonZero {
	cb := make([]int64, x*y)
	for i := range cb {
		cb[i] = -1
	}
	return NonZero{cb, x, y}
}

// RunStep ...
func (n *NonZero) RunStep(s Step, blank int64) {
	xB := blank / n.xMax
	yB := blank % n.xMax
	xO := 2*xB - s.x
	yO := 2*yB - s.y
	newB := n.get(s.x, s.y)
	n.set(s.x, s.y, -1)
	if xO >= 0 && xO < n.xMax && yO >= 0 && yO < n.yMax {
		n.set(xO, yO, n.get(xO, yO)-1)
		newB--
	}
	n.set(xB, yB, newB)
}

// Get ...
func (n *NonZero) get(i, j int64) int64 {
	return n.chessboard[n.xMax*i+j]
}

func (n *NonZero) set(i, j int64, num int64) {
	n.chessboard[n.xMax*i+j] = num
}

// Set ...
func (n *NonZero) Set(i, j int64, num int64) {
	n.set(i-1, j-1, num)
}

// GetAllStep ...
func (n *NonZero) GetAllStep(blank int64) []Step {
	x := blank / n.xMax
	y := blank % n.xMax
	stepList := make([]Step, 0, 4)
	if x > 0 && n.get(x-1, y) > 0 {
		stepList = append(stepList, Step{x - 1, y})
	}
	if y > 0 && n.get(x, y-1) > 0 {
		stepList = append(stepList, Step{x, y - 1})
	}
	if x < n.xMax-1 && n.get(x+1, y) > 0 {
		stepList = append(stepList, Step{x + 1, y})
	}
	if y < n.yMax-1 && n.get(x, y+1) > 0 {
		stepList = append(stepList, Step{x, y + 1})
	}
	return stepList
}

// Copy ...
func (n *NonZero) Copy() NonZero {
	newCb := make([]int64, len(n.chessboard))
	copy(newCb, n.chessboard)
	return NonZero{newCb, n.xMax, n.yMax}
}

// String ...
func (n NonZero) String() string {
	buf := bytes.Buffer{}
	for i := int64(0); i < n.yMax; i++ {
		for j := int64(0); j < n.xMax; j++ {
			buf.WriteString(fmt.Sprintf("%4d", n.chessboard[i*n.xMax+j]))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// Solve ...
func (n *NonZero) Solve() ([]string, error) {
	if _, ok := n.checkValid(); !ok {
		return nil, fmt.Errorf("not a game")
	}
	result, ok := n.solve()
	if !ok {
		return nil, fmt.Errorf("no result")
	}
	length := len(result)
	s := make([]string, length)
	for i, v := range result {
		s[length-i-1] = v
	}
	return s, nil
}

func (n *NonZero) solve() ([]string, bool) {
	index, ok := n.checkValid()
	if !ok {
		return nil, false
	}
	if n.checkEnd() {
		return []string{}, true
	}
	stepList := n.GetAllStep(index)
	for _, step := range stepList {
		// fmt.Println("=========== Step ===========")
		// fmt.Println(step)
		nn := n.Copy()
		nn.RunStep(step, index)
		// fmt.Println("============================")
		// fmt.Println(nn)
		result, ok := nn.solve()
		if ok {
			return append(result, step.String()), true
		}
	}
	return nil, false
}

func (n *NonZero) checkEnd() bool {
	zeroNum := int64(0)
	for _, v := range n.chessboard {
		if v == 0 {
			zeroNum++
		}
	}
	return zeroNum+1 == n.xMax*n.yMax
}

func (n *NonZero) checkValid() (index int64, _ bool) {
	minusNum := 0
	for i, v := range n.chessboard {
		if v < 0 {
			minusNum++
			index = int64(i)
		}
	}
	if minusNum != 1 {
		return index, false
	}
	return index, true
}
