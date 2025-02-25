/*Assignment 1: Employee Management System
Use these: Structs, Methods, Pointers, Slices, Maps
Problem Statement:Create a simple Employee Management System that allows adding employees, updating their salary, and displaying their details.
Requirements:
1. Define a struct Employee with fields:
    * ID (int)
    * Name (string)
    * Salary (float64)
2. Implement methods:
    * UpdateSalary(newSalary float64) → Updates the salary of an employee.
    * Display() → Prints employee details.
3. Use a map to store employees (map[int]*Employee).
4. Implement functions:
    * AddEmployee(emp *Employee, db map[int]*Employee) → Adds an employee to the map.
    * RemoveEmployee(id int, db map[int]*Employee) → Removes an employee by ID.*/

package main

import "fmt"

// Employee struct definition
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Method to increase an employee's salary by 10%
func (e *Employee) UpdateSalary() {
	e.Salary *= 1.1 // Increasing salary by 10%
}

// Method to display employee details
func (e Employee) Display() {
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

// Function to add an employee to the map
func AddEmployee(emp *Employee, db map[int]*Employee) {
	db[emp.ID] = emp
}

// Function to remove an employee from the map
func RemoveEmployee(id int, db map[int]*Employee) {
	delete(db, id)
}

func main() {
	// Initializing the map to store employees
	emp := make(map[int]*Employee)

	// Creating employee instances
	employee1 := &Employee{ID: 234, Name: "Susmitha", Salary: 78000}
	employee2 := &Employee{ID: 456, Name: "Sush", Salary: 89000}
	employee3 := &Employee{ID: 678, Name: "Susmi", Salary: 98000}

	// Adding employees to the map
	AddEmployee(employee1, emp)
	AddEmployee(employee2, emp)
	AddEmployee(employee3, emp)

	// Displaying employee details before salary update
	fmt.Println("Employee Details Before Salary Update:")
	for _, emp := range emp {
		emp.Display()
	}

	// Updating salaries for all employees using the method itself
	for _, emp := range emp {
		emp.UpdateSalary()
	}

	// Displaying employee details after salary update
	fmt.Println("\nEmployee Details After Salary Update:")
	for _, emp := range emp {
		emp.Display()
	}

	// Removing an employee (example: removing employee2)
	RemoveEmployee(456, emp)

	// Displaying employee details after removal
	fmt.Println("\nEmployee Details After Removing Employee 2:")
	for _, emp := range emp {
		emp.Display()
	}
}

/*

	)
	//created a structure for employee
		type Employee struct{
			ID int
			Name string
			Salary float64
		}
	// created a method to update a salary with 10% raise
		func (e *Employee) UpdatedSalary(newSalary float64) {
			return e.Salary * 1.1 // Assumed 10% raise
		}

		//Method to display employee details
		func(e Employee) Display(){
			fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
	}
	// Function to add an employee to the map
	func AddEmployee(emp *Employee, db map[int]*Employee) {
		db[emp.ID] = emp
	}

	// Function to remove an employee from the map
	func RemoveEmployee(id int, db map[int]*Employee) {
		delete(db, id)
	}



		func main() {

		//Map to store employee
		//Declared the map
	 var emp map[int]*Employee
	 //Initialized the map
	  emp = make(map[int]*Employee)
			//created an instance for employee
			employee1 := Employee{
				ID : 234,
				Name : "Susmitha",
				Salary : 78000,
			}
	//Another way of creating an instance for struct
			employee2 := Employee{456, "Sush", 89000}
			employee3 := Employee{678, "Susmi", 98000}

			// Adding employees to the map
		AddEmployee(employee1, emp)
		AddEmployee(employee2, emp)
		AddEmployee(employee3, emp)
			// Displaying employee details before salary update
		fmt.Println("Employee Details Before Salary Update:",employee1, employee2,employee3)
	// Updating salaries for employees
	employee1.UpdatedSalary(employee1.Salary * 1.1) // 10% raise
	employee2.UpdatedSalary(employee2.Salary * 1.1) // 10% raise
	employee3.UpdatedSalary(employee3.Salary * 1.1) // 10% raise
	 // Displaying employee details after salary update
	 fmt.Println("\nEmployee Details After Salary Update:")
	 for id, emp := range emp {
		fmt.Printf("Employee ID: %d → ", id)
		emp.Display()
	}


	 // Removing an employee (example: removing employee2)
	 RemoveEmployee(456, emp)

	 // Displaying employee details after removal
	 fmt.Println("\nEmployee Details After Removing Employee 2:")
	 for _, emp := range emp {
		 emp.Display()
	 }

		}
*/
