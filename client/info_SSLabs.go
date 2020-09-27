package client

import (
    "fmt"
    "net/http"
    "os"
    "io"
    "encoding/json"
    "crypto/tls"
    "../models"
)

func ObtenerInfoServidor(fecha string) ([]models.Buyer,error) {
    // Create New http Transport
    transCfg := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
    }
    // Create Http Client
    client := &http.Client{Transport: transCfg}
    response, err := client.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com​/buyers?date="+fecha)

    if err != nil {
        fmt.Print("ERROR: " + err.Error())
        os.Exit(1)
    }
    defer response.Body.Close()
    /*responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    fmt.Println(responseObject)*/

    return decodeBuyers(response.Body)
    
}

func decodeBuyers(r io.Reader) ([]models.Buyer, error) {
	buyers := []models.Buyer{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&buyers)
	return buyers, err
}
