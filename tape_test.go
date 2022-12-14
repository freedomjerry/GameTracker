package poker

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	NewFileContents, _ := ioutil.ReadAll(file)

	got := string(NewFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
