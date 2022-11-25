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
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func Start(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := ConnectionDB()

	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	arrayEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arrayEmpleado = append(arrayEmpleado, empleado)
	}
	//fmt.Println(arrayEmpleado)

	templates.ExecuteTemplate(w, "home", arrayEmpleado)
}

func Add(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "add", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		conexionEstablecida := ConnectionDB()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(name,email) VALUES('?', '?')")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(name, email)

		http.Redirect(w, r, "/", 301)
	}
}
