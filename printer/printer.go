package printer

type Printer interface {
	Print(printerName string, data []byte, copies int) error
}
