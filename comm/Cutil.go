package comm

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func StringSlice2C(strs []string) *string {
	cStrs := make([]*C.char, len(strs))
	for i := 0; i < len(strs); i++ {
		char := C.CString(strs[i])
		cStrs[i] = char
	}

	return (*string)(unsafe.Pointer(&cStrs[0]))
}
