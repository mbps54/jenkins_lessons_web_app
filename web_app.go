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
                </body>
            </html>
        `)
    })

    err := http.ListenAndServe("0.0.0.0:8081", nil)
    if err != nil {
        panic(err)
    }
}
