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

	reader := bufio.NewReader(openfile)
	writer := bufio.NewWriter(outfile)
	// buf := make([]byte, 1024)
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
				// n, err := reader.Read(buf)
				line, isPrefix, err := reader.ReadLine()
				if err != nil {
                    if err != io.EOF {
                        fmt.Println(err)
                        fmt.Println("fuking read error!")
                        os.Exit(1)
                    } else {
                        break
                    }
				}
                if isPrefix {
                    fmt.Println("a too long line")
                }
				// if n == 0 {
				// 	fmt.Println("break")
				// 	break
				// }
				// _, err = writer.Write(buf[:n])
				_, err = writer.Write(line)
				_, err = writer.WriteString("\n")
				if err != nil {
                    fmt.Println(err)
				}
			}
		} else {
			writer.WriteString(line)
		}
	}
	//data := rw.Read(path)
	//wordLine := strings.Split(string(data), "\n")
	/*
	   var index int = 0
	   for ; index < len(wordLine); index++ {
	       if wordLine[index] >= line {
	           break
	       }
	   }
	   result := make([]string, len(wordLine)+1)
	   copy(result, wordLine[:index])
	   copy(result[index+1:], wordLine[index:])
	   result[index] = line

	   var wordLines string = ""
	   for _, value := range result {
	       wordLines += value + "\n"
	   }
	*/

	//rw.Write(wordLines, "./test_result.txt")
}
