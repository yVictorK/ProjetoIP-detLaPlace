	package main

	import (
		"fmt"
		"math/rand"
		"time"
	)

	func imprimeMatriz(mat [][]int) {
		var contI int
		var contJ int
		for contI = 0; contI < len(mat); contI++ {
			for contJ = 0; contJ < len(mat[0]); contJ++ {
				fmt.Print(mat[contI][contJ], " ")
			}
			fmt.Println()
		}
	}

	func iniciaMatrizRandomico(mat [][]int) {
		var contI int
		var contJ int
		for contI = 0; contI < len(mat); contI++ {
			for contJ = 0; contJ < len(mat[0]); contJ++ {
				mat[contI][contJ] = rand.Intn((len(mat) * len(mat)) / 2 + 1)
			}
		}
	}

	func verificaQuadradaOrdem(mat [][]int) (bool, int) {
		var numLinhas int
		var numColunas int
		var ehQuadrada bool

		numLinhas = len(mat)
		numColunas = len(mat[0])

		ehQuadrada = false
		if numLinhas == numColunas {
			ehQuadrada = true
		}

		return ehQuadrada, numLinhas
	}

	func calculaSinal(indiceL int, indiceC int) int {
		var sinal int

		sinal = -1
		if ((indiceL + indiceC) % 2) == 0 {
			sinal = 1
		}

		return sinal
	}

	func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
		var contAi, contAj, contBi, contBj, temp, numL, numC int
		numL = len(menor)
		numC = len(menor[0])

		contAi = 0
		for contBi = 0; contBi < numL; contBi++ {
			if contAi == isqn {
				contAi++
			}
			contAj = 0
			for contBj = 0; contBj < numC; contBj++ {
				if contAj == jsqn {
					contAj++
				}
				temp = maior[contAi][contAj]
				menor[contBi][contBj] = temp
				contAj++
			}
			contAi++
		}
	}

	func detOrdem1(mat [][]int) int {
		return mat[0][0]
	}

	func detOrdem2(mat [][]int) int {
		var diagonalP int
		var diagonalI int

		diagonalP = mat[0][0] * mat[1][1]
		diagonalI = mat[1][0] * mat[0][1]
		return (diagonalP - diagonalI)
	}

	func detOrdemN(mat [][]int) int {
		var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
		var matMenor [][]int
		numL = len(mat)
		numC = len(mat[0])

		resposta = 0
		contL = 0
		for contC = 0; contC < numC; contC++ {
			cofator = mat[contL][contC]
			sinal = calculaSinal(contL, contC)
			matMenor = make([][]int, numL-1)
			for cont = 0; cont < (numL - 1); cont++ {
				matMenor[cont] = make([]int, numC-1)
			}

			copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
			detTemp = determinante(matMenor)
			resposta = resposta + (cofator * sinal * detTemp)
		}

		return resposta
	}

	func verificaLinhaOuColunaComMaisZeros(mat [][]int) (bool, int) {
		var contL, contC int
		var maxZerosLinha, indiceLinhaMaisZeros, qntdZerosLinha int
		var maxZerosColuna, indiceColunaMaisZeros, qntdZerosColuna int
		var indiceMaisZeros int
		var ehLinha bool
		//inicializa quantidade de zeros com -1 para evitar erros
		maxZerosLinha = -1
		maxZerosColuna = -1
		ehLinha = false
		//busca a linha com mais zeros
		for contL = 0; contL < len(mat); contL++ {
			qntdZerosLinha = 0
			for contC = 0; contC < len(mat[0]); contC++ {
				if mat[contL][contC] == 0 {
					qntdZerosLinha++
				}
			}
			if qntdZerosLinha > maxZerosLinha {
				indiceLinhaMaisZeros = contL
				maxZerosLinha = qntdZerosLinha
			}
		}
		//busca coluna com mais zeros
		for contC = 0; contC < len(mat); contC++ {
			qntdZerosColuna = 0
			for contL = 0; contL < len(mat); contL++ {
				if mat[contL][contC] == 0 {
					qntdZerosColuna++
				}
			}
			if qntdZerosColuna > maxZerosColuna {
				indiceColunaMaisZeros = contC
				maxZerosColuna = qntdZerosColuna
			}
		}
		if maxZerosLinha > maxZerosColuna {
			indiceMaisZeros = indiceLinhaMaisZeros
			ehLinha = true
		} else {
			indiceMaisZeros = indiceColunaMaisZeros
		}

		return ehLinha, indiceMaisZeros
	}

	func detOrdemNOtimizado(mat [][]int) int {
		var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
		var matMenor [][]int
		var indiceMaisZeros int
		var ehLinha bool
		numL = len(mat)
		numC = len(mat[0])
		resposta = 0
		// verifica se é um alinha ou coluna que possui mais zeros
		ehLinha, indiceMaisZeros = verificaLinhaOuColunaComMaisZeros(mat)
		//faz o calculo do determinante com base na linha ou coluna com mais zeros
		if ehLinha {
			for contC = 0; contC < numC; contC++ {
				cofator = mat[indiceMaisZeros][contC]
				if cofator != 0 {
					sinal = calculaSinal(indiceMaisZeros, contC)
					matMenor = make([][]int, numL-1)
					for cont = 0; cont < (numL - 1); cont++ {
						matMenor[cont] = make([]int, numC-1)
					}
					copiaMatrizMaiorParaMenor(mat, matMenor, indiceMaisZeros, contC)
					detTemp = determinante(matMenor)
					resposta = resposta + (cofator * sinal * detTemp)
				}
			}
		} else {
			for contL = 0; contL < numL; contL++ {
				cofator = mat[contL][indiceMaisZeros]
				if cofator != 0 {
					sinal = calculaSinal(contL, indiceMaisZeros)
					matMenor = make([][]int, numL-1)
					for cont = 0; cont < (numL - 1); cont++ {
						matMenor[cont] = make([]int, numC-1)
					}
					copiaMatrizMaiorParaMenor(mat, matMenor, contL, indiceMaisZeros)
					detTemp = determinante(matMenor)
					resposta = resposta + (cofator * sinal * detTemp)
				}
			}
		}

		return resposta
	}

	func determinante(mat [][]int) int {
		var ordem int
		var ehQuadrada bool
		var det int

		ehQuadrada, ordem = verificaQuadradaOrdem(mat)
		det = 0
		if ehQuadrada {
			switch ordem {
			case 1:
				//fmt.Println("Ordem 1")
				det = detOrdem1(mat)
			case 2:
				//fmt.Println("Ordem 2")
				det = detOrdem2(mat)
			default:
				//fmt.Println("Ordem ", ordem)
				det = detOrdemN(mat)

			}
			// imprimeMatriz(mat)
			// fmt.Println("Det ", det)

		} else {
			fmt.Println("Matriz nao eh quadrada!! retornando 0")
		}
		return det
	}

	func determinanteOtimizado(mat [][]int) int {
		var ordem int
		var ehQuadrada bool
		var det int

		ehQuadrada, ordem = verificaQuadradaOrdem(mat)
		det = 0
		if ehQuadrada {
			switch ordem {
			case 1:
				//fmt.Println("Ordem 1")
				det = detOrdem1(mat)
			case 2:
				//fmt.Println("Ordem 2")
				det = detOrdem2(mat)
			default:
				//fmt.Println("Ordem ", ordem)
				det = detOrdemNOtimizado(mat)

			}
			// imprimeMatriz(mat)
			// fmt.Println("Det ", det)

		} else {
			fmt.Println("Matriz nao eh quadrada!! retornando 0")
		}
		return det
	}

	func main() {
		var contOrdem int
		var numRepeticoes int
		var contRepeticoes int
		var inicio time.Time
		var fim time.Time
		var ordens []int
		var tempoBaseline []int64
		var tempoOtimizado []int64
		var tempoExperimento int64
		var numLinhas int
		var numColunas int
		var cont int
		var matrix [][]int
		//inicialização das variavies
		ordens = []int{3,5, 7, 9, 11}
		numRepeticoes = 3
		tempoBaseline = make([]int64, len(ordens))
		tempoOtimizado = make([]int64, len(ordens))

		for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
			numLinhas = ordens[contOrdem]
			numColunas = ordens[contOrdem]
			for contRepeticoes = 0; contRepeticoes < numRepeticoes; contRepeticoes++ {
				fmt.Print("Ordem: ", ordens[contOrdem])
				fmt.Println(" - Repeticao: ", contRepeticoes+1)
				matrix = make([][]int, numLinhas)
				for cont = 0; cont < numLinhas; cont++ {
					matrix[cont] = make([]int, numColunas)
				}
				// inicia e imprime a matriz
				iniciaMatrizRandomico(matrix)
				imprimeMatriz(matrix)
				//medir o tempo do baseline com a matriz
				inicio = time.Now()
				fmt.Println("Determinante baseline: ", determinante(matrix))
				fim = time.Now()
				tempoExperimento = fim.UnixNano() - inicio.UnixNano()
				tempoBaseline[contOrdem] = tempoBaseline[contOrdem] + tempoExperimento

				//medir o tempo da otimizada com a mesma matriz anterior
				inicio = time.Now()
				fmt.Println("Determinante otimizado: ", determinanteOtimizado(matrix))
				fim = time.Now()
				tempoExperimento = fim.UnixNano() - inicio.UnixNano()
				tempoOtimizado[contOrdem] = tempoOtimizado[contOrdem] + tempoExperimento
			}
			tempoBaseline[contOrdem] = tempoBaseline[contOrdem] / int64(numRepeticoes)
			tempoOtimizado[contOrdem] = tempoOtimizado[contOrdem] / int64(numRepeticoes)
		}
		fmt.Println()
		fmt.Println("Tempo médio baseline (ns):", tempoBaseline)
		fmt.Println("Tempo médio otimizado (ns):", tempoOtimizado)
	}
