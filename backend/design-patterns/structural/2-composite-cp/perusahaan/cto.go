package perusahaan

import "fmt"

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	fmt.Sprintf("Total Divison Salary CTO %s", c.GetSalary())
	total := c.GetSalary()
	for _, subordinate := range c.Subordinate {
		total += subordinate.TotalDivisonSalary()
	}
	return total
}
