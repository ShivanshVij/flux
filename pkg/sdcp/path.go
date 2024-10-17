package sdcp

import "fmt"

type Path string

func USBPath(p string) Path {
	return Path(fmt.Sprintf("/usb/%s", p))
}

func LocalPath(p string) Path {
	return Path(fmt.Sprintf("/local/%s", p))
}
