package main

import (
	"errors"
	"fmt"
)

const (
	TempMin        = 15
	TempMax        = 30
	MinDepartments = 1
	MaxDepartments = 1000
)

var (
	ErrBadOperator     = errors.New("bad operator")
	ErrTempOutOfRange  = errors.New("temperature out of range")
	ErrReadDepartments = errors.New("failed to read departments")
	ErrReadEmployees   = errors.New("failed to read employees")
)

// DeptTemp يمثل مجال درجات الحرارة لكل قسم
type DeptTemp struct {
	Min int
	Max int
}

// UpdateTemp يقوم بتحديث الحد الأعلى أو الأدنى حسب رأي الموظف
func (d *DeptTemp) UpdateTemp(op string, temp int) error {
	switch op {
	case ">=":
		if temp > d.Min {
			d.Min = temp
		}
	case "<=":
		if temp < d.Max {
			d.Max = temp
		}
	default:
		return ErrBadOperator
	}

	return nil
}

// CurrentTemp يعطي درجة الحرارة الحالية أو -1 إذا كانت غير صالحة
func (d *DeptTemp) CurrentTemp() (int, error) {
	if d.Min > d.Max {
		return -1, ErrTempOutOfRange
	}

	return d.Min, nil
}

func main() {
	var numDeps int
	if _, err := fmt.Scan(&numDeps); err != nil || numDeps < MinDepartments || numDeps > MaxDepartments {
		fmt.Println("Error:", ErrReadDepartments)
		return
	}

	for i := 0; i < numDeps; i++ {
		var numEmps int
		if _, err := fmt.Scan(&numEmps); err != nil || numEmps < 1 {
			fmt.Println("Error:", ErrReadEmployees)
			return
		}

		dept := DeptTemp{Min: TempMin, Max: TempMax}

		for j := 0; j < numEmps; j++ {
			var op string
			var temp int
			if _, err := fmt.Scan(&op, &temp); err != nil {
				fmt.Println("Error: invalid input")
				return
			}

			if err := dept.UpdateTemp(op, temp); err != nil {
				fmt.Println("Error:", err)
				return
			}

			val, err := dept.CurrentTemp()
			if err == nil {
				fmt.Println(val)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
