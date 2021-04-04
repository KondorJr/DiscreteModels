package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
import "strconv"

func main() {

	f, err := ioutil.ReadFile("l2-2.txt")
	if err != nil {
		panic(err)
	}

	var cities = [8][8]int{}

	for i, v := range strings.Split(string(f), "\n") {
		if i == 0 {
			continue
		}

		var tmp [8]int
		for i2, v2 := range strings.Split(v, " ") {
			number, err := strconv.ParseInt(v2, 16, 32)
			if err != nil {
				panic(err)
			}

			tmp[i2] = int(number)
		}

		cities[i-1] = tmp
	}

	gg := permutation([]int{1, 2, 3, 4, 5, 6, 7})

	total := 0

	shortestDistance := -1
	shortestPaths := make([]string, 0)

	for _, elem := range gg {

		fmt.Println("route:", routesToStr(elem))

		lastCity := 0

		for _, city := range elem {

			total += cities[lastCity][city]

			lastCity = city
		}

		total += cities[lastCity][0]

		fmt.Println("total distance:", total)

		if shortestDistance == -1 || shortestDistance > total {

			shortestDistance = total
			shortestPaths = append(shortestPaths, routesToStr(elem))
		}

		total = 0

	}

	fmt.Println()
	fmt.Println()
	fmt.Println("shortestDistance:", shortestDistance)
	fmt.Println("shortestPaths:", shortestPaths)
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func routesToStr(arr []int) string {

	result := "(0,"

	for _, o := range arr {
		result += strconv.Itoa(o)
		result += ","
	}

	result += "0)"

	return result

}
