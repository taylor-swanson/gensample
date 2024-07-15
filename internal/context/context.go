package context

import "math/rand/v2"

type Context struct {
	Rand     *rand.Rand
	Current  int
	Total    int
	Progress float64
}
