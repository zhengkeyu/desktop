package link

import (
	_"unsafe"
)

//go:linkname linkRealize Desktop/记录/linkname/linktwo.Link
func linkRealize(){
	println("go linkname test")
}
