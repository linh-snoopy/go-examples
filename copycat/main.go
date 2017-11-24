package main


import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Role string
	Age  int32
	Test string
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name      *string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func main() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Test: "lalalala"}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Test: "lalalala"}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Test: "lalalala"}}
		employee  = Employee{}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("1st: %#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRule: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("2nd: %#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("3rd: %#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, DoubleAge: 60, EmployeId: 0, SuperRule: "Super Dev"},
	// }
}