package calculations

import (
	"errors"
	"fmt"
	"os"
)

// Calculating n! with or without logs
func Calculate(n int64, f bool) int64 {
	var result int64
	var err error
	if f == true {
		fmt.Println("Start calculations...")
		fmt.Println("Calculate <", n, ">!")
		result, err = fact(n)
		if err != nil {
			fmt.Println("Calculation error: ", err)
			os.Exit(1)
		} else {
			fmt.Println("Calculations complete!")
		}
	} else {
		result, err = fact(n)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return result
}

// Calculating n!
func fact(n int64) (int64, error) {
	var result int64 = 1
	var err error
	if n < 0 {
		err = errors.New("Invalid value for calculating")
	} else if n == 0 {
		result = 1
		err = nil
	} else {
		for i := int64(1); i <= n; i++ {
			result = result * i
		}
		err = nil
	}
	return result, err
}
