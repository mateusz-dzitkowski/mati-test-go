package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmt := exec.Command("date")
	dateOut, _ := dateCmt.Output()
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err := exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	_ = grepCmd.Start()
	_, _ = grepIn.Write([]byte("hello grep\ngoodbye grep"))
	_ = grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	_ = grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -l -a -h")
	lsOut, _ := lsCmd.Output()
	fmt.Println("> ls -l -a -h")
	fmt.Println(string(lsOut))

}
