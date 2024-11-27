package makestring

import (
	m "lesson-8/makeslice"
)

func MakeString(n int) string {
	return string(m.MakeSlice(n))
}
