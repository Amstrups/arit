module arit

go 1.23.0

replace github.com/amstrups/nao => ../nao

require (
	github.com/amstrups/nao v0.0.0-00010101000000-000000000000
	golang.org/x/term v0.28.0
)

require golang.org/x/sys v0.29.0 // indirect

require github.com/davecgh/go-spew v1.1.1 // direct
