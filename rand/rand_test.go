package rand

import (
	"fmt"
	"testing"
)

// TestRand
func TestRandBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomBytes(10))
	}
}

func TestRandInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomInt(0, 0))
	}
}

// TestRandString ...
func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomString(20))
	}
}

func TestRandomCode(t *testing.T) {
	res := RandomCode(6)
	fmt.Println(res)
}
