package main

import (
	"fmt"
	"log"
	"ungraded-challenge-3/config"
	"ungraded-challenge-3/handler"
)

func main() {
	router, server := config.SetupServer()
	db := &handler.NewInventoryHandler{DB: config.GetDatabase()}
	
	router.GET("/inventories", db.GetInventory)
	router.GET("/inventories/:id", db.GetInventoryById)
	router.POST("/inventories", db.AddNewInventory)
	router.PUT("/inventories/:id", db.UpdateInventory)
	router.DELETE("/inventories/:id", db.DeleteInventory)


	fmt.Println("Server running on port :8080")
	log.Fatal(server.ListenAndServe())
}
