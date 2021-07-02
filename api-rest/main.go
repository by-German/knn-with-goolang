package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Slice struct {
	sort.Float64Slice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(list []float64) *Slice {
	s := &Slice{
		Float64Slice: sort.Float64Slice(list),
		idx:          make([]int, len(list))}

	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

type Data struct {
	Departamento  string
	Parentesco    float64
	MiembroHogar  float64
	Edad          float64
	NivelEstudios float64
	Sexo          float64
	EstadoCivil   float64
	Discapacidad  int
}

func LoadData() []Data {
	data := make([]Data, 0)
	response, err := http.Get("https://raw.githubusercontent.com/by-German/knn-with-goolang/master/CAP_100_POBLACION%20TOTAL_Muestra.csv")

	if err != nil {
		fmt.Println("Error al leer el archivo")
	}
	r := csv.NewReader(response.Body)
	r.Comma = '|'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || record[0] == "NOMBREDD" {
			continue
		}
		temp := Data{}
		temp.Departamento = record[0]
		temp.Parentesco, _ = strconv.ParseFloat(record[7], 64)
		temp.MiembroHogar, _ = strconv.ParseFloat(record[8], 64)
		temp.Edad, _ = strconv.ParseFloat(record[10], 64)
		temp.Sexo, _ = strconv.ParseFloat(record[9], 64)
		temp.Discapacidad, _ = strconv.Atoi(record[17])
		if record[14] == "" {
			temp.NivelEstudios = 9
			data = append(data, temp)
			continue
		}
		if record[13] == "" {
			temp.EstadoCivil = 1
			data = append(data, temp)
			continue
		}
		temp.EstadoCivil, _ = strconv.ParseFloat(record[13], 64)
		temp.NivelEstudios, _ = strconv.ParseFloat(record[14], 64)
		data = append(data, temp)
	}
	return data
}

func EuclideanDistance(i, n int, x []Data, y Data, ch chan []float64) {
	count := 0
	distancia := make([]float64, n-i)

	for v := i; v < n; v++ {

		dato1 := x[v].NivelEstudios
		dato2 := y.NivelEstudios
		dato3 := x[v].Edad
		dato4 := y.Edad
		dato5 := x[v].Sexo
		dato6 := y.Sexo
		dato7 := x[v].EstadoCivil
		dato8 := y.EstadoCivil

		//DISTRIBUCION A NODOS
		//NODO1------------------------
		conn1, _ := net.Dial("tcp", "localhost:8001")
		defer conn1.Close()
		fmt.Fprintln(conn1, dato1)
		fmt.Fprintln(conn1, dato2)
		r1 := bufio.NewReader(conn1)
		resp1, _ := r1.ReadString('\n')
		//NODO2------------------------
		conn2, _ := net.Dial("tcp", "localhost:8002")
		defer conn2.Close()
		fmt.Fprintln(conn2, dato3)
		fmt.Fprintln(conn2, dato4)
		r2 := bufio.NewReader(conn2)
		resp2, _ := r2.ReadString('\n')
		//NODO3------------------------
		conn3, _ := net.Dial("tcp", "localhost:8003")
		defer conn3.Close()
		fmt.Fprintln(conn3, dato5)
		fmt.Fprintln(conn3, dato6)
		r3 := bufio.NewReader(conn3)
		resp3, _ := r3.ReadString('\n')
		//NODO4------------------------
		conn4, _ := net.Dial("tcp", "localhost:8004")
		defer conn4.Close()
		fmt.Fprintln(conn4, dato7)
		fmt.Fprintln(conn4, dato8)
		r4 := bufio.NewReader(conn4)
		resp4, _ := r4.ReadString('\n')
		/*
			nodo1 := math.Pow(x[v].NivelEstudios-y.NivelEstudios, 2)
			nodo2 := math.Pow(x[v].Edad-y.Edad, 2)
			nodo3 := math.Pow(x[v].Sexo-y.Sexo, 2)
			nodo4 := math.Pow(x[v].EstadoCivil-y.EstadoCivil, 2)
		*/

		nodo1, _ := strconv.ParseFloat(resp1, 64)
		nodo2, _ := strconv.ParseFloat(resp2, 64)
		nodo3, _ := strconv.ParseFloat(resp3, 64)
		nodo4, _ := strconv.ParseFloat(resp4, 64)

		distancia[count] += math.Sqrt(nodo1 + nodo2 + nodo3 + nodo4)
		count++
	}
	ch <- distancia
}

func traindata(x []Data, k int, nProcesos int) Data {
	var y Data
	rand.Seed(time.Now().UnixNano())

	y.Departamento = "TUMBES"
	y.Sexo = float64(rand.Intn(2 - 0))
	y.Edad = float64(rand.Intn(70 - 0))
	y.EstadoCivil = float64(rand.Intn(3 - 0))
	y.MiembroHogar = float64(rand.Intn(2 - 0))
	y.NivelEstudios = float64(rand.Intn(9 - 0))
	y.Parentesco = float64(rand.Intn(10 - 0))

	channels := make([]chan []float64, nProcesos)
	aumento := len(x) / nProcesos
	count := 0
	for i := 0; i < len(x); i += aumento {
		channels[count] = make(chan []float64)
		go EuclideanDistance(i, i+aumento, x, y, channels[count])
		count++
	}

	distancia := make([]float64, 0)
	for i := 0; i < nProcesos; i++ {
		distancia = append(distancia, <-channels[i]...)
	}
	s := NewSlice(distancia)
	sort.Sort(s)
	count = 0

	for i := 0; i < k; i++ {
		if x[s.idx[i]].Discapacidad == 1 {

			count++
		}
		// fmt.Println(x[s.idx[i]])
	}
	if count > k-count {
		y.Discapacidad = 1
	} else {
		y.Discapacidad = 0
	}
	return y
}

func findknn(x []Data, k int, y Data, nProcesos int) Data {

	// secction of algorithm
	channels := make([]chan []float64, nProcesos)
	aumento := len(x) / nProcesos
	count := 0
	for i := 0; i < len(x); i += aumento {
		channels[count] = make(chan []float64)
		go EuclideanDistance(i, i+aumento, x, y, channels[count])
		count++
	}

	distancia := make([]float64, 0)
	for i := 0; i < nProcesos; i++ {
		distancia = append(distancia, <-channels[i]...)
	}
	s := NewSlice(distancia)
	sort.Sort(s)
	count = 0
	y.Discapacidad = 0
	for i := 0; i < k; i++ {
		if x[s.idx[i]].Discapacidad == 1 {
			count++
		}
		fmt.Println(x[s.idx[i]])
	}
	if count > k-count {
		y.Discapacidad = 1
	}
	return y
}
func RemoveIndex(s []Data, index int) []Data {
	return append(s[:index], s[index+1:]...)
}
func Knn(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	// k, err := strconv.Atoi(req.FormValue("k"))
	// if err != nil {
	// 	k = 3
	// }
	nProceso := 4
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error al leer el body")
	}
	k := 2
	rand.Seed(time.Now().UnixNano())

	var myData Data
	json.Unmarshal(body, &myData)
	y := myData
	x := LoadData()

	for i := 0; i < k; i++ {
		dto := traindata(x, k, nProceso)
		x = append(x, dto)
		x = RemoveIndex(x, rand.Intn((len(x) - 1)))
		fmt.Println(i, dto)
	}

	y = findknn(x, k, y, nProceso)
	// end secction
	json, _ := json.Marshal(y)
	fmt.Println("//////////////////////")
	io.WriteString(res, string(json))
	fmt.Println("//////////////////////")
	fmt.Println("knn calculated ", y)

}

func GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	data := LoadData()

	json, _ := json.Marshal(data)
	io.WriteString(res, string(json))
	fmt.Println("respuesta exitosa")
}

func router() {
	http.HandleFunc("/list", GetAll)
	http.HandleFunc("/knn", Knn)
	http.ListenAndServe(":8080", nil)
}

func main() {
	router()

}
