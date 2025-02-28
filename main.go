package main

import (
	"log"
	"net/http"
	"slices"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

var packSizes = []int{
	250,
	500,
	1000,
	2000,
	5000,
}

type Response struct {
	Packs map[int]int `json:"packs"`
}

func main() {
	router := gin.Default()
	router.GET("/packs/:quantity", handleInput)
	error := router.Run()
	if error != nil {
		log.Fatal(error)
	}
}

func handleInput(context *gin.Context) {
	quantity, error := strconv.Atoi(context.Param("quantity"))
	if error == nil && quantity >= 0 {
		context.IndentedJSON(http.StatusOK, getMinimumPacks(quantity))
	} else {
		context.IndentedJSON(http.StatusBadRequest, "Invalid Input")
	}
}

func getMinimumPacks(quantity int) Response {
	var smallestSize int = slices.Min(packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	var modResult = quantity % smallestSize
	var minimumQuantity int
	if modResult == 0 {
		minimumQuantity = quantity
	} else {
		minimumQuantity = (smallestSize - (quantity % smallestSize)) + quantity
	}

	var packs map[int]int = make(map[int]int)

	for _, size := range packSizes {
		if minimumQuantity >= size {
			var numPacks int = minimumQuantity / size
			minimumQuantity %= size
			packs[size] = numPacks
		}
	}

	return Response{Packs: packs}
}
