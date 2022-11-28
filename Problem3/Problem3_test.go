package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimaSegiEmpat(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "17\t19\t\n23\t29\t\n31\t37\t\n156", PrimaSegiEmpat(2, 3, 13), "Hasil harus sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.NotEqual(t, "17\t19\t\n23\t29\t\n31\t37\t\n156", PrimaSegiEmpat(5, 2, 1), "Hasil harus tidak sesuai")
	})

}
