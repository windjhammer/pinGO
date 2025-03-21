package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

func testConnection(ip string) bool {
	conn, err := net.DialTimeout("tcp", ip+":80", 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	arquivoIps := "ips.txt"
	arquivoResultado := "resultado.txt"

	if _, err := os.Stat(arquivoIps); os.IsNotExist(err) {
		fmt.Println("O arquivo ips.txt não foi encontrado no diretório do script.")
		return
	}

	data, err := ioutil.ReadFile(arquivoIps)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo ips.txt:", err)
		return
	}

	ips := strings.Split(string(data), "\n")
	var resultados []string

	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip == "" {
			continue
		}

		if testConnection(ip) {
			resultado := fmt.Sprintf("Ping para %s: Sucesso", ip)
			fmt.Println(resultado)
			resultados = append(resultados, resultado)
		} else {
			resultado := fmt.Sprintf("Ping para %s: Falhou", ip)
			fmt.Println(resultado)
			resultados = append(resultados, resultado)
		}
	}

	err = ioutil.WriteFile(arquivoResultado, []byte(strings.Join(resultados, "\n")), 0644)
	if err != nil {
		fmt.Println("Erro ao salvar os resultados:", err)
	} else {
		fmt.Println("Resultados salvos em", arquivoResultado)
	}
}
