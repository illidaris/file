package path

import (
	"path/filepath"
	"testing"
)

func TestMkdirIfNotExist(t *testing.T) {
	p := "./test.txt"
	absPath, _ := filepath.Abs(p)
	t.Log("temp file is", absPath)
	defer func(path string) {
		err := Delete(path)
		if err != nil {
			t.Error(err)
		}
	}(p)
	err := MkdirIfNotExist(p)
	if err != nil {
		t.Error(err)
	}
	b, err := ExistOrNot(p)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error(p, " create failed")
	}
	t.Log(p, " create success")
}
