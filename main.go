package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/go-chi/cors"	
	"github.com/go-chi/chi"
	"./Servidor"
)

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
	r.Get("/getServers/{dominio}", servidor.GetServers)
	fmt.Println("Servidor iniciado en el puerto: 3000")
	log.Fatal(http.ListenAndServe(":3000", r)) 

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