package packageone

import "fmt"

var PackageVar = "Package Variable in packageone package"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
