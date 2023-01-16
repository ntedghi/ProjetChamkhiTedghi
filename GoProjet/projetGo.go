package main

import "fmt"

func main() {
	var lignes, colonnes, i, j int
	var mat1 [10][10]int
	var mat2 [10][10]int
	var multiMat [10][10]int

	fmt.Print("Enter the Multiplication Matrix Rows and Columns = ")
	fmt.Scan(&lignes, &colonnes)

	fmt.Print("Enter the First Matrix Items to Multiplication = ")
	for i = 0; i < colonnes; i++ {
		for j = 0; j < colonnes; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}
	fmt.Print("Enter the Second Matrix Items to Multiplication = ")
	for i = 0; i < lignes; i++ {
		for j = 0; j < colonnes; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}

	for i = 0; i < lignes; i++ {
		for j = 0; j < colonnes; j++ {
			multiMat[i][j] = mat1[i][j] * mat2[i][j]
		}
	}
	fmt.Println("The Go Result of Matrix Multiplication = ")
	for i = 0; i < lignes; i++ {
		for j = 0; j < colonnes; j++ {
			fmt.Print(multiMat[i][j], "\t")
		}
		fmt.Println()
	}
}

//ET ON TESTETSTET SI CAMARCEH
