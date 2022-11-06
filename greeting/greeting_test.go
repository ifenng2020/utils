package greeting

import "testing"

func TestHello(t *testing.T) {
	greet := Hello("zf")
	want := "Hello, zf"
	if greet != want {
		t.Errorf(`Hello("zf") expect %s, got %s`, want, greet)
	}
}

func TestHi(t *testing.T) {
	greet := Hi("zf")
	want := "Hi, zf"
	if greet != want {
		t.Errorf(`Hi("zf") expect %s, got %s`, want, greet)
	}
}
