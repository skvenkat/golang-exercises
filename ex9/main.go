package main

import "fmt"

func checkLeapYear(year int) bool {
	if (year % 4) == 0 {
		if (year % 100) == 0 {
			if (year % 400) == 0 {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

func main() {
	years := []int{1897, 1900, 1927, 1947, 1965, 1977, 1981, 1990, 1996, 2000, 2006, 2010, 2014, 2016, 2020}
	for _, year := range years {
		fmt.Printf("is %d, leap year : %v\n", year, checkLeapYear(year))
	}

}
