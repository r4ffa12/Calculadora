# Calculadora Simples em Go

Este projeto é uma calculadora simples desenvolvida em Go (Golang). O programa permite ao usuário inserir uma operação matemática no formato:

```
numero operador numero
```

Exemplo:
```
12 + 34
```
E o programa retornará o resultado da operação.

## Tecnologias Utilizadas
- Linguagem: **Go (Golang)**
- Biblioteca **bufio** para leitura da entrada do usuário
- Biblioteca **strings** para manipulação de texto
- Biblioteca **strconv** para conversão de strings em números

## Como Executar

1. Certifique-se de ter o Go instalado.
2. Clone este repositório:
   ```sh
   git clone https://github.com/seu-usuario/calculadora-go.git
   ```
3. Acesse a pasta do projeto:
   ```sh
   cd calculadora-go
   ```
4. Execute o programa:
   ```sh
   go run main.go
   ```

## Funcionamento do Código

O código segue os seguintes passos:

### 1. Importação de Pacotes

```go
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
```
- **fmt**: Exibe mensagens e recebe entrada do usuário.
- **bufio**: Permite ler a entrada completa do usuário.
- **os**: Necessário para entrada via terminal.
- **strconv**: Converte strings para números inteiros.
- **strings**: Manipula a string digitada pelo usuário.

### 2. Entrada do Usuário

```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter the operation (e.g., 12 + 34):")
input, _ := reader.ReadString('\n')
input = strings.TrimSpace(input)
```
- Lê a linha digitada pelo usuário.
- Remove espaços extras no início e final da entrada.

### 3. Separação dos Elementos

```go
splitInput := strings.Fields(input)
if len(splitInput) != 3 {
	fmt.Println("Invalid format. Use: number operator number")
	return
}
```
- Divide a entrada em partes (número, operador, número).
- Verifica se o usuário digitou a operação corretamente.

### 4. Conversão de Números

```go
num1, err1 := strconv.Atoi(splitInput[0])
num2, err2 := strconv.Atoi(splitInput[2])
operator := splitInput[1]

if err1 != nil || err2 != nil {
	fmt.Println("Invalid numbers")
	return
}
```
- Converte os números digitados para inteiros.
- Se a conversão falhar, exibe erro.

### 5. Cálculo do Resultado

```go
result, err := getResult(num1, num2, operator)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
```
- Chama a função `getResult()` para calcular o resultado.
- Se houver erro (divisão por zero ou operador inválido), ele é tratado.
- Exibe o resultado final.

### 6. Função de Cálculo

```go
func getResult(num1, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Operator not valid")
	}
}
```
- Realiza o cálculo com base no operador digitado.
- Retorna erro se o operador for inválido ou houver divisão por zero.

## Exemplos de Entrada e Saída

| Entrada       | Saída         |
|--------------|--------------|
| `10 + 5`    | `10 + 5 = 15` |
| `20 - 4`    | `20 - 4 = 16` |
| `7 * 6`     | `7 * 6 = 42` |
| `50 / 2`    | `50 / 2 = 25` |
| `10 / 0`    | `Cannot divide by zero` |
| `abc + 3`   | `Invalid numbers` |
| `12 & 4`    | `Operator not valid` |

## Contribuição
Se desejar contribuir para este projeto, sinta-se à vontade para fazer um **fork** e enviar um **pull request**!

## Licença
Este projeto está sob a licença MIT. Sinta-se à vontade para usá-lo e modificá-lo.

