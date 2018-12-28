package interfaces

import (
	"gotest.tools/assert"
	"testing"
)

func TestTotalExpense(t *testing.T) {
	// Test examples
	tables := []struct {
		emps []SalaryCalculator
		exp	int
	}{
		// single contract employee
		{
			[]SalaryCalculator{
				Contract{"a", 400},
			},
			400,
		},
		// single permanent employee
		{
			[]SalaryCalculator{
				Permanent{"b", 500, 60},
			},
			560,
		},
		// permanent & contract employees
		{
			[]SalaryCalculator{
				Contract{"a", 400},
				Permanent{"b", 500, 60},
			},
			960,
		},
	}

	// go through each test example
	for _, table := range tables {
		exp := totalExpense(table.emps)
		assert.Equal(t, exp, table.exp)
	}
}