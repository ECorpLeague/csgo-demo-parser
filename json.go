package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv"
)

type j struct {
    Cl []string `json:"cl"`
    Gr []string `json:"gr"`
    Cr []string `json:"cr"`
}

func main() {
    // create an instance of j as a slice
    var data []j
    // using a for loop for create dummy data fast
    for i := 0; i < 5; i++ {
        v := strconv.Itoa(i)
        data = append(data, j{
            Cl: []string{"foo " + v},
            Gr: []string{"bar " + v},
            Cr: []string{"foobar " + v},
        })
    }

    // printing out json neatly to demonstrate
    b, _ := json.MarshalIndent(data, "", " ")
    fmt.Println(string(b))

    // writing json to file

    _ = ioutil.WriteFile("file.json", b, 0644)

    // to append to a file
    // create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
    // f, err := os.OpenFile("/var/log/debug-web.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // write to file, f.Write()
    // f.Write(b)

    // if you are doing alot of I/O work you may not want to write out to file often instead load up a bytes.Buffer and write to file when you are done... assuming you don't run out of memory when loading to bytes.Buffer


}