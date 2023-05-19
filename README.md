# Brazil-Code

This is validator, generator and formats the brazil code as CPF and CNPJ

> Formats to use
- CPF
- CNPJ
- CNH
 - VoterRegistration(Título de Eleitor)

## Development
Import package with command in shell
```shell
$ go get github.com/potatowski/brazilcode
```
In code just use the function with import
> Example:
```code
package main

import (
	"fmt"

	"github.com/potatowski/brazilcode"
)

func main() {
	doc, err := brazilcode.CNPJGenerate()
	if err != nil {
		panic(err)
	}

	docFormatted, err := brazilcode.CNPJFormat(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(doc, docFormatted)
}
```
## License

The MIT License © 2023 João Vitor Lima da Rocha
