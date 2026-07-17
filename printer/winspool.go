//go:build windows

package printer

import (
	"fmt"

	winprinter "github.com/alexbrainman/printer"
)

func writeRaw(printerName string, data []byte, copies int) error {

	p, err := winprinter.Open(printerName)
	if err != nil {
		return fmt.Errorf("open printer: %w", err)
	}
	defer p.Close()

	for i := 0; i < copies; i++ {

		if err := p.StartDocument("Go Print Gateway", "RAW"); err != nil {
			return fmt.Errorf("start document: %w", err)
		}

		if err := p.StartPage(); err != nil {
			return fmt.Errorf("start page: %w", err)
		}

		if _, err := p.Write(data); err != nil {
			return fmt.Errorf("write printer: %w", err)
		}

		if err := p.EndPage(); err != nil {
			return fmt.Errorf("end page: %w", err)
		}

		if err := p.EndDocument(); err != nil {
			return fmt.Errorf("end document: %w", err)
		}
	}

	return nil
}
