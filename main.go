package main

import (
	"database/sql"
	"net/http"

	//"log"
	"fmt"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func ConnectionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Pass := ""
	Nombre := "sistema de empleados"

	conexion, err := sql.Open(Driver, Usuario+":"+Pass+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/add", Add)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}
func Start(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := ConnectionDB()
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES('javier', 'correo@gmail.com')")

	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec()
	templates.ExecuteTemplate(w, "home", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "add", nil)
}
