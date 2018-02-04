package main

import "syscall"
import "os"
import "os/exec"
import "fmt"


func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("> which ls\n", binary)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// Note that Go does not offer a classic Unix fork function.
// Usually this isn’t an issue though, since starting goroutines,
// spawning processes, and exec’ing processes covers most use cases for fork.
