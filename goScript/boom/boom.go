package main

import (	
	"fmt"	
    "log"
    "os"
	"os/exec"		
	"sync"
	"bufio"
    "io"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:  "boom",
        Usage: "make an explosive entrance",
        Action: func(*cli.Context) error {
            
            fmt.Println("boom! I say!")                        
            Command("ls")
            Command("cd a; ls -l") // cd成功了

            fmt.Println("222222222")
            Command("ls")

            return nil
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
