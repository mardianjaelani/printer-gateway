package printer

type Manager struct {
	driver Printer
}

func NewManager() *Manager {
	return &Manager{
		driver: NewWindowsPrinter(),
	}
}

func (m *Manager) Print(printer string, data []byte, copies int) error {
	return m.driver.Print(printer, data, copies)
}
