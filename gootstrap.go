//Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var readmeTempl = `#%s

This is the awesome description for %s.`

var docTempl = `// Add some documentation to your package.
package %s`

var mainTempl = `package main

import (
	"fmt"
)

func main() {
	fmt.Prinln("Hello from Gootstrap!")
}
`

func run(command, pack_name string, out io.Writer) {
	switch command {
	case "new":
		createPackage(pack_name, out)
		fmt.Fprintf(out, "===> Package created! cd %s to access.", pack_name)
	}
}

func main() {

	pack_name := os.Args[2]
	log.Printf("===> Creating package %s", pack_name)
	run(os.Args[1], os.Args[2], os.Stdout)

}

func createPackage(pack_name string, out io.Writer) {
	sep := string(filepath.Separator)

	// Creates the project's folder

	if _, err := os.Stat(pack_name); os.IsNotExist(err) {
		os.Mkdir(pack_name, 0777)
		fmt.Fprintf(out, "===> Creating directory")
	}

	//Creates README.md
	readme := fmt.Sprintf("%s%sREADME.md", pack_name, sep)
	readme_file, err := os.Create(readme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer readme_file.Close()
	fReadme := fmt.Sprintf(readmeTempl, pack_name, pack_name)
	readme_file.WriteString(fReadme)
	fmt.Fprintf(out, "===> Creating README.md")

	// Creates main .go file
	mainpack := fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name)
	mainpack_file, err := os.Create(mainpack)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer mainpack_file.Close()
	mainpack_file.WriteString(mainTempl)
	fmt.Fprintf(out, "===> Creating main .go file.")

	// Creates main doc.go file
	doc := fmt.Sprintf("%s%sdoc.go", pack_name, sep)
	doc_file, err := os.Create(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer doc_file.Close()
	dReadme := fmt.Sprintf(docTempl, pack_name)
	doc_file.WriteString(dReadme)
	fmt.Fprintf(out, "===> Creating doc.go")
}
