package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
)

func main() {
    var isPy bool
    var isPyPy bool
    var np string
    re := regexp.MustCompile("(?Ui)C:.Python.+;")
    repypy := regexp.MustCompile("(?Ui)C:.pypy.+;")
    args := os.Args[1:]
    var path = re.ReplaceAllString(os.Getenv("path"), "")
    path = repypy.ReplaceAllString(path, "")
    if len(args) > 0 {
        isPyPy, _ = regexp.MatchString("(?i)PyPy.*", args[0])
        if isPyPy {
            np = fmt.Sprintf("C:\\%v;C:\\%v\\bin;", args[0], args[0])
        } else {
            np = fmt.Sprintf("C:\\%v;C:\\%v\\Scripts;", args[0], args[0])
        }
        //fmt.Println(np + path)
        os.Setenv("PATH", np+path)
        fmt.Println(os.Getenv("path"))
    } else {
        files, _ := ioutil.ReadDir("c:/")

        for _, f := range files {
            isPy, _ = regexp.MatchString("Python[0-9][0-9]", f.Name())
            isPyPy, _ = regexp.MatchString("(?i)PyPy.*", f.Name())
            if (isPy || isPyPy) && f.IsDir() {
                fmt.Println(f.Name())
            }
        }
    }

}
