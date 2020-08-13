package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	consola()
}

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
		fmt.Println("Otro Comando")
	} else if data == "rmdisk" {
		fmt.Println("Otro Comando")
	} else if data == "pause" {
		fmt.Println("Otro Comando")
	} else if data == "fdisk" {
		fmt.Println("Otro Comando")
	} else if data == "mount" {
		fmt.Println("Otro Comando")
	} else if data == "unmount" {
		fmt.Println("Otro Comando")
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

}
func pause(commandArray []string) {

}
func mkdisk(commandArray []string) {

}
func rmdisk(commandArray []string) {

}
func fdisk(commandArray []string) {

}
func mount(commandArray []string) {

}
func unmount(commandArray []string) {

}
