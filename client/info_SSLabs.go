package client

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

type Response struct{
    Servidores []Servidor `json:"endpoints"`
} 

type Servidor struct {
    Address string `json:"ipAddress"`
    Ssl_grade string `json:"grade"`
}

func ObtenerInfoServidor(dominio string) Response {
    response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host="+dominio)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

    return responseObject
}