package Problem0532

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0532(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2, 3, 4, 5}, -1},
			ans{0},
		},

		question{
			para{[]int{3, 1, 4, 1, 5}, 2},
			ans{2},
		},

		question{
			para{[]int{1, 2, 3, 4, 5}, 1},
			ans{4},
		},

		question{
			para{[]int{3, 1, 4, 1, 5}, 0},
			ans{1},
		},

		question{
			para{[]int{3, 3, 5, 5, 5}, 2},
			ans{1},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findPairs(p.one, p.k), "输入:%v", p)
	}
}
