// ###############################################################################################################################################################################
// package main

// import "fmt"

// func main() {
// 	var prato string
// 	fmt.Println("Digite seu prato preferido...")
// 	fmt.Println("P - Pizza")
// 	fmt.Println("H - Hambúrguer")
// 	fmt.Println("B - Bife com fritas")
// 	fmt.Println("S - Salada Caesar")
// 	fmt.Println("F - Salada de Frutas")
// 	fmt.Println("E - Estrogonofe")
// 	fmt.Println("O - Outros")
// 	fmt.Scan(&prato)

// 	switch prato {
// 	case "B":
// 		fmt.Println("Com batatas Palito ou Noisete?")
// 	case "H":
// 		fmt.Println("Com Queijo ou com Ovo?")
// 	case "P":
// 		fmt.Println("Calabresa ou Napolitana?")
// 	case "S":
// 		fmt.Println("Alface ou Rúcula?")
// 	case "F":
// 		fmt.Println("Kiwi ou Frango?")
// 	case "E":
// 		fmt.Println("Carne ou Frango?")
// 	case "O":
// 		fmt.Println("Não gostou de nosso cardápio?")
// 	default:
// 		fmt.Println("Não entendi seu paladar...")
// 	}
// }
// ###############################################################################################################################################################################
// package main

// import "fmt"

// func main() {
// 	cidade, populacao, capital := devolveCidadeEPopulacao()
// 	if capital {
// 		fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
// 	} else {
// 		fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
// 	}
// }

// func devolveCidadeEPopulacao() (string, int, bool) {
// 	return "Vila Sem Nome", 4328, true
// }
// ###############################################################################################################################################################################
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	estados := devolveEstadosDoSudeste()
// 	fmt.Println(estados)
// }

// func devolveEstadosDoSudeste() [4]string {
// 	var estados [4]string
// 	estados[0] = "RJ"
// 	estados[1] = "SP"
// 	estados[2] = "MG"
// 	estados[3] = "ES"
// 	return estados
// }
// ###############################################################################################################################################################################
// package main
// package main

// import "fmt"

// func main() {
// 	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
// 	for i, ponto := range pontosPlanningPoker {
// 		fmt.Println("O ponto na posição", i, " tem o valor", ponto)
// 	}
// }

// ###############################################################################################################################################################################
// package main

// import (
// 	"fmt"
// 	"time"
// )

// const aquecimento = 5
// const corridaLeve = 10
// const corridaForte = 10
// const alongamento = 1

// func main() {
// 	fmt.Println("Período de alongamento...")
// 	time.Sleep(alongamento * time.Minute)

// 	fmt.Println("Período de aquecimento...")
// 	time.Sleep(aquecimento * time.Minute)

// 	fmt.Println("Período de corrida leve...")
// 	time.Sleep(corridaLeve * time.Minute)

// 	fmt.Println("Período de corrida forte...")
// 	time.Sleep(corridaForte * time.Minute)

// 	fmt.Println("Período de alongamento...")
// 	time.Sleep(alongamento * time.Minute)
// }

// ###############################################################################################################################################################################
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	arquivo, err := os.Open("sites.txt")
// 	if err != nil {
// 		fmt.Println("Ocorreu um erro:", err)
// 	}
// 	leitor := bufio.NewReader(arquivo)
// 	linha, _ := leitor.ReadString('\n')

// 	fmt.Println(linha)
// }
// ###############################################################################################################################################################################