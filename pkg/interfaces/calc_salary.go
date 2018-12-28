package interfaces

// Based on: https://golangbot.com/interfaces-part-1/
// In Go, an interface is a set of method signatures.
// When a type provides definition for all the methods in the interface,
// it is said to implement the interface.
// Interface specifies what methods a type should have and the type decides
// how to implement these methods.

type SalaryCalculator interface {
	CalculateSalary() int
}

// A permanent employee has a basic pay & bonus
type Permanent struct {
	empId	string
	basicpay int
	bonus	int
}

// A contract employee has just a basic pay
type Contract struct {
	empId  string
	basicpay int
}

// salary of permanent employee is sum of basic pay and bonus
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.bonus
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// total expense is calculated by iterating though the SalaryCalculator
//  slice and summing the salaries of the individual employees.
// Since both permanent & contract employees have CalculateSalary()
// methods, they implement the SalaryCalculator interface.
func totalExpense(s []SalaryCalculator) int {
	var expense int = 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	return expense
}
