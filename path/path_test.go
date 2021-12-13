package path

import (
	"testing"
)

func TestExistOrNot(t *testing.T) {
	p := "./test.txt"
	b, err := ExistOrNot(p)
	if err != nil {
		t.Error(err)
	}
	if b {
		t.Error(p, " is exist")
	}
	t.Log(p, " is not exist")
}

func ExampleExistOrNot() {
	fPath := "./test.txt"
	b, err := ExistOrNot(fPath)
	if err != nil {
		println(err)
	}
	if b {
		println("file path is exist")
	}
}

func BenchmarkExistOrNot(b *testing.B) {
	//do something
}
