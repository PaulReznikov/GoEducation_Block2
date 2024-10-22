package main

import "fmt"

// Employee /////////////////////////////////////////////////
type Employee struct {
	Name         string
	Position     string
	Salary       float64
	Experience   int
	VacationDays int
}

// Manager /////////////////////////////////////////////////////
type Manager struct {
	Employee
	TeamSize int
	Bonus    float64
}

func (e *Employee) GetDetails() {
	fmt.Printf("Name: %v; Position: %v; Salary: %v; Experience: %v.\n\n", e.Name, e.Position,
		e.Salary, e.Experience)
}

func (e *Employee) CalculateSalary() {
	if e.Experience > 5 {
		e.Salary *= 1.1
	}
}

func (e *Employee) PromoteEmployee() error {
	positionsTop := map[int]string{1: "jun", 2: "mid", 3: "senor", 4: "team lead"}
	if positionsTop[4] == e.Position {
		return fmt.Errorf("сотрудник %v занимает максимальную должность - %v", e.Name, e.Position)
	}

	for k, v := range positionsTop {
		if v == e.Position {
			e.Position = positionsTop[k+1]
			return nil
		}
	}

	return fmt.Errorf("такой должности нет в компании - %v", e.Position)

}

func (e *Employee) Vacation() {
	vacationsList := map[string]int{"jun": 28, "mid": 35, "senor": 40, "team lead": 45}
	e.VacationDays = vacationsList[e.Position]
	if e.Experience > 5 {
		e.VacationDays += 5
	} else if e.Experience > 10 {
		e.VacationDays += 10
	} else if e.Experience > 15 {
		e.VacationDays += 15
	}

	fmt.Printf("Количество дней отпуска у работника %v = %v\n", e.Name, e.VacationDays)

}

////////////////////////////////////////////////

func (m *Manager) CalculateBonus() {
	m.Bonus = 30 * float64(m.TeamSize) * float64(m.Experience)
}

func (m *Manager) CalculateSalary() {
	if m.Experience > 5 {
		m.Salary = m.Salary*1.1 + m.Bonus
	}
}

func main() {
	m := Manager{
		Employee: Employee{
			"Dima",
			"mid",
			4000.0,
			5,
			0,
		},
		TeamSize: 5,
		Bonus:    0,
	}

	fmt.Println(m)
	m.GetDetails()
	m.CalculateBonus()
	m.Vacation()
	m.PromoteEmployee()
	fmt.Println(m)
}
