package main

import (	
	"fmt"	
    "log"
    "os"
	"os/exec"		
	"sync"
	"bufio"
    "io"
    // "sort"

    "github.com/urfave/cli/v2"
)

func main() {
    // app := &cli.App{
    //     Name:  "boom",
    //     Usage: "make an explosive entrance",
    //     Action: func(ctx *cli.Context) error {

    //         // Arguments 参数  
    //         fmt.Printf("step1 boom! I say %q \n", ctx.Args().Get(0)) 

    //         // Command("ls -l")
    //         // Command("cd a; ls") // cd成功了
    //         // fmt.Println("step2 22222222")
    //         // Command("ls")

    //         return nil
    //     },
    // }




    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "add",
                Aliases: []string{"a"},
                Usage:   "add a task to the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("added task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "complete",
                Aliases: []string{"c"},
                Usage:   "complete a task on the list",
                Action: func(cCtx *cli.Context) error {
                    fmt.Println("completed task: ", cCtx.Args().First())
                    return nil
                },
            },
            {
                Name:    "template",
                Aliases: []string{"t"},
                Usage:   "options for task templates",
                Subcommands: []*cli.Command{
                    {
                        Name:  "add",
                        Usage: "add a new template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("new task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                    {
                        Name:  "remove",
                        Usage: "remove an existing template",
                        Action: func(cCtx *cli.Context) error {
                            fmt.Println("removed task template: ", cCtx.Args().First())
                            return nil
                        },
                    },
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

// go调用shell
func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)  // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}
