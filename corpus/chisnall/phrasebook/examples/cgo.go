package arc4random 
// #include <stdlib.h>
import "C"

func arc4random() uint32 {
	return uint32(C.arc4random())
}
