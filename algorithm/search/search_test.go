package search

import (
	"math/rand"
	"sort"
	"testing"
)

func TestOrderStatistics(t *testing.T) {
	for i := 0; i < 10; i++ {
		perm := rand.Perm(10)
		ans := OrderStatistics(perm, i)
		if ans != i {
			t.Errorf("%d != %d", ans, i)
		}
	}
	for i := 0; i < 100000; i++ {
		perm := make([]int, 30)
		for j := 0; j < len(perm); j++ {
			perm[j] = rand.Intn(10)
		}
		ind := rand.Intn(30)
		ans := OrderStatistics(perm, ind)
		v := append([]int{}, perm...)
		sort.Ints(perm)
		if ans != perm[ind] {
			t.Errorf("%d != %d %v %v %d\n", ans, perm[ind], perm, v, ind)
		}
	}
	/*sl := make([]int, 100000)
	OrderStatistics(sl, 50000)*/
}
