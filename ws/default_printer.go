package ws

import (
	"os/exec"
)

func GetDefaultPrinter() (string, error) {

	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-CimInstance Win32_Printer -Filter 'Default=True').Name",
	)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
