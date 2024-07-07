package playground

import (
	"fmt"
	"go/build"
	"log"
)

func TestImportPaths() {

	buildPkg, err := build.Import("github.com/gpbPiazza/notebook/algo_data/queues", "", build.IgnoreVendor)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(buildPkg.Imports)
}
