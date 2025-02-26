# Employement-Management-Systems

# Employee Management System

## Description
This program is a simple Employee Management System implemented in Golang. It allows adding employees, updating their salary, displaying their details, and removing employees from the system.

## Features
- Add new employees to the system.
- Update employee salaries by increasing them by 10%.
- Display employee details.
- Remove employees by their ID.

## Struct Definition
The program defines an `Employee` struct with the following fields:
```go
// Employee struct definition
type Employee struct {
    ID     int
    Name   string
    Salary float64
}
```

## Methods
### UpdateSalary
This method increases an employee's salary by 10%.
```go
func (e *Employee) UpdateSalary() {
    e.Salary *= 1.1 // Increasing salary by 10%
}
```

### Display
This method prints the details of an employee.
```go
func (e Employee) Display() {
    fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}
```

## Functions
### AddEmployee
Adds an employee to the employee database (a map of employees).
```go
func AddEmployee(emp *Employee, db map[int]*Employee) {
    db[emp.ID] = emp
}
```

### RemoveEmployee
Removes an employee from the database by ID.
```go
func RemoveEmployee(id int, db map[int]*Employee) {
    delete(db, id)
}
```

## Main Function Workflow
1. **Initialize a map** to store employees:
   - The map `emp` is created using `make(map[int]*Employee)`, where the key is the employee's ID, and the value is a pointer to an `Employee` struct.

2. **Create employee instances** with unique IDs, names, and salaries:
   - Three employees are created with sample details using struct literals.
   ```go
   employee1 := &Employee{ID: 234, Name: "Susmitha", Salary: 78000}
   employee2 := &Employee{ID: 456, Name: "Sush", Salary: 89000}
   employee3 := &Employee{ID: 678, Name: "Susmi", Salary: 98000}
   ```

3. **Add employees** to the map using `AddEmployee()`:
   - This function stores employee pointers in the map for quick access and modification.
   ```go
   AddEmployee(employee1, emp)
   AddEmployee(employee2, emp)
   AddEmployee(employee3, emp)
   ```

4. **Display employee details** before updating salaries:
   - A loop iterates over the map and calls the `Display()` method for each employee.

5. **Increase salary by 10%** for all employees using `UpdateSalary()`:
   - Another loop iterates over the map and calls `UpdateSalary()` on each employee pointer, modifying the salary directly.

6. **Display updated salaries**:
   - After updating salaries, the program prints the updated details.

7. **Remove an employee** using `RemoveEmployee()`:
   - The function `RemoveEmployee(456, emp)` removes the employee with ID `456` from the map using `delete()`.

8. **Display final employee details** after removal:
   - The remaining employees are displayed after deletion.


   ```
 The program will output:
   ```
   Employee Details Before Salary Update:
   ID: 234, Name: Susmitha, Salary: 78000.00
   ID: 456, Name: Sush, Salary: 89000.00
   ID: 678, Name: Susmi, Salary: 98000.00

   Employee Details After Salary Update:
   ID: 234, Name: Susmitha, Salary: 85800.00
   ID: 456, Name: Sush, Salary: 97900.00
   ID: 678, Name: Susmi, Salary: 107800.00

   Employee Details After Removing Employee 2:
   ID: 234, Name: Susmitha, Salary: 85800.00
   ID: 678, Name: Susmi, Salary: 107800.00
   ```

