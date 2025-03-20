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

	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip == "" {
			continue
		}

		if testConnection(ip) {
			fmt.Printf("Ping para %s: Sucesso\n", ip)
		} else {
			fmt.Printf("Ping para %s: Falhou\n", ip)
		}
	}
}
