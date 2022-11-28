package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimaX(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 2, PrimeX(1), "Hasil harus sesuai")
		assert.Equal(t, 11, PrimeX(5), "Hasil harus sesuai")
		assert.Equal(t, 19, PrimeX(8), "Hasil harus sesuai")
		assert.Equal(t, 23, PrimeX(9), "Hasil harus sesuai")
		assert.Equal(t, 29, PrimeX(10), "Hasil harus sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.NotEqual(t, 1, PrimeX(1), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, PrimeX(5), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, PrimeX(8), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, PrimeX(9), "Hasil harus tidak sesuai")
		assert.NotEqual(t, 1, PrimeX(10), "Hasil harus tidak sesuai")
	})

}
