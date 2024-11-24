package main

import (
	"fmt"
)

type Employee struct {
	ID           int
	Name         string
	DepartmentID int
	Projects     []string
}

type Department struct {
	ID        int
	Name      string
	Employees []Employee
}

type Project struct {
	ID           int
	Name         string
	DepartmentID int
	Employees    []Employee
}

// Function to create and return the Employee entity
func createEmployee(ID int, Name string, DepartmentID int, Projects []string) Employee {
	return Employee{ID, Name, DepartmentID, Projects}
}

// Function to create and return the Project entity
func createProject(ID int, Name string, DepartmentID int, Employees []Employee) Project {
	return Project{ID, Name, DepartmentID, Employees}
}

// Function to print employees according to department
func printEmployeesByDepartment(d Department) {
	fmt.Printf("Employees of %s department: %v\n", d.Name, d.GetEmployees())
}

// Function to print project information
func printProject(p Project) {
	fmt.Printf("Project: %s \n", p.Name)
	fmt.Print("Employees: ")
	for _, v := range p.Employees {
		fmt.Printf("%s ", v.Name)
	}
	fmt.Println()
	fmt.Println("-------------------------")
}

// Function to print employees
func printEmployee(e Employee) {
	fmt.Printf("Employee: %s \n", e.Name)
	fmt.Printf("DepartmentID: %d \n", e.DepartmentID)
	fmt.Print("Projects: ")
	for _, v := range e.Projects {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
	fmt.Println("-------------------------")
}

/****** Methods ******/

// Method of Employee struct
func (e *Employee) AssignToProject(project *Project) {
	e.Projects = append(e.Projects, project.Name)
	project.AddEmployee(e)
}

// Methods of Department struct
func (d *Department) AddEmployee(employee *Employee) {
	d.Employees = append(d.Employees, *employee)
}
func (d *Department) GetEmployees() []Employee {
	return d.Employees
}

// Method of Project struct
func (p *Project) AddEmployee(employee *Employee) {
	p.Employees = append(p.Employees, *employee)
}

func main() {
	backendDepartment := Department{1, "Backend", []Employee{}}
	frontendDepartment := Department{2, "Frontend", []Employee{}}
	aiDepartment := Department{3, "AI", []Employee{}}
	mobileDepartment := Department{4, "Mobile", []Employee{}}

	employee1 := createEmployee(1, "John", 1, []string{})
	employee2 := createEmployee(2, "Alice", 2, []string{})
	employee3 := createEmployee(3, "Michael", 3, []string{})
	employee4 := createEmployee(4, "Viktor", 4, []string{})
	employee5 := createEmployee(5, "Andrew", 2, []string{})
	employee6 := createEmployee(6, "Tim", 1, []string{})

	project1 := createProject(1, "ASDUM", backendDepartment.ID, []Employee{})
	project2 := createProject(2, "CDP", frontendDepartment.ID, []Employee{})
	project3 := createProject(3, "E-pay", backendDepartment.ID, []Employee{})
	project4 := createProject(4, "TranslatorApp", aiDepartment.ID, []Employee{})

	// Assign employees to department
	// When I created employee I already assigned to specific department

	backendDepartment.AddEmployee(&employee1)
	backendDepartment.AddEmployee(&employee6)

	frontendDepartment.AddEmployee(&employee2)
	frontendDepartment.AddEmployee(&employee5)

	aiDepartment.AddEmployee(&employee3)

	mobileDepartment.AddEmployee(&employee4)

	// Printing all employees according to assigned department
	printEmployeesByDepartment(backendDepartment)
	printEmployeesByDepartment(frontendDepartment)
	printEmployeesByDepartment(aiDepartment)
	printEmployeesByDepartment(mobileDepartment)

	// Assign employees to project ASDUM
	employee1.AssignToProject(&project1)
	employee2.AssignToProject(&project1)

	// Assign employees to project CDP
	employee6.AssignToProject(&project2)
	employee5.AssignToProject(&project2)

	// Assign employees to project E-pay
	employee6.AssignToProject(&project3)
	employee4.AssignToProject(&project3)

	// Assign employees to project TranslatorApp
	employee3.AssignToProject(&project4)
	employee4.AssignToProject(&project4)

	// Print projects
	fmt.Println()
	printProject(project1)
	printProject(project2)
	printProject(project3)
	printProject(project4)

	// Print employee
	fmt.Println()
	printEmployee(employee1)
	printEmployee(employee2)
	printEmployee(employee3)
	printEmployee(employee4)
	printEmployee(employee5)
	printEmployee(employee6)
}
