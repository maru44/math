package problem

const Plus = "1"
const Minus = "2"
const Times = "3"

const Round = 20

type Problem struct {
	Correct int
	Count   int
}

func (p *Problem) Rotate(method string) {
	switch method {
	case Plus:
		p.Count++
	}
}
