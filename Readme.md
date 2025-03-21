
# pinGOu

Um programa em Go para verificar a conectividade através do protocolo ICMP de uma lista de endereços IP.

## 🚀 Começando

Instruções para instalação do programa para fins de teste do seu funcionamento.

Consulte **[Implantação](#-implantação)** para saber como implantar o projeto.

### 📋 Pré-requisitos

Você precisa ter o Go instalado no seu sistema.

```
# No Linux (Ubuntu/Debian)
sudo apt install golang

# No Arch Linux
yay -S go

```
Além disso, é necessária a criação de um arquivo `ips.txt` contendo os endereços IP que deseja testar, um por linha.

### 🔧 Instalação

Clone o repositório e navegue até o diretório do projeto:

```
git clone git@github.com:windjhammer/pinGOu.git
cd pinGOu
```

Compile o programa:

```
go build -o pinGOu
```

Ou execute diretamente sem compilar:

```
go run main.go
```

Para implantar o programa em um ambiente ativo, basta compilar e mover o executável para um local acessível no sistema, como `/usr/local/bin/` no linux ou qualquer pasta que possua permissão de execução no *PATH* no windows.


Agora você pode rodar o programa de qualquer lugar:

```
pinGOu
```


