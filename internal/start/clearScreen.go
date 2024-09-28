package start

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
