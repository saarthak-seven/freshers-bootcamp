package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type FullTime struct {
	monthlyPay int
}

func (f FullTime) CalculateSalary() int {
	return f.monthlyPay
}

type Contractor struct {
	dailyPay   int
	daysWorked int
}

func (c Contractor) CalculateSalary() int {
	return c.dailyPay * c.daysWorked
}

type Freelance struct {
	hourlyRate  int
	hoursWorked int
}

func (f Freelance) CalculateSalary() int {
	return f.hourlyRate * f.hoursWorked
}

func main() {

	fulltime := FullTime{monthlyPay: 15000}
	contractor := Contractor{dailyPay: 3000, daysWorked: 22}
	freelancer := Freelance{hourlyRate: 100, hoursWorked: 20}

	employees := []SalaryCalculator{fulltime, contractor, freelancer}

	for _, emp := range employees {
		fmt.Println("Calculated Salary:", emp.CalculateSalary())
	}

}
