package brackets

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const Shell = "sh"

func (b *brackets)RunCommand(command string) *brackets {
	cmd := exec.Command(Shell, "-c", command)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		for _, l := range StringToLines(string(stdout)) {
			if l == "" {
				continue
			}
			b.PrintDangerLn(l)
		}
		b.PrintDangerLn(err.Error())
		return b
	}

	for _, l := range StringToLines(string(stdout)) {
		b.PrintInfoLn(l)
	}
	return b
}

func (b *brackets)RunCommandInteractive(command string) *brackets {
	cmd := exec.Command(Shell, "-c", command)
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			b.PrintInfoLn(scanner.Text())
			stdin.Write([]byte("some sample text\n"))
		}
	}(reader)
	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()

	return b
}


func StringToLines(s string)[]string {
	return strings.Split(s, "\n")
}
