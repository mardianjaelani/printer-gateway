package printer

func RawPrint(printerName string, data []byte, copies int) error {

	if printerName == "" {

		def, err := DefaultPrinter()
		if err != nil {
			return err
		}

		printerName = def
	}

	return writeRaw(printerName, data, copies)
}
