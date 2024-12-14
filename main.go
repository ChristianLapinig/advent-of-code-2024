package main

import "fmt"

func main() {
	star()
	top()
	for i := 0; i < 3; i++ {
		middle()
	}
	trunk()
}

func star() {
	fmt.Println("           *           ")
	fmt.Println("          * *          ")
	fmt.Println("         * * *         ")
	fmt.Println("          * *          ")
}

func top() {
	fmt.Println("           *          ")
	fmt.Println("         *   *        ")
	fmt.Println("        *  *  *       ")
	fmt.Println("      *  *   *  *     ")
	fmt.Println("   *   *   *   *   *  ")
}

func middle() {
	fmt.Println("         *   *          ")
	fmt.Println("        *  *  *         ")
	fmt.Println("      *  *   *  *       ")
	fmt.Println("   *   *   *   *   *    ")
}

func trunk() {
	fmt.Println("         *   *          ")
	fmt.Println("        *     *         ")
	fmt.Println("      *         *       ")
	fmt.Println("    * * * * * * * *     ")
}
