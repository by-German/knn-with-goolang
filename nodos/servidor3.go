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

	//rol de servidor
	//escucha
	ln, error := net.Listen("tcp", "localhost:8003") //IP:PUERTO
	if error != nil {
		log.Println("Falla al resolver la direccion", error.Error())
		os.Exit(1)
	}
	defer ln.Close() //garantiza que la conexion se cierre de manera correcta
	conn, error := ln.Accept()
	if error != nil {
		log.Println("Fallo al aceptar la conexion con el cliente", error.Error())
	}
	defer conn.Close()
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
	fmt.Fprintln(conn, respuesta)

}