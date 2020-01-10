package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	custom := now.Format("2009-01-02 15:04:05")
	fmt.Println(now)
	fmt.Println(custom)

	ansic := now.Format(time.ANSIC)
	fmt.Println(ansic)

	aa := now.AddDate(0, 0, -1)
	fmt.Println(aa.String())
	fmt.Println(aa.Year())
	fmt.Println(int(aa.Month()))
	fmt.Println(aa.Day())
	fmt.Printf("%T\n", string(aa.Format("20060102")))
	fmt.Println("")
	fmt.Println("")

	var m, _ = strconv.Atoi("01")
	fmt.Println(m)
	fmt.Println(time.Month(m))

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	curr_dt := time.Now()
	var_year := curr_dt.Year()
	var_month := string(int(curr_dt.Month()))
	var_day := curr_dt.Day()

	fmt.Println(var_year, var_month, var_day)
	fmt.Printf("%T, %T, %T \n", var_year, var_month, var_day)

	currentTime := time.Now()
	fmt.Println("Current Time in String: ", currentTime.String())
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("20060102"))

	tmp := now.AddDate(0, 0, -1)
	fmt.Println("Current Time in String: ", tmp.String())
	fmt.Println("YYYY-MM-DD : ", tmp.Format("20060102"))

	v := reflect.ValueOf(row)
	values := make([]interface{}, v.NumField())
	types := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()    // get value from struct
		types[i] = v.Field(i).Type().String() // get type from struct
	}

}
