package lesson03

import "testing"

func HelloTest(t *testing.T){
	want := "Hello,word!"
	if got := Hello(); got != want{
		t.Errorf("Hello()= %q,want %q",got, want)
	}
}