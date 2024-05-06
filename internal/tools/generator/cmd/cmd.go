package main

import "github.com/machi12/nas/internal/tools/generator"

func main() {
	generator.ParseSpecs()

	generator.GenerateNasMessage()

	generator.GenerateNasEncDec()

	generator.GenerateTestLarge()
}
