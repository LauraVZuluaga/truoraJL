package client

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func ObtenerInfoServidor(dominio string) {
    response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host="+dominio)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

}