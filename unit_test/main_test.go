package unit_test

import (
	"testing"

	Op "github.com/rizkysr90/kalkulator_cli/helper"
	"github.com/stretchr/testify/assert"
)

var tester = []float32{6, 4, 2}

func TestSum(t *testing.T) {
	result := Op.SumCalc(tester)
	assert.Equal(t, float32(12), result)
}

func TestSubtract(t *testing.T) {
	result := Op.Subtract(tester)
	assert.Equal(t, float32(0), result)
}

func TestDivide(t *testing.T) {
	result := Op.Divide(tester)
	assert.Equal(t, float32(0.75), result)
}

func TestMultiply(t *testing.T) {
	result := Op.Multiply(tester)
	assert.Equal(t, float32(48), result)
}
