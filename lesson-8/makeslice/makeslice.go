package makeslice

import "math/rand"

func MakeSlice(n int) []byte {
	var s = []byte{}
	for i := 0; i < n; i++ {
		s = append(s, byte(rand.Intn(100)))
	}
	return s
}
