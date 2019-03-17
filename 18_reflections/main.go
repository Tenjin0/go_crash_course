package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) string {

	if reflect.ValueOf(q).Kind() == reflect.Struct {

		t := reflect.TypeOf(q).Name()
		v := reflect.ValueOf(q)
		query := fmt.Sprintf("instert int %s values(", t)

		for i := 0; i < v.NumField(); i++ {

			switch v.Field(i).Kind() {

			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())

				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())

				}
			default:
				fmt.Println("Unsupported type")
				return ""
			}

		}
		query = fmt.Sprintf("%s)", query)
		return query
	}
	fmt.Println("Unsupported types")
	return ""
}

// func createQuery(o order) string {

// 	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordID, o.customerID)
// 	return i
// }

func main() {

	o := order{1234, 567}

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Science Park Road, Singapore",
		salary:  90000,
		country: "Singapore",
	}

	fmt.Println(createQuery(o))
	fmt.Println(createQuery(e))
}
