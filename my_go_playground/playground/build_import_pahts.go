package playground

import (
	"github.com/gpbPiazza/notebook/my_go_playground/playground/archp"
)

func TestImportPaths() {

	// notAllowedDependOn := map[string]bool{
	// 	"github.com/gpbPiazza/notebook/algo_data/queues": true,
	// }

	// buildPkg, err := build.Import("github.com/gpbPiazza/notebook/algo_data/graphs", "", build.ImportComment)
	// if err != nil {
	// 	log.Fatalf("Failed to import package: %v", err)
	// }

	// for _, imp := range buildPkg.Imports {
	// 	// if notAllowedDependOn[imp] {
	// 	// 	log.Fatalf("It is not allowed to graphs package depend on %v package", imp)
	// 	// }
	// 	fmt.Printf("Imported package: %s\n", imp)
	// }

	archp.NewVerifier("github.com/gpbPiazza/notebook/algo_data/graphs").Verify()
}
