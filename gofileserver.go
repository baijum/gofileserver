package main

import (
    "log"
    "net/http"
    "flag"
    "strconv"
    "os"
)

var port = flag.Int("port", 9999, "port number")
var dir = flag.String("dir", "NaD", "directory to serve")

func main() {
    flag.Parse()
    wd, _ := os.Getwd()
    if *dir != "NaD" {
        wd = *dir
    }

    fileinfo, fileerr := os.Stat(wd)

    if fileerr != nil {
        if os.IsNotExist(fileerr) {
            log.Fatal("Error: ", fileerr)
        } else {
            log.Fatal(fileerr)
        }
    }

    if fileinfo.IsDir() {
        log.Println("Serving directory: " + wd)
    } else {
        log.Fatal("Not a directory: " + wd)
    }

    http.Handle("/",
        http.StripPrefix("/", http.FileServer(http.Dir(wd))))
    log.Println("Listening port: " + strconv.Itoa(*port))
    err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
