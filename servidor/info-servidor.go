package servidor

//InfoServidor: Estructura para la información del servidor
//Id para llave primaria en la base de datos

type InfoServidor struct {
	ID      string `json:ID`
	Address string	`json:address`
	Ssl_grade string `json:ssl_grade`
	country string	`json:country`
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
