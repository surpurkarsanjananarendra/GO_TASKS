package main

import ("fmt"
		"io")

type Employees struct {
	Eid    int
	Name   string
	Age    int
	Salary float64
	Dept   string
}

type Department struct {
	Dname string
	List  []Employees
}

func (d *Department) AddEmp(e Employees) {
	d.List = append(d.List, e)
}

func (d *Department) RemoveEmp(id int) {
	for i, emp := range d.List {
		if emp.Eid == id {
			d.List = append(d.List[:i], d.List[i+1:]...)
			return
		}
	}
}

func (d Department) AvgSal() float64 {
	if len(d.List) == 0 {
		return 0
	}
	sum := 0.0
	for _, emp := range d.List {
		sum += emp.Salary
	}
	return sum / float64(len(d.List))
}

func main() {
	var lt []Employees

	finance := Department{Dname: "Finance"}
	marketing := Department{Dname: "Marketing"}
	it := Department{Dname: "IT"}

	for {
		fmt.Println("\nMENU\n1. Add Employees\n2. Delete Employee\n3. Update Employe\n4. Display All\n5. Raise Salary\n6. Exit")

		fmt.Print("Enter choice: ")
		var ch int
		fmt.Scan(&ch)

		switch ch {
		case 1:
			var n int
			fmt.Print("Enter number of employees to add: ")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				var e Employees

				fmt.Print("\nEnter Eid: ")
				_, err := fmt.Scan(&e.Eid)
				if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
				fmt.Println("Input finished or unexpected end of input:", err)
				} else {
				fmt.Println("Error reading eid:", err)
				}
				}
				fmt.Println("Successfully read Eid:", e.Eid)

				fmt.Print("Enter Name: ")
				_, err = fmt.Scan(&e.Name)
				if err != nil {
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						fmt.Println("\nInput finished:", err)
						return
					}
					fmt.Println("Error reading name:", err)
					return
				}
				fmt.Println("Successfully read Name:", e.Name)

				fmt.Print("Enter Age: ")
				_, err = fmt.Scan(&e.Age)
				if err != nil {
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						fmt.Println("\nInput finished:", err)
						return
					}
					fmt.Println("Error reading Age:", err)
					return
				}
				fmt.Println("Successfully read Age:", e.Age)

				fmt.Print("Enter Salary: ")
				_, err = fmt.Scan(&e.Salary)
				if err != nil {
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						fmt.Println("\nInput finished:", err)
						return
					}
					fmt.Println("Error reading Salary:", err)
					return
				}
				fmt.Println("Successfully read Salary:", e.Salary)

				fmt.Print("Enter Department: ")
				_, err = fmt.Scan(&e.Dept)
				if err != nil {
					if err == io.EOF || err == io.ErrUnexpectedEOF {
						fmt.Println("\nInput finished:", err)
						return
					}
					fmt.Println("Error reading Dept:", err)
					return
				}
				fmt.Println("Successfully read Dept:", e.Dept)

				lt = append(lt, e)

				if e.Dept == "Finance" {
					finance.AddEmp(e)
				} else if e.Dept == "Marketing" {
					marketing.AddEmp(e)
				} else {
					it.AddEmp(e)
				}
			}
		case 2:
			var id int
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)

			finance.RemoveEmp(id)
			marketing.RemoveEmp(id)
			it.RemoveEmp(id)

			for i, emp := range lt {
				if emp.Eid == id {
					lt = append(lt[:i], lt[i+1:]...)
					break
				}
			}
		case 3:
    	var id int
    	fmt.Print("Enter Employee ID to update: ")
    	fmt.Scan(&id)
    
    	for i := range lt {
    		if lt[i].Eid == id {
    			oldDept := lt[i].Dept
    
    			fmt.Print("Enter New Name: ")
    			fmt.Scan(&lt[i].Name)
    
    			fmt.Print("Enter New Age: ")
    			fmt.Scan(&lt[i].Age)
    
    			fmt.Print("Enter New Salary: ")
    			fmt.Scan(&lt[i].Salary)
    
    			fmt.Print("Enter New Department: ")
    			fmt.Scan(&lt[i].Dept)
    
    			if oldDept == "Finance" {
    				finance.RemoveEmp(id)
    			} else if oldDept == "Marketing" {
    				marketing.RemoveEmp(id)
    			} else {
    				it.RemoveEmp(id)
    			}
				
    			if lt[i].Dept == "Finance" {
    				finance.AddEmp(lt[i])
    			} else if lt[i].Dept == "Marketing" {
    				marketing.AddEmp(lt[i])
    			} else {
    				it.AddEmp(lt[i])
    			}
    
    			fmt.Println("Updated Successfully")
    			break
    		}
    	}
		case 4:
			fmt.Println("\nALL EMPLOYEES")
			for _, e := range lt {
				fmt.Printf("\nID:%d Name:%s Age:%d Salary:%.2f Dept:%s",
					e.Eid, e.Name, e.Age, e.Salary, e.Dept)
			}

			fmt.Println("\n\nFinance:", finance.List)
			fmt.Println("Marketing:", marketing.List)
			fmt.Println("IT:", it.List)
    	case 5:
    	var id int
    	var percent float64
    
    	fmt.Print("Enter Employee ID: ")
    	fmt.Scan(&id)
    
    	fmt.Print("Enter Raise Percentage: ")
    	fmt.Scan(&percent)
    
    	for i := range lt {
    		if lt[i].Eid == id {
    
    			lt[i].Salary += lt[i].Salary * percent / 100
    			dept := lt[i].Dept
    			if dept == "Finance" {
    				for j := range finance.List {
    					if finance.List[j].Eid == id {
    						finance.List[j].Salary = lt[i].Salary
    					}
    				}
    			} else if dept == "Marketing" {
    				for j := range marketing.List {
    					if marketing.List[j].Eid == id {
    						marketing.List[j].Salary = lt[i].Salary
    					}
    				}
    			} else {
    				for j := range it.List {
    					if it.List[j].Eid == id {
    						it.List[j].Salary = lt[i].Salary
    					}
    				}
    			}
    			fmt.Println("Salary Raised!")
    			break
    		}
    	}
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