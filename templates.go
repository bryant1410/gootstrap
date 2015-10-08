package main

//README.md template
var readmeTempl = `#%s

This is the awesome description for %s.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).`

//doc.go template
var docTempl = `// Add some documentation to your package.
package main`

//Main .go file template
var mainTempl = `package main

import (
	"fmt"
)

func main() {
	fmt.Prinln("Hello from Gootstrap!")
}
`

//Main _test.go file template
var mainTestTempl = `package main

import (
	"testing"
)

func Test(t *testing.T) {
	
}
`

var travisTempl = `language: go
sudo: false

go:
  - 1.3
  - 1.4
  - 1.5
  - tip

script:
  - go test -v ./...
`

var mitLicenseTempl = `The MIT License (MIT)

Copyright (c) %d %s

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.`
