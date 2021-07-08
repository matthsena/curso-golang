package main // indica que esse é o pacote principal da aplicação

import "fmt"

func main() {
    // para declarar variável em go usamos a estrutura
    // var nome_variavel tipo = valor
    var nome string = "Matheus"
    var versao float32 = 1.1
    var idade int
    // Golang da erro de compilação caso uma variável tenha
    // sido declarada, porém não utilizada

    // em Go, toda função de um pacote externo é invocada com
    // nome-pacote.Função, sempre com a primeira letra maiúscula
    // para indicar que pertence a um pacote externo
    fmt.Println("Olá!", nome, "sua idade é", idade)
    fmt.Println("Esse programa está na versão", versao)

}
