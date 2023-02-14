package commander

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func ExecuteCommand(indicator string, in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Print(">> You have entered inside the commander. Now you can execute os commands here.\n")
	for {
		fmt.Print(indicator)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" || line == "" {
			os.Exit(0)
		}

		parts := strings.Split(line, " ")
		if len(parts) == 0 {
			os.Exit(0)
		}

		command, args := strings.TrimSpace(parts[0]), parts[1:]

		runner := exec.Command(command, args...)
		result, err := runner.CombinedOutput()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		out.Write(result)
	}
}
