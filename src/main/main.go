package main

import "module/ir"

func main() {
	ir.OpenUsbDevice(ir.VENDOR_ID, ir.PRODUCT_ID)
}
