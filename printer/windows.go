package printer

type WindowsPrinter struct {
}

func NewWindowsPrinter() *WindowsPrinter {
	return &WindowsPrinter{}
}

func (w *WindowsPrinter) Print(
	printer string,
	data []byte,
	copies int,
) error {

	return writeRaw(
		printer,
		data,
		copies,
	)

}
