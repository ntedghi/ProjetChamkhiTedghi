package main

import "fmt"

func main() {
	var lignes, colonnes, i, j int

	var mat1 [10][10]int
	var mat2 [10][10]int
	var matMultiplier [10][10]int

	fmt.Print("Entrer le nombre de lignes puis celui de colonnes  = ")
	fmt.Scan(&lignes, &colonnes)

	fmt.Print("Enter the First Matrix Items to Multiplication = ")
	for i = 0; i < lignes; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&multimat1[i][j])
		}
	}

	fmt.Print("Enter the Second Matrix Items to Multiplication = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&multimat2[i][j])
		}
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			multiplicationnmat[i][j] = multimat1[i][j] * multimat2[i][j]
		}
	}
	fmt.Println("The Go Result of Matrix Multiplication = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(multiplicationnmat[i][j], "\t")
		}
		fmt.Println()
	}
}
