package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/////Structs
type mbr struct {
	Numero   uint8
	Caracter byte
	Cadena   [20]byte
}

type mbra struct {
	Size      uint64
	Date      [16]byte
	Signature uint8
	Primary   [4]partition
	Extended  [4]partition
}
type partition struct {
	status   [1]byte
	tipo_par [1]byte
	fit      [1]byte
	start    uint16
	size     uint64
	name     [16]byte
}

////MAIN
func main() {
	//create(10, "/home/autoestima/", "name.disk", "k")
	//var date [16]byte
	asd := time.Now()
	fe := asd.Format("2006-01-02 15:04:00")
	fmt.Println(fe)
	disco := mbra{}

	a := time.Now()

	fmt.Println(a)
	//tiempo := strconv.Itoa(fecha.Day()) + "/" + strconv.Itoa(fecha.Year()) + " " + strconv.Itoa(fecha.Hour()) + ":" + strconv.Itoa(fecha.Minute()) + ":" + strconv.Itoa(fecha.Second())
	//t := time.Now()
	//fmt.Println(t)
	//date[0] = t.Year()
	//date[1] = t.Month()
	//date[2] = t.Day()
	//da/te[3] = t.Hour()
	//date[4] = t.Minute()
	//da//te[5] = t.Second()
	//hi()
	consola()
}

//////////////METODOS PARA CONSOLA
func consola() {
	finalizar := 0
	salida := "exit"
	mensaje := "Bienvenido a la consola del sistema de archivos LWH \n(exit para salir)"
	fmt.Println(mensaje)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	comando, _ := reader.ReadString('\n')
	comando = strings.ReplaceAll(comando, "\n", "")
	comando = leer(comando)
	fmt.Print(comando + "\n")

	if comando == salida {
		finalizar = 1

	} else {

		if comando != "" {
			lineaComando(comando)
		}
	}

	for finalizar != 1 {
		fmt.Println("Insertar comando: ")
		reader := bufio.NewReader(os.Stdin)
		comando, _ := reader.ReadString('\n')
		comando = strings.ReplaceAll(comando, "\n", "")
		comando = leer(comando)
		fmt.Print(comando + "\n")
		if comando == salida {
			finalizar = 1
			break

		} else {
			if comando != "" {
				lineaComando(comando)
			}
		}
		if finalizar == 1 {
			break
		}
	}
}

//Recibe la linea
func lineaComando(comando string) {
	var commandArray []string
	commandArray = strings.Split(comando, " ")
	//fmt.Println(commandArray[1])
	ejecutarComando(commandArray) //Ejecutamos el comando.
}

//Recibe array del comando (pos 0) comando y el resto son atributos
func ejecutarComando(commandArray []string) {
	data := strings.ToLower(commandArray[0])
	if data == "exec" {
		exec(commandArray)

	} else if data == "pause" {
		pause(commandArray)

	} else if data == "mkdisk" {

		mkdisk(commandArray)

	} else if data == "rmdisk" {
		rmdisk(commandArray)
	} else if data == "fdisk" {
		fmt.Println("fdisk Comando")
	} else if data == "mount" {
		fmt.Println("mount Comando")
	} else if data == "unmount" {
		fmt.Println("unmount Comando")
	} else {
		if len(data) > 0 {
			if data[0] == '#' {
				//fmt.Println(data)
			} else {
				fmt.Println("Comando no existe")
			}

		}

	}

}

//Metodo que concatena \* con cada linea
func leer(comando string) string {

	if strings.Contains(comando, "\\*") {

		comando = strings.ReplaceAll(comando, "\\*", " ")
		reader := bufio.NewReader(os.Stdin)
		a, _ := reader.ReadString('\n')
		a = strings.ReplaceAll(a, "\n", "")
		comando += leer(a)

	}

	return comando
}

