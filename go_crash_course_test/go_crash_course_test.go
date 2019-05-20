package go_crash_course_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Tenjin0/go_crash_course/codewars/aretheythesame"
)

func dotest(array1 []int, array2 []int, exp bool) {
	var ans = aretheythesame.Comp(array1, array2)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Tour", func() {

	It("should handle basic cases", func() {

		var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		var a2 = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
		// dotest(a1, a2, true)
		// a2 = []int{231, 14641, 20736, 361, 25921, 361, 20736, 361}
		// dotest(a1, a2, false)
		// a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
		// dotest(a1, a2, false)
		a1 = []int{}
		// a2 = []int{}
		// dotest(a1, a2, true)
		a2 = []int{25, 49}
		dotest(a1, a2, false)

	})
})
