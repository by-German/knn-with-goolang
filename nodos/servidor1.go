package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
)

func main() {
	//http.ListenAndServe(":8081", nil)

	//rol de servidor
	//escucha
	ln, error := net.Listen("tcp", ":8081") //IP:PUERTO
	if error != nil {
		log.Println("Falla al resolver la direccion", error.Error())
		os.Exit(1)
	}
	defer ln.Close() //garantiza que la conexion se cierre de manera correcta
	conn, error := ln.Accept()
	if error != nil {
		log.Println("Fallo al aceptar la conexion con el cliente", error.Error())
	}
	//recuperar lo que envia el cliente
	r1 := bufio.NewReader(conn)
	r2 := bufio.NewReader(conn)
	msg1, _ := r1.ReadString('\n') //lee haste el salto de linea finaliza el mensaje para guardar el dato que llega
	msg2, _ := r2.ReadString('\n')
	var1, _ := strconv.ParseFloat(msg1, 64)
	var2, _ := strconv.ParseFloat(msg2, 64)
	//fmt.Println(var1) //imprime mensaje del cliente

	//responde al cliente
	respuesta := math.Pow(var1-var2, 2)
	s := fmt.Sprintf("%f", respuesta)
	fmt.Fprintln(conn, s)

}
