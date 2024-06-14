package main

import (
	"bufio"
	"fmt"
	"os"
	lector "tp2/lector"
)

func main() {

	/*l := lector.CrearLector()

	l.Procesar("agregar_archivo test01.log")
	l.Procesar("ver_mas_visitados 3")*/

	entrada := bufio.NewScanner(os.Stdin)
	lectorDs := lector.CrearLector()

	for entrada.Scan() {
		comando := entrada.Text()
		
		instruccion, resultado, err := lectorDs.Procesar(comando)

		if err != nil {
			fmt.Println("ERROR", err)
		}

		switch instruccion {
		case "agregar_archivo":

			for _, ip := range resultado {
				fmt.Println("DoS: " + ip)
			}
			fmt.Println("OK")

		case "ver_visitantes":
			fmt.Println("Visitantes:")
			for _, ip := range resultado {
				fmt.Println("\t" + ip)
			}
			fmt.Println("OK")

		case "ver_mas_visitados":
			fmt.Println("Sitios más visitados:")
			for _, sitio := range resultado {
				if sitio != ""{
					fmt.Println("\t" + sitio)
				}
			}
			fmt.Println("OK")

		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}
