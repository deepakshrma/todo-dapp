package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"

	todo "github.com/deepakshrma/todo-dapp/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

var (
	contextKeyClient = "ethClient"
)

// func doSomething(ctx context.Context) {
// 	client, _ := ctx.Value(contextKeyClient).(ClientContext)

// 	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
// 	balance, err := client.client.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balance)
// }

// func doAnother(ctx context.Context, printCh <-chan int) {
// ctx, cancelCtx := context.WithCancel(ctx)

// printCh := make(chan int)
// go doAnother(ctx, printCh)

// for num := 1; num <= 3; num++ {
// 	printCh <- num
// }

// cancelCtx()

// time.Sleep(100 * time.Millisecond)

// fmt.Printf("doSomething: finished\n")
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			if err := ctx.Err(); err != nil {
// 				fmt.Printf("doAnother err: %s\n", err)
// 			}
// 			fmt.Printf("doAnother: finished\n")
// 			return
// 		case num := <-printCh:
// 			fmt.Printf("doAnother: %d\n", num)
// 		}
// 	}
// }
func Logger(eth *ethclient.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(eth)
		c.Set(contextKeyClient, eth)
		c.Next()
	}
}

const ( // iota is reset to 0
	NewTodo     = iota // c0 == 0
	DoneTodo    = iota // c1 == 1
	DeletedTodo = iota // c2 == 2
)

type TodoItem struct {
	Id          big.Int `form:"Id"`
	Status      int     `form:"Status"`
	Title       string  `form:"Title"`
	Description string  `form:"Description"`
	CreatedBy   string  `form:"CreatedBy"`
}

func main() {
	err := godotenv.Load(".env")
	var envs map[string]string
	envs, err = godotenv.Read()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("STACK")
	_ = val
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.Use(Logger(client))
	v1.GET("/deploy", func(c *gin.Context) {
		client := c.MustGet(contextKeyClient).(*ethclient.Client)
		auth := getAuth(client, envs["PRIVATE_KEY"])

		address, _, instance, err := todo.DeployTodo(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		_ = instance
		file := fmt.Sprintf(`
		PRIVATE_KEY=%s
		CONTRACT_ADDRESS=%s
		`, envs["PRIVATE_KEY"], address.Hex())
		env, _ := godotenv.Unmarshal(file)
		godotenv.Write(env, "./.env")
		c.JSON(http.StatusOK, gin.H{
			"data": address.Hex(),
		})
	})
	v1.GET("/todos/:id", func(c *gin.Context) {
		client := c.MustGet(contextKeyClient).(*ethclient.Client)
		// todoId := c.Param("id")
		address := common.HexToAddress(envs["CONTRACT_ADDRESS"])
		instance, err := todo.NewTodo(address, client)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Param("id"))
		todoId := new(big.Int)
		id, _ := todoId.SetString(c.Param("id"), 10)

		todo, err := instance.GetTodo(nil, id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": todo,
		})
	})
	v1.POST("/todos/", func(c *gin.Context) {
		var newTodo TodoItem
		if c.ShouldBind(&newTodo) == nil {
			client := c.MustGet(contextKeyClient).(*ethclient.Client)

			address := common.HexToAddress(envs["CONTRACT_ADDRESS"])
			instance, err := todo.NewTodo(address, client)
			if err != nil {
				log.Fatal(err)
			}
			auth := getAuth(client, envs["PRIVATE_KEY"])

			fmt.Println(newTodo)
			title := newTodo.Title
			description := newTodo.Description
			todoRes, err := instance.CreateTodo(auth, title, description)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"data": todoRes,
			})

			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something fail",
		})

	})
	v1.GET("/todos/", func(c *gin.Context) {
		client := c.MustGet(contextKeyClient).(*ethclient.Client)

		address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
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

	})
	v1.GET("/noOfTodos", func(c *gin.Context) {
		client := c.MustGet(contextKeyClient).(*ethclient.Client)

		address := common.HexToAddress(envs["CONTRACT_ADDRESS"])
		instance, err := todo.NewTodo(address, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("contract is loaded")
		noOfTodos, err := instance.NoOfTodos(nil)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"NoOfTodos": noOfTodos,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAuth(client *ethclient.Client, privateKeyStr string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	return auth
}

func createKs() {
	ks := keystore.NewKeyStore("./keys", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	file := "./keys/UTC--2022-08-18T12-27-36.757848000Z--b3c69e270ba54f0709201c1e413542f6f6990786"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
func backup() {
	// client, err := ethclient.Dial("http://localhost:8545")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("we have a connection")
	// _ = client // we'll use this in the upcoming sections
	/*
		account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
		balance, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(balance) // 25893180161173005034
		blockNumber := big.NewInt(0)
		balance, err = client.BalanceAt(context.Background(), account, blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(balance) // 25729324269165216042
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

		fmt.Println(ethValue) // 25.729324269165216041

		pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
		fmt.Println(pendingBalance) // 25729324269165216042

		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af7

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])
		fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
	*/
	// createKs()
	// importKs()

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(header.Number.String()) // 5671744
	// blockNumber := big.NewInt(0)
	// for i := 0; i < 6; i++ {
	// 	blockNumber.Add(blockNumber, big.NewInt(1))
	// 	block, err := client.BlockByNumber(context.Background(), blockNumber)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(block.Number().Uint64())     // 5671744
	// 	fmt.Println(block.Time())                // 1527211625
	// 	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	// 	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// 	fmt.Println(len(block.Transactions()))   // 144

	// 	count, err := client.TransactionCount(context.Background(), block.Hash())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("b: %v c: %v", block.Number().Uint64(), count) // 144

	// 	// for _, tx := range block.Transactions() {
	// 	// 	fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	// 	// 	fmt.Println(tx.Value().String())    // 10000000000000000
	// 	// 	fmt.Println(tx.Gas())               // 105000
	// 	// 	fmt.Println(tx.GasPrice().Uint64()) // 102000000000
	// 	// 	fmt.Println(tx.Nonce())             // 110644
	// 	// 	fmt.Println(tx.Data())              // []
	// 	// 	fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	// 	// }
	// }

	/*
		Send Transctions
	*/
	/*
		privateKey, err := crypto.HexToECDSA("")
		if err != nil {
			log.Fatal(err)
		}
		_ = privateKey
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		value := big.NewInt(1000000000000000000) // in wei (1 eth)
		gasLimit := uint64(21000)                // in units
		// gasPrice := big.NewInt(30000000000)      // in wei (30 gwei)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		fmt.Printf("Price: %v", gasPrice)
		if err != nil {
			log.Fatal(err)
		}

		toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")

		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
	*/

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	println(fromAddress.Hash().Hex())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	tokenAddress := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
