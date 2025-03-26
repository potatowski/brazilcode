# Brazil Code

The `brazilcode` package provides functionality to validate, generate, and format Brazilian identification codes, including CPF, CNPJ, CNH, and Título de Eleitor (Voter Registration).

## Supported Formats
- **CPF**: Cadastro de Pessoas Físicas (Individual Taxpayer Registry)
- **CNPJ**: Cadastro Nacional de Pessoas Jurídicas (National Register of Legal Entities)
- **CNH**: Carteira Nacional de Habilitação (National Driver's License)
- **Voter Registration (Título de Eleitor)**: Brazilian Electoral Registration Number

## Installation
To install the `brazilcode` package, run the following command in your shell:
```shell
$ go get github.com/potatowski/brazilcode
```

## Usage
Once the package is installed, you can use the functions provided to generate, validate, and format documents.

### Example Code:
```go
package main

import (
	"fmt"
	"github.com/potatowski/brazilcode"
)

func main() {
	// Generate a CNPJ document
	doc, err := brazilcode.CNPJ.Generate()
	if err != nil {
		panic(err)
	}

	// Format the CNPJ document
	docFormatted, err := brazilcode.CNPJ.Format(doc)
	if err != nil {
		panic(err)
	}

	// Print both unformatted and formatted documents
	fmt.Println(doc, docFormatted)
}
```

Alternatively, you can use the generic `Generate` and `Format` functions to generate and format any document type:
```go
doc, err := brazilcode.Generate("CNPJ")
if err != nil {
	panic(err)
}

docFormatted, err := brazilcode.Format("CNPJ", doc)
if err != nil {
	panic(err)
}

fmt.Println(doc, docFormatted)
```

## License
This project is licensed under the MIT License.

© 2023 João Vitor Lima da Rocha