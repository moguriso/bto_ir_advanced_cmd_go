package ir

import (
	"log"

	"github.com/kylelemons/gousb/usb"
)

var (
	gVendor  usb.ID = 0x0
	gProduct usb.ID = 0x0
)

func setVendor(vid usb.ID) {
	gVendor = vid
}

func setProduct(pid usb.ID) {
	gProduct = pid
}

func getVendor() usb.ID {
	return gVendor
}

func getProduct() usb.ID {
	return gProduct
}

func listUsbDevice(desc *usb.Descriptor) bool {
	if desc.Vendor == getVendor() && desc.Product == getProduct() {
		return true
	} else {
		return false
	}
}

func OpenUsbDevice(vid, pid usb.ID) {
	ctx := usb.NewContext()
	defer ctx.Close()

	setVendor(vid)
	setProduct(pid)

	devs, err := ctx.ListDevices(listUsbDevice)
	if err != nil {
		log.Fatalf("list: %s", err)
	}

	// All Devices returned from ListDevices must be closed.
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	dev, err2 := ctx.OpenDeviceWithVidPid(int(getVendor()), int(getProduct()))
	if err2 != nil {
		log.Fatalf("list: %s", err)
	}
	_ = dev /* tentative */
}
