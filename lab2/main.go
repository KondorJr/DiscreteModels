package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	N = 8
	K = 100000
)

func readMatrix() ([N][N]int, error) {
	var matrix [N][N]int
	bytes, err := ioutil.ReadFile("l2-2.txt")
	if err != nil {
		return [8][N]int{}, err
	}

	fmt.Println("Weight matrix")

	for i, v := range strings.Split(string(bytes), "\n") {
		if i == 0 {
			continue
		}

		for i2, v2 := range strings.Split(v, " ") {
			n, err := strconv.ParseInt(v2, 10, 32)
			if err != nil {
				return [8][N]int{}, err
			}

			fmt.Print(n, "\t")

			matrix[i-1][i2] = int(n)
		}

		fmt.Println()
	}

	return matrix, err
}

func sumOfEdgeVertices(matrix [N][N]int) int {
	sum := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum += matrix[i][j]
		}
	}

	return sum
}

func euler(matrix, conMatrix [N][N]int) [N][N]int {
	var cm [N][N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cm[i][j] = conMatrix[i][j]
		}
	}

	var (
		p      [K + 1]int
		p1     = 0
		k      = sumOfEdgeVertices(cm) / 2
		p2     = k + 1
		weight = 0
	)

	p[0] = 0

	for p1 >= 0 {
		var i, v = p[p1], p[p1]

		if p1 < 0 {
			break
		}

		for i = 0; i < N; {

			if cm[v][i] != 0 {
				cm[v][i] = cm[v][i] - 1
				cm[i][v] = cm[i][v] - 1

				p1 += 1
				if p1 == K+1 {
					p1 = -1
					p2 = 0
					break
				}

				p[p1] = i
				v = i
				i = 0
			} else {
				i++
			}

			if i >= N {
				if p1 < 0 || p2 < 0 {
					break
				}
				p2 -= 1

				p[p2] = p[p1]
				p1--

			}
		}
	}

	if p2 > 0 {
		return [8][N]int{}
	} else {
		for i := 0; i < k; i++ {
			fmt.Println(" ", p[i]+1, " - ", p[i+1]+1, " : ", matrix[p[i]][p[i+1]])
			weight += matrix[p[i]][p[i+1]]
		}
		fmt.Println("Weight: ", weight)
	}

	return cm
}

func addEdges(matrix, conMatrix [N][N]int, vd [N]int) int {
	for i := 0; i < N; i++ {
		if vd[i]%2 != 0 {
			for j := 0; j < N; j++ {
				if vd[j]%2 != 0 && conMatrix[i][j] != 0 {
					conMatrix[i][j] = conMatrix[i][j] + 1
				}
			}
		}
	}

	fmt.Println("Changed Connectivity Matrix")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(conMatrix[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
	euler(matrix, conMatrix)

	return 0
}

func connectivityMatrix(matrix [N][N]int) [N][N]int {
	fmt.Println("Connectivity Matrix")
	var conMatrix [N][N]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] > 0 {
				conMatrix[i][j] = 1
			} else {
				conMatrix[i][j] = matrix[i][j]
			}
			fmt.Print(conMatrix[i][j], "\t")
		}
		fmt.Println()
	}

	return conMatrix
}

func ifEuler(matrix, conMatrix [N][N]int, vd [N]int) int {
	for i := 0; i < N; i++ {
		if vd[i] != 0 {
			addEdges(matrix, conMatrix, vd)
			break
		} else {
			euler(matrix, conMatrix)
		}
	}

	return 0
}

func vertexDegree(matrix, conMatrix [N][N]int) [N]int {
	fmt.Println()
	var vd [N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			vd[i] += conMatrix[i][j]
		}
		if vd[i]%2 != 0 {
			i++
		}
	}

	ifEuler(matrix, conMatrix, vd)

	return vd
}

func main() {
	matrix, err := readMatrix()
	if err != nil {
		panic(err)
	}

	conMatrix := connectivityMatrix(matrix)

	vertexDegree(matrix, conMatrix)

}
