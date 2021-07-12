package randomStrings

import "testing"


func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_= RandStr(32, AlphaNum)
	}
}


func TestRandStr(t *testing.T) {
	res := RandStr(32, AlphaNum)
	if len(res) == 0 {
		t.Fatal("no result")
	}

	if len(res) != 32 {
		t.Fatal("Length not 32")
	}
}