package main

import "fmt"
import "io/ioutil"
import "os/exec"


func main() {

	dateCmd := exec.Command("date")

	// running a command, wait for it to finish,
	// and collecting its output.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	// get input/output pipes
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	// start command in background
	grepCmd.Start()

	// grab stdin/stdout
	grepIn.Write([]byte("hello grep\ngood bye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)

	// wait for done
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))


	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

