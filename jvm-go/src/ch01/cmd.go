package main

import (
	"flag"
	"fmt"
	"os"
)

/**
	命令结构体
 */
type Cmd struct {

	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

/**
	解析Cmd
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print the message and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd;
}



/**
	打印输出
 */
func printUsage()  {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

/**
	启动JVM
 */
func startJVM(cmd *Cmd)  {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}