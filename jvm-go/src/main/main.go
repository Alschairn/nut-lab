package main

import (
	"classpath"
	"fmt"
	"os"
	"strings"

)

func main() {
	cmd := classpath.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

/**
	打印输出
 */
func printUsage()  {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}


func startJVM(cmd *classpath.Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath:%v class:%v, args:%v \n",
		cp, cmd.Class, cmd.Args)

	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.Class)
		return
	}
	fmt.Printf("class data:%v \n", classData)

}
