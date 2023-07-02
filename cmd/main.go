package main

import (
	"fmt"

	"github.com/Jahankohan/mpc_wallet/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()


	// Initialize Ethereum client
	// client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }

	// // Initialize key manager
	// keyManager := &key_manager.KeyManager{}

	// Initialize transaction builder
	// transactionBuilder := transaction.NewTransactionBuilder(client, keyManager)

	// // Initialize the transaction handler
	// transactionHandler := handlers.NewTransactionHandler(transactionBuilder)

	// API routes
	api := router.Group("/api")
	{
		// User routes
		api.POST("/register", handlers.RegisterUser) // assuming RegisterUser is implemented elsewhere
		api.POST("/authenticate", handlers.AuthenticateUser) // assuming AuthenticateUser is implemented elsewhere
		
		// Transaction route
		// api.POST("/transaction", transactionHandler.HandleTransaction)
	}

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)
}
