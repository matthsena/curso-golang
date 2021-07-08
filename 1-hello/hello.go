package main // indica que esse é o pacote principal da aplicação

import "fmt"
import "reflect" // pacote para descobrir tipo de variável

func main() {
    // para declarar variável em go usamos a estrutura
    // var nome_variavel tipo = valor
    // var nome string = "Matheus"
    var versao float32 = 1.1
    var idade int
    // Golang da erro de compilação caso uma variável tenha
    // sido declarada, porém não utilizada

    // inferencia de tipos
    // o Go consegue inferir o tipo da variável
    // ao invés de `var nome string = "Matheus"`,
    // podemos simplesmente utilizar:

    // var nome = "Matheus"
    // não precisamos utilizar o nome var, podemos utilizar
    // um encurtador e definir como:
    nome := "Matheus" // forma mais usual


    // em Go, toda função de um pacote externo é invocada com
    // nome-pacote.Função, sempre com a primeira letra maiúscula
    // para indicar que pertence a um pacote externo
    fmt.Println("Olá!", nome, "sua idade é", idade)
    fmt.Println("Esse programa está na versão", versao)
    fmt.Println("O tipo de nome é", reflect.TypeOf(nome))
}
