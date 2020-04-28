package bd

import (
	"database/sql"
	"fmt"
	"log"
)

/*
En la consola SQL es necesario ejecutar el comando:
cockroach sql --certs-dir=certs
(Modo seguro)
*/

func cockroachdb() {
	//Conectar con la base de datos Truora
	//Crear DB Truora
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/truora?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Error en la conexión", err)

	//
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXIST `
		`INSERT INTO tbl_employee (full_name, department, designation, created_at, update_at) 
		VALUES ('Julian Salgado', 'IT', 'Developer', NOW(), NOW());`); err != nil {
		log.Fatal(err)

	}

}
