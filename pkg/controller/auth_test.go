package controller

import (
	"testing"
)

func BenchmarkCreateToken(b *testing.B) {
	_,_ = createToken("test", "client")
}

func BenchmarkHashPassword(b *testing.B) {
	_,_ = HashPassword("test")
}

func BenchmarkIsPasswordValid(b *testing.B) {
	_ = IsPasswordValid("test", "test")
	
}
