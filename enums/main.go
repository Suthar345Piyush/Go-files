package main

import "fmt"

// enums  are  enumerated types in go
// // we can use it  with type  and  const keyword with iota
// type OrderStaus int

//for string enums

type OrderStaus string

const (
	Received   OrderStaus = "received"
	Processing            = "processing"
	Shipped               = "shipped"
	Deliverd              = "deliverd"
	Cancelled             = "cancelled"
)

// const (
// 	Received OrderStaus = iota
// 	Processing
// 	Shipped
// 	Deliverd
// 	Cancelled
// )

func changeOrderStatus(status OrderStaus) {
	fmt.Println("order status is:", status)
}

func main() {
	changeOrderStatus(Shipped)
}
