package _038_count_and_say

import (
	"fmt"
	"testing"
)

type question38 struct {
	para38
	ans38
}

type para38 struct {
	n int
}

type ans38 struct {
	ans string
}

func Test_Problem38(t *testing.T) {

	qs := []question38{
		{
			para38{2},
			ans38{"11"},
		},
		{
			para38{3},
			ans38{"21"},
		},
		{
			para38{4},
			ans38{"1211"},
		},
		{
			para38{5},
			ans38{"111221"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 38------------------------\n")

	for _, q := range qs {
		a, p := q.ans38, q.para38
		if ans := countAndSay(p.n); ans != a.ans {
			t.Errorf("countAndSay(%d) expected be %s, but %s got", p.n, a.ans, ans)
		} else {
			fmt.Printf("【input】:%d\t【output】:%s\n", p.n, countAndSay(p.n))
		}
	}
	fmt.Printf("\n\n\n")
}
