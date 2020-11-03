package server

import(
	"net/http"
	"../client"
	"github.com/go-chi/chi"
	"encoding/json"
	"../bd"
	"fmt"
)

//InfoServidor: Estructura para la información del servidor
//Id para llave primaria en la base de datos

type InfoServidor struct {
	ID      string `json:"ID"`
	Address string	`json:"address"`
	Ssl_grade string `json:"ssl_grade"`
	Country string	`json:"country"`
	Owner   string	`json:"owner"`
}

//InfoSerCompleta: Estructura para la información completa del servidor
//asociada a la información parcial de la información anterior
//(Se anida)
type InfoSerCompleta struct {
	Servers   []InfoServidor `json:"servers"`
	Servers_changed   bool	`json:"servers_changed"`
	Ssl_grade         string	`json:"ssl_grade"`
	Previous_ssl_grade string	`json:"PreviousSslGrade"`
	Logo             string	`json:"logo"`
	Title            string	`json:"title"`
	Is_down           bool	`json:"is_down"`
}

type Servers []InfoServidor
var servers = Servers{
	{
		ID: "1",
		Address: "server1",
		Ssl_grade: "B",
		Country: "US",
		Owner:"Amazon.com, Inc.",
	},
	{
		ID:"2",
		Address: "server2",
		Ssl_grade: "A+",
		Country: "US",
		Owner: "Amazon.com, Inc.",
	},
	{
		ID:"3",
		Address: "server3",
		Ssl_grade: "A",
		Country: "US",
		Owner: "Amazon.com, Inc.",
	},
}
var info = InfoSerCompleta {
	Servers: servers,
	Servers_changed: true,
	Ssl_grade: "B",
	Previous_ssl_grade: "A+",
	Logo: "https://server.com/icon.png",
	Title: "Title of the page",
	Is_down: false,
}

//Permite leer el dominio que se escribe como parametro en la ruta URL
// de manera que con esté se pueda consultar la información de los servidores
//que esta en las apis utilizadas
func GetServers(w http.ResponseWriter, r *http.Request){
	date := chi.URLParam(r,"fecha")
	
	buyers,err := client.GetBuyers(date)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Something went wrong")
	}

	products := client.GetProducts(date)
	transactions := client.GetTransactions(date)

	bd.SaveBuyers(buyers);
	bd.SaveProducts(products);
	bd.SaveTransactions(transactions);

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,"OK")
}


func GetBuyers(w http.ResponseWriter, r *http.Request){

	buyers, err := bd.GetBuyers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Something went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buyers)
}