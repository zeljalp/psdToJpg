package main

import (
    "fmt"
    "os/exec"
    "runtime"
    "io/fs"
    "strings"
    "path/filepath"
)

func execute(cmd string, args []string) {

    // here we perform the pwd command.
    // we can store the output of this in our out variable
    // and catch any errors in err
    out, err := exec.Command(cmd, args...).Output()

    // if there is an error with our execution
    // handle it here
    if err != nil {
        fmt.Printf("%s\n", err)
    }
    // as the out variable defined above is of type []byte we need to convert
    // this to a string or else we will see garbage printed out in our console
    // this is how we convert it to a string
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)
}

func find(root, ext string) []string {
   var a []string
   filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
      if e != nil { return e }
      if filepath.Ext(d.Name()) == ext {
         a = append(a, s)
      }
      return nil
   })
   return a
}

func main() {
    //args := []string{"rm363-a002-rm390-mockup-10.psd[0]","-density", "150x150", "-resize", "25%x25%", "slika.jpg"}
    app := "/usr/bin/convert";
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
       // execute(app, args)
    }
   println("Now try to convert all .psd files ...")

   var psdName string;
   var jpgName string;
   for _, s := range find("/home/v4ing", ".psd") {
      println("Psd file for conversion is ")
      println(s)
      psdName = s;
      jpgName = strings.ReplaceAll(s, ".psd", ".jpg")
      args := []string{psdName+"[0]","-density", "150x150", "-resize", "25%x25%", jpgName}
      execute(app, args);
   }
}


