package math

// I am aware that these exist ynder the package: https://pkg.go.dev/golang.org/x/exp/constraints
// But to facilitate learning and not relying on imports I'm implementing them below.

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

func DifferenceInts[T Integer](a, b T) int {
	if a > b {
		return int(a - b)
	} else {
		return int(b - a)
	}
}
