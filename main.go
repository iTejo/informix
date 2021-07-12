package main

import (
	"database/sql"
	"fmt"

	_ "github.com/alexbrainman/odbc"
)

func main() {

	var nombre string

	fmt.Println(" Conexion Go al Informix ")

	db, err := sql.Open("odbc", "DSN=Infdrv1;locale=es_ES.819;") // nombre del conector, nombre del driver

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error en el ping : ", err)
	}

	//Query a ejecutar
	rows, err := db.Query("select nombre from socios")

	if err != nil {
		fmt.Println("No va la Query :( ", err)
	}

	//Parse
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&art_descri)
		if err != nil {
			fmt.Println("Error al parsear :( ", err)
		}

		fmt.Println(art_descri)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	//cerrar conexion
	defer db.Close()
}
