package commondenominators

// kata: https://www.codewars.com/kata/54d7660d2daf68c619000d95
// example: https://www.youtube.com/watch?v=WdsvijS5MRA
// lcm: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// Pgcd is used to find the pgcd between two numbers
func Pgcd(a, b int) int {

	var sup, inf int

	if a > b {
		sup = a
		inf = b
	} else if b > a {
		sup = b
		inf = a
	}

	var rest int

	for {
		rest = sup % inf
		if rest == 0 {
			break
		}
		sup = inf
		inf = rest
	}

	return inf
}

func ConvertFracts(a [][]int) string {
	// your code
	return ""
}
