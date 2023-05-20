package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Bem-vindo ao jogo de adivinhação!")

    var allAttempts [][]int

    for {
        answer := rand.Intn(100) + 1
        fmt.Println("Vou gerar um número aleatório entre 1 e 100. Tente adivinhar qual é!")

        attempts := 0
        userAttempts := []int{}

        for {
            fmt.Print("Digite um número inteiro: ")
            var guess int
            fmt.Scanln(&guess)

            attempts++
            userAttempts = append(userAttempts, guess)

            if guess == answer {
                fmt.Printf("Parabéns! Você acertou o número em %d tentativa(s).\n", attempts)
                break
            } else if guess < answer {
                fmt.Println("O número é maior que", guess)
            } else {
                fmt.Println("O número é menor que", guess)
            }
        }

        allAttempts = append(allAttempts, userAttempts)

        fmt.Print("Deseja jogar novamente? (s/n): ")
        var playAgainInput string
        fmt.Scanln(&playAgainInput)

        if playAgainInput != "s" {
            break
        }

        fmt.Println()
    }

    fmt.Println("Fim do jogo. Obrigado por jogar!")

    fmt.Println("Histórico de tentativas:")
    for i, attempts := range allAttempts {
        fmt.Printf("Jogada %d: ", i+1)
        for _, guess := range attempts {
            fmt.Printf("%d ", guess)
        }
        fmt.Println()
    }
}
