package main

import (
	"fmt"

	"ex10/open_close"
	"ex10/single_resp"
)

func main() {
	fmt.Println("SOLID Principles")
	single_resp.SingleResp()
	open_close.OpenCloseResp()

}
