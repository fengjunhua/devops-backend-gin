package cmd

import (
	"fmt"
	"os/exec"
)

func LinuxExec() string{

	cmd := exec.Command("/bin/ls","-l","/var/")
	output,err := cmd.Output()
	if err != nil{
		panic(err)
	}
	fmt.Println(string(output))
	return string(output)
}