package controllers

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	todo "github.com/deepakshrma/tododapp/contracts"
	"github.com/deepakshrma/tododapp/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	contextKeyClient = "ethClient"
)

func CreateNewTodo(c *gin.Context) {
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	contractAddress := c.MustGet("contract_address").(string)

	address := common.HexToAddress(contractAddress)
	instance, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}

	todos, err := instance.GetTodos(nil)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})

}
func DeleteTodo(c *gin.Context) {
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	auth := c.MustGet("auth").(*bind.TransactOpts)
	contractAddress := c.MustGet("contract_address").(string)

	address := common.HexToAddress(contractAddress)
	instance, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}
	deleteId := big.NewInt(0)
	id, _ := deleteId.SetString(c.Param("id"), 10)

	instance.UpdateStatus(auth, id, 1)

	todo, _ := instance.GetTodo(nil, id)
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})

}

func DeployTodo(c *gin.Context) {
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	auth := c.MustGet("auth").(*bind.TransactOpts)
	privateKey := c.MustGet("private_key").(string)

	address, _, instance, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
	file := fmt.Sprintf(`
	PRIVATE_KEY=%s
	CONTRACT_ADDRESS=%s
	`, privateKey, address.Hex())
	env, _ := godotenv.Unmarshal(file)
	godotenv.Write(env, "./.env")
	c.JSON(http.StatusOK, gin.H{
		"data": address.Hex(),
	})
}

func FindNumberOfTodos(c *gin.Context) {
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	contractAddress := c.MustGet("contract_address").(string)

	address := common.HexToAddress(contractAddress)
	instance, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}

	noOfTodos, err := instance.NoOfTodos(nil)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"NoOfTodos": noOfTodos,
	})
}

func GetTodoById(c *gin.Context) {
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	contractAddress := c.MustGet("contract_address").(string)
	address := common.HexToAddress(contractAddress)
	instance, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}
	todoId := new(big.Int)
	id, _ := todoId.SetString(c.Param("id"), 10)

	todo, err := instance.GetTodo(nil, id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

func GetAllTodos(c *gin.Context) {
	var input models.TodoItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := c.MustGet(contextKeyClient).(*ethclient.Client)
	auth := c.MustGet("auth").(*bind.TransactOpts)
	contractAddress := c.MustGet("contract_address").(string)
	address := common.HexToAddress(contractAddress)
	instance, err := todo.NewTodo(address, client)
	if err != nil {
		log.Fatal(err)
	}

	todo, err := instance.CreateTodo(auth, input.Title, input.Description)
	_ = todo
	if err != nil {
		log.Fatal(err)
	}
	Id, _ := instance.NoOfTodos(nil)
	newTodo, _ := instance.GetTodo(nil, Id)
	c.JSON(http.StatusOK, gin.H{
		"data": newTodo,
	})
	return

}
