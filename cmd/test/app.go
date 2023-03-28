package main

func main() {

	tt := new(TT)
	aa := new(AA)
	aa.a = 1
	tt.ta = &AA{a: 1}
	str := "aaa"
	tt.b = &str
	ga(aa)
}

type TT struct {
	a  int
	b  *string
	ta *AA
}

type AA struct {
	a int
}

func b(c TT) {
	c.a = 1
	*c.b = "4"
}

func ga(a AA) {
	a.a = 1
}
