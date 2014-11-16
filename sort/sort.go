package sort

import (
	//rw "github.com/logindave/newdict/src/rw"
	//"strings"
	"bufio"
	"fmt"
	"io"
	"os"
)

func Insert(insertStr, path string) {
	openfile, err := os.Open(path)
	if nil != err {
		fmt.Println("can't fuck open it")
	}
	outfile, err := os.Create("new.txt")
	if err != nil {
		fmt.Println("can;t fucking create it")
	}
	defer func() {
		outfile.Close()
		openfile.Close()
	}()
	defer openfile.Close()

	reader := bufio.NewReader(openfile)
	writer := bufio.NewWriter(outfile)
	buf := make([]byte, 1024)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
				fmt.Println(err)
			}
			break
		}
		if line > insertStr {
			writer.WriteString(insertStr + "findmehere\n")
			writer.WriteString(line)
			// write remain all to the file
			for {
				n, err := reader.Read(buf)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
						fmt.Println("fuking read error!")
						os.Exit(1)
					} else {
						writer.Flush()
						break
					}
				}
				_, err = writer.Write(buf[:n])
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			writer.WriteString(line)
		}
	}
}