func exec(commandArray []string) {
	if len(commandArray) == 2 {
		var atr []string
		atr = strings.Split(commandArray[1], "->")
		p := strings.ToLower(atr[0])
		if p == "-path" {
			file_data, err := ioutil.ReadFile("./prueba.sh")

			if err != nil {
				fmt.Println("Error al abrir el path")
			}
			data := string(file_data)
			var comandosArray []string
			comandosArray = strings.Split(data, "\n")
			command := ""
			for i := 0; i < len(comandosArray); i++ {
				command = comandosArray[i]
				if strings.Contains(comandosArray[i], "\\*") {
					command = strings.ReplaceAll(comandosArray[i], "\\*", " ")
					i++
				}

				for strings.Contains(comandosArray[i], "\\*") {
					command += strings.ReplaceAll(comandosArray[i], "\\*", " ")
					i++

				}
				if command == "\n" || command == "" {

				} else {
					fmt.Println(command)
					lineaComando(command)
					command = ""
				}
				command = ""

			}

		} else {
			fmt.Println("Error en atributo path")
		}
	} else {
		fmt.Println("Numero de parametros incorrectos")
	}
}
func pause(commandArray []string) {
	if len(commandArray) == 1 {
		fmt.Println("Press the any key to terminate the console screen!")
		fmt.Scanln() // wait for Enter Key
	} else {
		fmt.Println("Error pause no tiene atributos")
	}
}
func mkdisk(commandArray []string) {
	if len(commandArray) == 5 || len(commandArray) == 4 {
		cSize := false
		cPath := false
		cName := false
		//cUnit := false

		var size int
		var path string
		var name string
		var unit string = "m"

		for i := 1; i < len(commandArray); i++ {
			parameter := strings.Split(commandArray[i], "->")
			p := strings.ToLower(parameter[0])
			if p == "-path" {
				cPath = true
				path = parameter[1]
			} else if p == "-size" {
				cSize = true
				size, _ = strconv.Atoi(parameter[1])
				if size < 0 {
					cSize = false
					fmt.Println("Error en el tamaño")
					return
				}

			} else if p == "-name" {
				cName = true
				name = parameter[1]
				if strings.Contains(name, ".disk") {

				} else {
					cName = false
					fmt.Println("Error en el nombre")
					return

				}
			} else if p == "-unit" {
				unit = strings.ToLower(parameter[1])
				if unit == "k" || unit == "m" {

				} else {
					fmt.Println("Error unit invalido, sera m(predeterminado)")
					unit = "k"

				}

			}

		}

		if cSize == true && cPath == true && cName == true {
			create(int64(size), path, name, unit)
		} else {
			fmt.Println("No tiene los parametros obligatorios")
		}
	} else {
		fmt.Println("Numero de atributos incorrecto")
	}
}
func rmdisk(commandArray []string) {
	if len(commandArray) == 2 {

		parameter := strings.Split(commandArray[1], "->")
		p := strings.ToLower(parameter[0])
		if p == "-path" {

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Confirmacion de eliminacion de disco " + parameter[1] + " (Y/N)")
			conf, _ := reader.ReadString('\n')
			conf = strings.ReplaceAll(conf, "\n", "")
			if conf == "Y" || conf == "y" {
				nombreArchivo := "hola.txt" // El nombre o ruta absoluta del archivo
				err := os.Remove(nombreArchivo)
				if err != nil {
					fmt.Printf("Error eliminando archivo: %v\n", err)
				} else {
					fmt.Println("Eliminado correctamente")
				}
			} else {

			}
		} else {
			fmt.Println("rmdisk solo recibe de parametro path")
		}

	} else {
		fmt.Println("Numero de atributos incorrecto")
	}

}
func fdisk(commandArray []string) {

}
func mount(commandArray []string) {

}
func unmount(commandArray []string) {

}

//////////////////METODOS creacio disco (archivo binario)
func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

func create(size int64, path, name, unit string) {
	if unit == "m" {
		size = size * 1024 * 1024
	} else {
		size = size * 1024
	}

	err := os.Mkdir(path, 0755)

	file, err := os.Create(path + name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var otro int8 = 0

	s := &otro

	fmt.Println(unsafe.Sizeof(otro))
	//Escribimos un 0 en el inicio del archivo.
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	escribirBytes(file, binario.Bytes())

	file.Seek(size-int64(1), 0)

	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	escribirBytes(file, binario2.Bytes())

	file.Seek(0, 0)

	time := time.Now()
	fecha := time.Format("2006-01-02 15:04:00")
	random := uint8(rand.Intn(250))

	disco := mbra{}
	copy(disco.Date[:], fecha)
	disco.Size = uint64(size)
	disco.Signature = random

	s1 := &disco

	//Escribimos struct.
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, s1)
	escribirBytes(file, binario3.Bytes())

}
