package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	//rol de servidor
	//escucha
	ln, error := net.Listen("tcp", "localhost:8001") //IP:PUERTO
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
	r := bufio.NewReader(conn)

	msg, _ := r.ReadString('\n') //lee haste el salto de linea finaliza el mensaje para guardar el dato que llega
	var1 := msg
	fmt.Println(var1) //imprime mensaje del cliente

	//responde al cliente
	msgg := "servidor1 trabajo" + msg

	fmt.Println(msgg)
	fmt.Fprintln(conn, msgg)

}
