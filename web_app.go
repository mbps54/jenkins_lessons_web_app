package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, `
            <!DOCTYPE html>
<html>
    <head>
        <style>
            #hello {
                font-size: 60px;
            }
        </style>
    </head>
    <body>
        <div id="hello">Hello from lesson 11!</div>
        <div class="example"></div>
    </body>
    <script>
        document.getElementsByClassName('example')[0].innerHTML = "<script>alert('This is a bug!');</script>";
    </script>
</html>
<html>
    <head>
        <style>
            #hello {
                font-size: 60px;
            }
        </style>
    </head>
    <body>
        <div id="hello">Hello from lesson 11!</div>
        <div class="example"></div>
    </body>
    <script>
        document.getElementsByClassName('example')[0].innerHTML = "<script>alert('This is a bug!');</script>";
    </script>
</html>

        `)
    })

    err := http.ListenAndServe("0.0.0.0:8081", nil)
    if err != nil {
        panic(err)
    }
}
