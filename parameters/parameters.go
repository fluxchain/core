package parameters

var current *Parameters = nil

type Parameters struct {
	MinimumPoW [32]byte
}

func IsSet() bool {
	return current != nil
}

func Set(p *Parameters) {
	current = p
}

func Current() *Parameters {
	if !IsSet() {
		panic("no chain parameters set!")
	}

	return current
}
