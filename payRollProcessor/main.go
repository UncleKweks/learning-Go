package main

import (
	"fmt"
)

type Payable interface {
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se *SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12
}

func (se *SalariedEmployee) GetName() string {
	return fmt.Sprintf("Salaried: %s (Annual Salary: %.2f) Employee: %s", se.Name, se.AnnualSalary, se.Name)
}


type HourlyEmployee struct {
	Name       string
	HourlyRate float64
	HoursWorked float64
}

func (he *HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he *HourlyEmployee) GetName() string {
	return fmt.Sprintf("Hourly: %s (Hourly Rate: %.2f) Employee: %s", he.Name, he.HourlyRate, he.Name)
}

type CommissionedEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	Sales          float64
}

func (ce *CommissionedEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.Sales)
}

func (ce *CommissionedEmployee) GetName() string {
	return fmt.Sprintf("Commissioned: %s (Base Salary: %.2f, Commission Rate: %.2f) Employee: %s", ce.Name, ce.BaseSalary, ce.CommissionRate, ce.Name)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("Processing Payroll:")
	totalPayroll := 0.0
	for _, employee := range employees {
		pay := employee.CalculatePay()
		totalPayroll += pay
		fmt.Printf("%s - Pay: %.2f\n", employee.(interface{ GetName() string }).GetName(), pay)
	}
	fmt.Printf("Total Payroll: %.2f\n", totalPayroll)
}

func main() {
	employees := []Payable{
		&SalariedEmployee{Name: "Alice", AnnualSalary: 60000},
		&HourlyEmployee{Name: "Bob", HourlyRate: 20, HoursWorked: 160},
		&CommissionedEmployee{Name: "Charlie", BaseSalary: 30000, CommissionRate: 0.1, Sales: 50000},
	}

	payRollList := []Payable{
		&SalariedEmployee{Name: "David", AnnualSalary: 80000},
		&HourlyEmployee{Name: "Eve", HourlyRate: 25, HoursWorked: 120},
		&CommissionedEmployee{Name: "Frank", BaseSalary: 40000, CommissionRate: 0.15, Sales: 60000},
	}

	ProcessPayroll(employees)
	ProcessPayroll(payRollList)
}