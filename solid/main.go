package main

import (
	"fmt"

	"solid/dep_inv"
	"solid/intf_seg"
	"solid/liskov_sub"
	"solid/open_close"
	"solid/single_resp"
)

func main() {

	fmt.Println("SOLID Principles")
	fmt.Println("1) Single Responsibility:")
	single_resp.SingleResp()
	fmt.Println("2) Open Close:")
	open_close.OpenClose()
	fmt.Println("3) Liskov Substitution:")
	liskov_sub.LiskovSub()
	fmt.Println("4) Interface Segregation")
	intf_seg.IntfSeg()
	fmt.Println("5) Dependency Inversion")
	dep_inv.DepInv()
}
