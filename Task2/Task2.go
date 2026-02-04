package main

import (
	"fmt"
	"io"
)

type Employees struct { //can be accessed from other packages
	Eid    int
	Name   string
	Age    int
	Salary float64
	Dept   string
}

type Department struct { //public...
	Dname string
	List  []Employees
}

func (d *Department) AddEmp(e Employees) { //public..
	d.List = append(d.List, e)
}

func (d *Department) RemoveEmp(id int) bool { //public
	for i, emp := range d.List {
		if emp.Eid == id {
			d.List = append(d.List[:i], d.List[i+1:]...) //(variadic expression) take ele one by one from slice   [without ... go thinks we append 1 ele not all]
			return true
		}
	}
	return false
}

func (d Department) AvgSal() float64 { //public
	if len(d.List) == 0 {
		return 0
	}
	sum := 0.0
	for _, emp := range d.List {
		sum += emp.Salary
	}
	return sum / float64(len(d.List))
}

func scan_input(err error) { //private, limited for this package only
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			fmt.Println("\nInput finished or unexpected end of input:", err)
		} else {
			fmt.Println("Input error:", err)
		}
	}
}

func empExists(list []Employees, id int) bool { // private (it only checks if emp is there or not, doesnt care about its location, eg:add)
	for _, e := range list {
		if e.Eid == id {
			return true
		}
	}
	return false
}

func findEmpIndex(list []Employees, id int) int { //private (it cares about position eg: update)
	for i, e := range list {
		if e.Eid == id {
			return i
		}
	}
	return -1
}

func getDeptPtr(dept string, finance, marketing, it *Department) *Department { //private (to avoid repetitive if-else block)
	if dept == "Finance" {
		return finance
	} else if dept == "Marketing" {
		return marketing
	}
	return it //as it uses ptr it modifies the original value
}

func updateDeptSalary(d *Department, id int, sal float64) { //private
	for i := range d.List {
		if d.List[i].Eid == id {
			d.List[i].Salary = sal
			return
		}
	}
}

func main() {
	var lt []Employees
	var err error

	finance := Department{Dname: "Finance"}
	marketing := Department{Dname: "Marketing"}
	it := Department{Dname: "IT"}

	for {
		fmt.Println("\nMENU\n1. Add Employees\n2. Delete Employee\n3. Update Employe\n4. Display All\n5. Raise Salary\n6. Exit")
		fmt.Print("\nEnter choice: ")

		var ch int
		fmt.Scan(&ch)

		switch ch {

		case 1:
			var n int
			fmt.Print("\nEnter number of employees to add: ")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				var e Employees

				fmt.Print("\nEnter Eid: ")
				fmt.Scan(&e.Eid)
				if empExists(lt, e.Eid) {
					fmt.Println("Error: Employee ID already exists!")
					continue
				}

				fmt.Print("Enter Name: ")
				_, err = fmt.Scan(&e.Name)
				scan_input(err)

				fmt.Print("Enter Age: ")
				_, err = fmt.Scan(&e.Age)
				scan_input(err)

				fmt.Print("Enter Salary: ")
				_, err = fmt.Scan(&e.Salary)
				scan_input(err)

				fmt.Print("Enter Department: ")
				_, err = fmt.Scan(&e.Dept)
				scan_input(err)

				lt = append(lt, e)
				getDeptPtr(e.Dept, &finance, &marketing, &it).AddEmp(e)
				fmt.Println("Employee Added Successfully")
			}

		case 2:
			var id int
			fmt.Print("\nEnter Employee ID to delete: ")
			fmt.Scan(&id)

			if !empExists(lt, id) {
				fmt.Println("Error: Employee not found!")
				break
			}

			idx := findEmpIndex(lt, id)
			dept := lt[idx].Dept
			lt = append(lt[:idx], lt[idx+1:]...)
			getDeptPtr(dept, &finance, &marketing, &it).RemoveEmp(id)
			fmt.Println("Employee Deleted Successfully")

		case 3:
			var id int
			fmt.Print("\nEnter Employee ID to Update: ")
			fmt.Scan(&id)

			idx := findEmpIndex(lt, id)
			if idx == -1 {
				fmt.Println("Error: Employee not found!")
				break
			}

			oldDept := lt[idx].Dept

			fmt.Print("Enter New Name: ")
			fmt.Scan(&lt[idx].Name)

			fmt.Print("Enter New Age: ")
			fmt.Scan(&lt[idx].Age)

			fmt.Print("Enter New Salary: ")
			fmt.Scan(&lt[idx].Salary)

			fmt.Print("Enter New Department: ")
			fmt.Scan(&lt[idx].Dept)

			getDeptPtr(oldDept, &finance, &marketing, &it).RemoveEmp(id)
			getDeptPtr(lt[idx].Dept, &finance, &marketing, &it).AddEmp(lt[idx])

			fmt.Println("Updated Successfully")

		case 4:
			if len(lt) == 0 {
				fmt.Println("No Employees to Display")
				break
			}

			fmt.Println("\nALL EMPLOYEES")
			for _, e := range lt {
				fmt.Printf("\nID:%d Name:%s Age:%d Salary:%.2f Dept:%s",
					e.Eid, e.Name, e.Age, e.Salary, e.Dept)
			}

		case 5:
			var id int
			var percent float64

			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)

			idx := findEmpIndex(lt, id)
			if idx == -1 {
				fmt.Println("Error: Employee Not Found!")
				break
			}

			fmt.Print("Enter Raise Percentage: ")
			fmt.Scan(&percent)

			lt[idx].Salary += lt[idx].Salary * percent / 100
			updateDeptSalary(
				getDeptPtr(lt[idx].Dept, &finance, &marketing, &it),
				id,
				lt[idx].Salary,
			)

			fmt.Println("Salary Raised Successfully")

		case 6:
			fmt.Println("\nDepartment Averages:")
			fmt.Printf("Finance: %.2f\n", finance.AvgSal())
			fmt.Printf("Marketing: %.2f\n", marketing.AvgSal())
			fmt.Printf("IT: %.2f\n", it.AvgSal())
			fmt.Println("Exiting Program...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
