package main

import (
	"backed/app/models"
	"backed/internal/pkg"
)

func main() {
	generator := pkg.NewGenerator("database")

	generator.ApplyBasic(&models.UserInfo{})
	generator.ApplyBasic(&models.File{})
	generator.ApplyBasic(&models.GoodsCard{})
	generator.ApplyBasic(&models.FruitKind{})

	generator.Execute()
}
