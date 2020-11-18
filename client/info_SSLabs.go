package client

import (
  "fmt"
  "net/http"
  "os"
  "io"
  "encoding/json"
  "encoding/csv"
  "crypto/tls"
  "../models"
  "strconv"
  "io/ioutil"
  "bytes"
  "strings"
)

const url string = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com​";

func requestServer(endpoint string, param string) (*http.Response) {
  // Create New http Transport
  transCfg := &http.Transport {
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
  }
  // Create Http Client
  client := &http.Client{Transport: transCfg}
  response, err := client.Get(url + endpoint + param)
  if err != nil {
      fmt.Print("ERROR: " + err.Error())
      os.Exit(1)
  }
  return response
}


func GetBuyers(date string) ([]models.Buyer,error) {
  response := requestServer("/buyers?date=", date)
  defer response.Body.Close()
  return decodeBuyers(response.Body)
  /*responseData, err := ioutil.ReadAll(response.Body)
  if err != nil {
      log.Fatal(err)
  }
  var responseObject Response
  json.Unmarshal(responseData, &responseObject)
  fmt.Println(responseObject)*/
}


func GetProducts(date string) ([]models.Product) {
  response := requestServer("/products?date=", date)
  defer response.Body.Close()
  return decodeProducts(response.Body)
}

func GetTransactions(date string) ([]models.Transaction) {
  response := requestServer("/transactions?date=", date)
  defer response.Body.Close()
  return decodeTransactions(response.Body)
}

func decodeTransactions(r io.Reader) ([]models.Transaction) {
  var transaction models.Transaction
  var transactions []models.Transaction
  responseData, err := ioutil.ReadAll(r)
  if err != nil {
      panic(err)
  }
  data := bytes.Split(responseData, []byte(string(0)))
  
  for i := 0; i*6+1 < len(data); i++ {

    transaction.ID = string(data[i*6])
    transaction.BuyerID = string(data[i*6+1])
    transaction.Ip = string(data[i*6+2])
    transaction.Device = string(data[i*6+3])
    productsIDs := strings.ReplaceAll(strings.ReplaceAll(string(data[i*6+4]),"(",","),")",",")
    transaction.ProductIDs = strings.Split(productsIDs, ",") 
    
    transactions = append(transactions, transaction)
  }

  return transactions
}

func decodeBuyers(r io.Reader) ([]models.Buyer, error) {
  buyers := []models.Buyer{}
  dec := json.NewDecoder(r)
  err := dec.Decode(&buyers)
  return buyers, err
}

func decodeProducts(r io.Reader) ([]models.Product) {
  var product models.Product
  var products []models.Product
  dec := csv.NewReader(r)
  dec.Comma = (0x0027)
  dec.LazyQuotes = true
  records, err := dec.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  for _, rec := range records {
    product.ID = rec[0]
    product.Name = rec[1]
    product.Price, _ = strconv.Atoi(rec[2])

    products = append(products, product)
  }

  return products
}