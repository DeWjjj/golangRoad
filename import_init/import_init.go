package import_init

import "fmt"

func init() {
	fmt.Println("PACKAGE_INIT()")
}
func Plus(x,y int) int {
	return x + y
}
