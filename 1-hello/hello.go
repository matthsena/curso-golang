package main // indica que esse é o pacote principal da aplicação

import "fmt"

func main() {
    // em Go, toda função de um pacote externo é invocada com
    // nome-pacote.Função, sempre com a primeira letra maiúscula
    // para indicar que pertence a um pacote externo

    fmt.Println("Hello World!")

}
