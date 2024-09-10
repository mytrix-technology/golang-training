package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mytrix-technology/golang-training/D.Expert/large-language-models/gin-langchain/routes"
)

func main() {
	r := gin.Default()
	routes.GetVacationRouter(r)
	r.Run()
}
