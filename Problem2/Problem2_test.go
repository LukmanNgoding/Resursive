package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 0, Fibonacci(0), "Hasil harus sesuai")
		assert.Equal(t, 1, Fibonacci(2), "Hasil harus sesuai")
		assert.Equal(t, 34, Fibonacci(9), "Hasil harus sesuai")
		assert.Equal(t, 55, Fibonacci(10), "Hasil harus sesuai")
		assert.Equal(t, 144, Fibonacci(12), "Hasil harus sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.NotEqual(t, 1, Fibonacci(0), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 0, Fibonacci(2), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, Fibonacci(9), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, Fibonacci(10), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, Fibonacci(12), "Hasil harus tidak sesuai")
	})

}
