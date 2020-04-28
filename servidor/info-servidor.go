package servidor

//InfoServidor: Estructura para la información del servidor
//Id para llave primaria en la base de datos

type InfoServidor struct {
	ID      string
	SslGrade string 
	Address string
	Country string
	Owner   string
}

//InfoSerCompleta: Estructura para la información completa del servidor
//asociada a la información parcial de la información anterior
//(Se anida)
type InfoSerCompleta struct {
	InfoServidores   []InfoServidor
	ServersChanged   bool
	SslGrade         string
	PreviousSslGrade string
	Logo             string
	Title            string
	IsDown           bool
}
