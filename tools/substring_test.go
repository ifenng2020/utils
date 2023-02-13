package tools

import "testing"

type Case struct {
	input  string
	expect string
	got    string
}

func Test(t *testing.T) {
	c1 := Case{
		input:  "Hello,世界",
		expect: "",
	}

	if c1.got = SubString(c1.input, 0, 88); c1.got != "Hello,世界" {
		t.Fatal(c1.got)
	}
}
