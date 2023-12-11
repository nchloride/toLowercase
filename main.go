package main 

import (
        "fmt"
        "flag"
        "os"
        "strings"
        "os/exec"
        "unicode"

)

var filePath string

func main(){
        flag.StringVar(&filePath,"name","","turn all directory to lower case")
        flag.Parse()

        if filePath != "" {
                lowerCaseFileName()
                return
        }


}

func lowerCaseFileName(path ...string){
        tmpPath := ""
        if len(path) != 0 {
               tmpPath = strings.ToLower(path[0])
        }
        files,err := os.ReadDir(filePath+"/"+tmpPath)
        if err != nil {
                fmt.Println(err)
        }

        for _,file :=  range files {
              hC :=  hasCapitalLetters(file.Name())
              if !hC{
                      continue
              }
              currentFile := filePath+tmpPath+"/"+file.Name()
              renameAndRemove(currentFile,file.IsDir(),tmpPath+"/"+file.Name())
              fmt.Println("current file name:" + currentFile)
        }

}

func renameAndRemove(currentFile string,isDir bool,fileToAppend ...string){
        fta := ""
        if len(fileToAppend) != 0 {
                fta = fileToAppend[0]
        }
        renameCommand := exec.Command("cp","-r", currentFile, strings.ToLower(currentFile))
        var out strings.Builder
        renameCommand.Stdout = &out

        if err := renameCommand.Start(); err !=nil {
                fmt.Println(err)
        }
	renameCommand.Wait()

        fmt.Println(out.String())
        removeCommand := exec.Command("rm","-rf",currentFile)
        if err := removeCommand.Start(); err !=nil {
                fmt.Println(err)
        }
        removeCommand.Wait()
        fmt.Println("FTA:" +fta)
	if isDir {
		lowerCaseFileName(fta)
	}



}

func hasCapitalLetters(name string) bool{
      for i:=0; i<len(name); i++ {
              if  unicode.IsUpper(rune(name[i])){
                      return unicode.IsUpper(rune(name[i]))
                      break
              }
      }
      return false
}

