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

type DeptTemp struct {
	Min int
	Max int
}

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
		if _, err := fmt.Scan(&numEmps); err != nil {
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

			if val, err := dept.CurrentTemp(); err == nil {
				fmt.Println(val)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
