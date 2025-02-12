// time package  in golang

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	currentTime := time.Now()
// 	fmt.Println(currentTime)

// 	// datatype of time in golang
// 	fmt.Printf("type of time is %T\n", currentTime)

// 	// format function to get the  data into particular format

// 	// formatted := currentTime.Format("02-01-2006 , Monday , 15:04:05")
// 	// fmt.Println("the correct data is : ", formatted)

// 	formatted := currentTime.Format("02-01-2006 , Monday , 3:04 AM")
// 	fmt.Println("the correct data is : ", formatted)

// 	//parsing
// 	//-> in case of  parsing the  format of layout_str and  date_str should  be same

// 	layout_str := "2006-01-02"
// 	dateStr := "2024-01-07"
// 	formatted_time, _ := time.Parse(layout_str, dateStr)
// 	fmt.Println("formatted time: ", formatted_time)

// 	// add 1 more day in currenttime
// 	new_date := currentTime.Add(48 * time.Hour)
// 	fmt.Println("new_date time: ", new_date)
// 	formatted_new_date := new_date.Format("2006/01/02 Monday")
// 	fmt.Println("formatted_new_date time: ", formatted_new_date)

// }

//Time adding duration

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	futureTime := currentTime.Add(10 * time.Hour)
	fmt.Println("current time : ", currentTime)
	fmt.Println("the  future time is: ", futureTime)
}
