package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"	
	"github.com/go-chi/chi"
	"./Servidor"
		
)

var info = &servidor.InfoSerCompleta {
	{
		ID: "1",
		Addres: "server1",
		Asl_grade: "B",
		Country: "US",
		Owner:"Amazon.com, Inc.",
	},
	/*{
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
	},*/
	Servers_changed: true,
	Ssl_grade: "B",
	Previous_ssl_grade: "A+",
	Logo: "https://server.com/icon.png",
	Title: "Title of the page",
	Is_down: false,
}

func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Bienvenido al mundo JuliLau")
}

func main(){
	fmt.Println("Servidor iniciado en el puerto: 3000")
	//Crea la ruta
	r := chi.NewRouter()
	
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	//El controlador aplica la especificación CORS en la solicitud
	//y agrega encabezados CORS relevantes según sea necesario.
	r.Use(cors.Handler)
	r.HandleFunc("/",indexRoute)
	fmt.Println("Servidor iniciado en el puerto: 3000")
	http.ListenAndServe(":3000", r)

	/*
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
	*/

	/*r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Tenemos respuesta en la ruta, sea bienvenido"))
	})*/

	//Necesito agregar la ruta con r.Get para obtener la información, pero debe ser un
	//con un método de la lógica ya 

	//Main terminado, baby 

}

//middleware que agregará encabezados de respuesta CORS o imprimirá información de registro para las solicitudes emitidas.