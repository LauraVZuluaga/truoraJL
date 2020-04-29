package servidor

import(
	"net/http"
	"../client"
	"github.com/go-chi/chi"
)

//InfoServidor: Estructura para la información del servidor
//Id para llave primaria en la base de datos

type InfoServidor struct {
	ID      string `json:ID`
	Address string	`json:address`
	Ssl_grade string `json:ssl_grade`
	Country string	`json:country`
	Owner   string	`json:owner`
}

//InfoSerCompleta: Estructura para la información completa del servidor
//asociada a la información parcial de la información anterior
//(Se anida)
type InfoSerCompleta struct {
	Servers   []InfoServidor `json:servers`
	Servers_changed   bool	`json:servers_changed`
	Ssl_grade         string	`json:ssl_grade`
	Previous_ssl_grade string	`json:PreviousSslGrade`
	Logo             string	`json:logo`
	Title            string	`json:title`
	Is_down           bool	`json:is_down`
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

func GetServers(w http.ResponseWriter, r *http.Request){
	vars := chi.URLParam(r,"dominio")
	client.ObtenerInfoServidor(vars)
}