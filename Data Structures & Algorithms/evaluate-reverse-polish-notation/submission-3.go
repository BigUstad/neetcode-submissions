import (
	"slices"
)

var opermap = map[string]bool {
	"+" : true,
	"-" : true,
	"*" : true,
	"/" : true,
}

func doEval(o1, o2 int, oper string) int {
	res := 0
	// eq := strconv.Itoa(o1) + oper + strconv.Itoa(o2)
	// fmt.Println(eq)
	switch(oper) {
		case "+":
		{
			res = o1 + o2
		}
		case "-":
		{
			res = o1 - o2
		}
		case "*":
		{
			res = o1 * o2
		}
		case "/":
		{
			res = (o1 / o2)
		}
	}
	// fmt.Println(res)
	return res
}

func evalRPN(tokens []string) int {
	res := 0
	if len(tokens) == 0 {
		return res
	}
	if len(tokens) < 3 {
		res, _ = strconv.Atoi(tokens[0])
		return res
	}
	stack := make([]int, 0)
	for _, t := range tokens {
		if !opermap[t] {
			i, _ := strconv.Atoi(t)
			stack = append(stack,i )
			continue
		}
		n := len(stack)
		i, j := n - 1, n - 2
		o1 := stack[j]
		o2 := stack[i]
		stack = slices.Delete(stack, j, i+1)
		res = doEval(o1, o2, t)
		stack = append(stack, res)
	}
	return res
}
