package printer

import (
	"github.com/alexbrainman/printer"
)

func DefaultPrinter() (string, error) {
	return printer.Default()
}
