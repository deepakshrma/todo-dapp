package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/deepakshrma/tododapp/controllers"
	"github.com/deepakshrma/tododapp/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ContextBuilder(eth *ethclient.Client, envs map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ethClient", eth)
		c.Set("auth", utils.GetAuth(eth, envs["PRIVATE_KEY"]))
		c.Set("private_key", envs["PRIVATE_KEY"])
		c.Set("contract_address", envs["CONTRACT_ADDRESS"])
		c.Next()
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

const ( // iota is reset to 0
	NewTodo     = iota // c0 == 0
	DoneTodo    = iota // c1 == 1
	DeletedTodo = iota // c2 == 2
)

var envs map[string]string

func main() {
	err := godotenv.Load(".env")
	envs, err = godotenv.Read()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.Use(ContextBuilder(client, envs))
	v1.Use(CORSMiddleware())
	v1.GET("/configs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"CONTRACT_ADDRESS": envs["CONTRACT_ADDRESS"],
		})
	})
	v1.GET("/deploy", controllers.DeployTodo)
	v1.GET("/todos/:id", controllers.GetTodoById)
	v1.POST("/todos/", controllers.GetAllTodos)
	v1.POST("/todos/:id", controllers.DeleteTodo)
	v1.GET("/todos/", controllers.CreateNewTodo)
	v1.GET("/noOfTodos", controllers.FindNumberOfTodos)
	defer subscribe(client)
	r.Run(utils.Getenv("PORT", ":8080"))

}
func subscribe(client *ethclient.Client) {
	// contractAddress := common.HexToAddress(envs["CONTRACT_ADDRESS"])
	query := ethereum.FilterQuery{
		// Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
