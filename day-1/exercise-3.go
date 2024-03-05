package main

import "fmt"

type Employee interface {
	CalculateSalary() float64
}

type Fulltime struct {
	SalaryPerMonth float64
}

type Contractor struct {
	SalaryPerMonth float64
}

type Freelancer struct {
	HourlyRate  float64
	HoursWorked int
}

func (f Fulltime) CalculateSalary() float64 {
	return f.SalaryPerMonth
}

func (c Contractor) CalculateSalary() float64 {
	return c.SalaryPerMonth
}

func (fl Freelancer) CalculateSalary() float64 {
	return fl.HourlyRate * float64(fl.HoursWorked)
}

func main() {
	for {
		var employeeType string
		fmt.Println("Enter the type of employee (Fulltime, Contractor, or Freelancer) or 'quit' to exit:")
		_, err := fmt.Scan(&employeeType)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if employeeType == "quit" {
			break
		}

		var employee Employee
		switch employeeType {
		case "Fulltime":
			var salary float64
			fmt.Println("Enter the monthly salary for Full-time employee:")
			_, err := fmt.Scan(&salary)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			employee = Fulltime{SalaryPerMonth: salary}
		case "Contractor":
			var salary float64
			fmt.Println("Enter the monthly salary for Contractor:")
			_, err := fmt.Scan(&salary)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			employee = Contractor{SalaryPerMonth: salary}
		case "Freelancer":
			var hourlyRate float64
			var hoursWorked int
			fmt.Println("Enter the hourly rate for Freelancer:")
			_, err := fmt.Scan(&hourlyRate)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Enter the hours worked for Freelancer:")
			_, err = fmt.Scan(&hoursWorked)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			employee = Freelancer{HourlyRate: hourlyRate, HoursWorked: hoursWorked}
		default:
			fmt.Println("Invalid employee type entered.")
			continue
		}

		fmt.Println("Employee Type:", employeeType)
		fmt.Println("Employee Salary:", employee.CalculateSalary())
	}
}
