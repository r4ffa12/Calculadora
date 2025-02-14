# Calculadora GO

1. Importação de Pacotes
go
Copiar
Editar
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
fmt: Usado para exibir mensagens e capturar entrada do usuário.
bufio: Usado para ler a entrada do usuário com mais flexibilidade.
os: Permite interagir com o sistema operacional, como ler a entrada do teclado.
strconv: Fornece funções para converter texto (string) em números inteiros.
strings: Contém funções para manipular strings, como remover espaços em branco.
2. Função main()
A função principal do programa executa os seguintes passos:

a) Criar um leitor para capturar a entrada do usuário
go
Copiar
Editar
reader := bufio.NewReader(os.Stdin)
Isso permite que o programa leia uma linha completa digitada pelo usuário, incluindo espaços.
b) Exibir a mensagem de entrada
go
Copiar
Editar
fmt.Println("Enter the operation (e.g., 12 + 34):")
Isso pede ao usuário para inserir uma operação matemática no formato correto.
c) Capturar e limpar a entrada do usuário
go
Copiar
Editar
input, _ := reader.ReadString('\n')
input = strings.TrimSpace(input)
ReadString('\n'): Lê a linha inteira digitada pelo usuário até pressionar "Enter".
TrimSpace(input): Remove espaços em branco extras no início e no final da entrada.
3. Processamento da Entrada
a) Separar os números e o operador
go
Copiar
Editar
splitInput := strings.Fields(input)
strings.Fields(input): Divide a entrada em partes separadas por espaços. Exemplo:
Entrada: "12 + 34"
Resultado: ["12", "+", "34"] (lista com três elementos)
b) Verificar se a entrada está no formato correto
go
Copiar
Editar
if len(splitInput) != 3 {
	fmt.Println("Invalid format. Use: number operator number")
	return
}
Garante que a entrada tem exatamente três partes (número operador número).
Se o usuário digitou algo inválido (como "12 +" ou "12+34" sem espaços), o programa exibe um erro e termina.
c) Converter os números de string para inteiros
go
Copiar
Editar
num1, err1 := strconv.Atoi(splitInput[0])
num2, err2 := strconv.Atoi(splitInput[2])
operator = splitInput[1]
strconv.Atoi() converte um texto ("12") em número (12).
Se a conversão falhar (exemplo: usuário digitou "abc + xyz"), o programa não consegue continuar.
d) Verificar se os números são válidos
go
Copiar
Editar
if err1 != nil || err2 != nil {
	fmt.Println("Invalid numbers")
	return
}
Se algum dos números não for válido, o programa exibe um erro e para.
4. Chamar a Função de Cálculo
go
Copiar
Editar
result, err := getResult(num1, num2, operator)
if err != nil {
	fmt.Println(err)
	return
}
A função getResult(num1, num2, operator) faz o cálculo e retorna o resultado.
Se houver erro (como divisão por zero ou operador inválido), ele é tratado e exibido.
5. Exibir o Resultado
go
Copiar
Editar
fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
Exibe a operação completa e o resultado formatado corretamente. Exemplo:
txt
Copiar
Editar
12 + 34 = 46
6. Função getResult()
Essa função recebe os dois números e o operador e retorna o resultado ou um erro, se houver.

a) Estrutura de Decisão (switch case)
go
Copiar
Editar
switch operator {
Aqui, verificamos qual operação o usuário quer fazer.
b) Executar a Operação Correspondente
go
Copiar
Editar
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
Dependendo do operador (+, -, *, /), o cálculo é feito e o resultado retornado.
Se a operação for uma divisão e o segundo número for 0, o programa retorna um erro.
c) Lidar com Operadores Inválidos
go
Copiar
Editar
default:
	return 0, fmt.Errorf("Operator not valid")
Se o operador não for um dos permitidos (+, -, *, /), o programa exibe um erro.
Resumo do Fluxo do Programa
Usuário digita uma operação (Exemplo: 10 * 5).
O programa lê a entrada e separa os valores (10, *, 5).
Converte os números para inteiros e verifica erros.
Chama a função getResult() para calcular o resultado.
Exibe a resposta formatada (10 * 5 = 50).
Se houver erro, exibe uma mensagem apropriada.
Exemplos de Entrada e Saída
Entrada	Saída
10 + 5	10 + 5 = 15
20 - 4	20 - 4 = 16
7 * 6	7 * 6 = 42
50 / 2	50 / 2 = 25
10 / 0	Cannot divide by zero
abc + 3	Invalid numbers
12 & 4	Operator not valid
Esse código permite ao usuário fazer cálculos simples e lida com erros de entrada, garantindo que o formato da operação seja válido.
