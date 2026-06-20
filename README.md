# Brazil Code

The `brazilcode` package provides functionality to validate, generate, and format Brazilian identification codes, including CPF, CNPJ, CNH, RENAVAM, and Título de Eleitor (Voter Registration).

## Supported Formats
- **CPF**: Cadastro de Pessoas Físicas (Individual Taxpayer Registry)
- **CNPJ**: Cadastro Nacional de Pessoas Jurídicas (National Register of Legal Entities)
- **CNH**: Carteira Nacional de Habilitação (National Driver's License)
- **Voter Registration (Título de Eleitor)**: Brazilian Electoral Registration Number
- **RENAVAM**: Registro Nacional de Veículos Automotores (National Registry of Motor Vehicles)

## Installation

```shell
go get github.com/potatowski/brazilcode/v3
```

## Usage

### Using the Facade (root package)

The root package provides both direct document variables and string-based dispatch:

```go
package main

import (
	"fmt"
	"github.com/potatowski/brazilcode/v3"
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

	fmt.Println(doc, docFormatted)

	// Validate a CPF
	if err := brazilcode.CPF.IsValid("123.456.789-09"); err != nil {
		fmt.Println("invalid CPF:", err)
	}

	// Generate a Voter Registration for a specific state (UF)
	voter, err := brazilcode.VoterRegistration.Generate(brazilcode.WithUF("MG"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Voter Registration:", voter)
}
```

You can also use the string-based generic functions:

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

### Using Sub-Packages Directly

Each document type is also available as a standalone package:

```go
package main

import (
	"fmt"
	"github.com/potatowski/brazilcode/v3/cpf"
	"github.com/potatowski/brazilcode/v3/cnpj"
	"github.com/potatowski/brazilcode/v3/voter"
	"github.com/potatowski/brazilcode/v3/internal/digit"
)

func main() {
	c := cpf.CPF{}

	// Validate
	if err := c.IsValid("123.456.789-09"); err != nil {
		fmt.Println("invalid:", err)
	}

	// Generate
	doc, _ := cnpj.CNPJ{}.Generate()
	fmt.Println("CNPJ:", doc)

	// Generate Voter Registration with UF option
	v, _ := voter.VoterRegistration{}.Generate(digit.WithUF("SP"))
	fmt.Println("Voter:", v)
}
```

Available sub-packages:

| Package | Import Path | Document |
|---------|-------------|----------|
| `cpf` | `github.com/potatowski/brazilcode/v3/cpf` | CPF |
| `cnpj` | `github.com/potatowski/brazilcode/v3/cnpj` | CNPJ |
| `cnh` | `github.com/potatowski/brazilcode/v3/cnh` | CNH |
| `renavam` | `github.com/potatowski/brazilcode/v3/renavam` | RENAVAM |
| `voter` | `github.com/potatowski/brazilcode/v3/voter` | Título de Eleitor |

## Migrating from v2

### Breaking Changes

- **Module path**: `github.com/potatowski/brazilcode/v2` → `github.com/potatowski/brazilcode/v3`
- **Package structure**: The `src/` directory has been removed. Documents are now in root-level sub-packages (`cpf/`, `cnpj/`, etc.).
- **`Generate` signature**: Changed from `Generate(params map[string]string)` to `Generate(opts ...Option)` using the functional options pattern.
- **Voter Registration**: Package renamed from `voterRegistration` to `voter`.

### Migration Example

```diff
- import "github.com/potatowski/brazilcode/v2"
+ import "github.com/potatowski/brazilcode/v3"

- doc, err := brazilcode.VoterRegistration.Generate(map[string]string{"uf": "MG"})
+ doc, err := brazilcode.VoterRegistration.Generate(brazilcode.WithUF("MG"))
```

## License
This project is licensed under the MIT License.

© 2023 João Vitor Lima da Rocha