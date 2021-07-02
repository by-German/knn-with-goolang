package main

import (
	"bufio"
	"fmt"
	"net"
)

type Alumno struct {
	Codigo string `json:cod`
	Nombre string `json:nom`
	Dni    int    `json:dni`
}

func main() {
	var menu int
	fmt.Println("Ingrese un numero entre 1 y 4")
	fmt.Scanln(&menu)

	if menu == 1 {
		//envio, Dial para comunicarse
		conn1, _ := net.Dial("tcp", "localhost:8001")
		defer conn1.Close()
		//enviamos el mensaje al servidor
		//fmt.Fprintln(conn1, "cliente 1 enviando datos al server 1!")
		fmt.Fprintln(conn1, "kfdskjldsfds")
		//recibe comunicacion
		r1 := bufio.NewReader(conn1)
		resp, _ := r1.ReadString('\n')
		fmt.Println("Respuesta al servidor 1", resp)
	} else if menu == 2 {
		//envio, Dial para comunicarse
		conn2, _ := net.Dial("tcp", "localhost:8002")
		defer conn2.Close()
		//enviamos el mensaje al servidor
		fmt.Fprintln(conn2, "cliente 1 enviando datos al server 2!")
		//recibe comunicacion
		r2 := bufio.NewReader(conn2)
		resp, _ := r2.ReadString('\n')
		fmt.Println("Respuesta al servidor 2", resp)
	} else if menu == 3 {
		//envio, Dial para comunicarse
		conn3, _ := net.Dial("tcp", "localhost:8003")
		defer conn3.Close()
		//enviamos el mensaje al servidor
		fmt.Fprintln(conn3, "cliente 1 enviando datos al server 3!")
		//recibe comunicacion
		r3 := bufio.NewReader(conn3)
		resp, _ := r3.ReadString('\n')
		fmt.Println("Respuesta al servidor 3", resp)
	} else {
		//envio, Dial para comunicarse
		conn4, _ := net.Dial("tcp", "localhost:8004")
		defer conn4.Close()
		//enviamos el mensaje al servidor
		fmt.Fprintln(conn4, "cliente 1 enviando datos al server 4!")
		//recibe comunicacion
		r4 := bufio.NewReader(conn4)
		resp, _ := r4.ReadString('\n')
		fmt.Println("Respuesta al servidor 4", resp)
	}

}
