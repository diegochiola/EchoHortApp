package main

import (
	"log"
	"net/http"
) //librerias

type PreUrl struct {
	Url    string       `json:"datos"` //no debe haber espacio
	Client *http.Client //guarda la informacion de la peticion del cliente
}

// A modo de pruebas hacemos funcion main aqui
func main() {
	GetPreUrl()
}
func GetPreUrl() (string, error) {
	url := "https://opendata.aemet.es/opendata/api/prediccion/especifica/municipio/horaria/08001/?api_key=eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJkaWVnb2NoaW9sYUBnbWFpbC5jb20iLCJqdGkiOiI3MGJkYmQ5OC0yNWRkLTQyZWQtOGIwNC05MGJlYWNlNDFjM2IiLCJpc3MiOiJBRU1FVCIsImlhdCI6MTcyNzk3NTYyNiwidXNlcklkIjoiNzBiZGJkOTgtMjVkZC00MmVkLThiMDQtOTBiZWFjZTQxYzNiIiwicm9sZSI6IiJ9.Ky3n7e_0X-Sb-4rVEXBzXqKRHCmDmyV4QQ84G3V-P3w" //url de la primera peticion
	peticio, _ := http.NewRequest("GET", url, nil)
	peticio.Header.Add("cache-control", "no-cache")
	answer, error := http.DefaultClient.Do(peticio) //ejecutar peticion

	if error != nil {
		log.Println("Error al conectar con la API", error)
		return "", error
	} //get peticion, envio url y body vacio (nil)
	//pendiente gestionar la respuesta
}
