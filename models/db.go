package models

import (
	"database/sql"
	"fmt"
	// Justificación del "_": Porque sí :v
	_ "github.com/go-sql-driver/mysql"
)

const bASE, uSUARIO, cONTRASEÑA, hOST string = "proyecto", "olakase00", "Jimmy", "localhost"

// DB -> conexion con la base
var DB *sql.DB

// Inicializar -> Crea y verifica la conexion
func Inicializar() error {
	var err error
	DB, err = sql.Open("mysql", cONTRASEÑA+":"+uSUARIO+"@/"+bASE)
	if err != nil {
		DB = nil
		return err
	}
	_, err = DB.Begin()
	if err != nil {
		DB = nil
		return err
	}
	err = DB.Ping()
	if err != nil {
		DB = nil
		return err
	}
	/////////////////
	rows, err := DB.Query("SELECT * FROM test;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var entero int
		var fecha string
		if err := rows.Scan(&entero, &fecha); err != nil {
			panic(err)
		}
		fmt.Printf("Datos de la base: %d, %s\n", entero, fecha)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return nil
}

//
// func EjecutarSentencia(sentencia string) (sql.Rows, error) {
// 	var err error
// 	var rows *sql.Rows
// 	rows, err = DB.Query(sentencia)
// 	if err != nil {
// 		return *rows, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var entero int
// 		var fecha string
// 		if err := rows.Scan(&entero, &fecha); err != nil {
// 			return *rows, err
// 		}
// 		fmt.Printf("Datos de la base: %n, %s\n", entero, fecha)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return *rows, err
// 	}
// 	return *rows, nil
// }
