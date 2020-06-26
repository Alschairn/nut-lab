package classpath

import (
	"flag"
	"fmt"
	"os"
)

/**
	命令结构体
 */
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	XjreOption  string
	Class       string
	Args        []string
}

/**
	解析Cmd
 */
func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print the message and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "print help message")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd;
}



/**
	打印输出
 */
func printUsage()  {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}

