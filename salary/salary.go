package salary

import "fmt"

type Employee interface {
	GetSalary()
}

type FullTimeEmployee struct {
	Name          string
	MonthlySalary int
}

type PartTimeEmployee struct {
	Name        string
	HourlyRate  int
	HoursWorked int
}

func (f FullTimeEmployee) GetSalary() {
	fmt.Printf("gaji karyawan fulltime dengan nama %s adalah %d\n", f.Name, f.MonthlySalary)
}

func (p PartTimeEmployee) GetSalary() {
	fmt.Printf("gaji karyawan partime dengan nama %s adalah %d\n", p.Name, p.HourlyRate*p.HoursWorked)
}

func ViewSalary(e Employee) {
	e.GetSalary()
}
