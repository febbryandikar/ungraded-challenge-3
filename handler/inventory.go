package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ungraded-challenge-3/entity"

	"github.com/julienschmidt/httprouter"
)

type NewInventoryHandler struct {
	*sql.DB
}

func (h *NewInventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var inventories []entity.Inventory

	rows, err := h.Query("SELECT id, item_name, item_code, stock, description, status FROM inventory")
	if err!= nil {
        log.Fatal(err)
    }
	defer rows.Close()

	for rows.Next() {
		var inventory entity.Inventory
        err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.ItemCode, &inventory.Stock, &inventory.Description, &inventory.Status)
        if err!= nil {
            log.Fatal(err)
        }
        inventories = append(inventories, inventory)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status: "success",
        Message: "Inventories retrieved successfully",
        Data: inventories,
	})
}

func (h *NewInventoryHandler) GetInventoryById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var inventory entity.Inventory
	paramsId := p.ByName("id")

	query := "SELECT id, item_name, item_code, stock, description, status FROM inventory WHERE id = ?"
	if err := h.QueryRow(query, paramsId).Scan(&inventory.ID, &inventory.Name, &inventory.ItemCode, &inventory.Stock, &inventory.Description, &inventory.Status); err!= nil {
		log.Fatal(err)
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status: "success",
        Message: "Inventory retrieved successfully",
        Data: inventory,
	})
}

func (h *NewInventoryHandler) AddNewInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	newInventory := entity.Inventory{
		Name: "Captain Marvel's Suit",
    	ItemCode: "CMSU-001",
    	Stock: 5,
    	Description: "Suit worn by Captain Marvel, providing enhanced durability and flight capability",
    	Status: "active",
	}

	query := "INSERT INTO inventory (item_name, item_code, stock, description, status) VALUES (?, ?, ?, ?, ?)"
	_, err := h.Exec(query, newInventory.Name, newInventory.ItemCode, newInventory.Stock, newInventory.Description, newInventory.Status)
	if err!= nil {
		log.Fatal(err)
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status: "success",
        Message: "New inventory added successfully",
	})
}

func (h *NewInventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var inventory entity.Inventory
	paramsId := p.ByName("id")

	query := "SELECT id, item_name, item_code, stock, description, status FROM inventory WHERE id = ?"
	if err := h.QueryRow(query, paramsId).Scan(&inventory.ID, &inventory.Name, &inventory.ItemCode, &inventory.Stock, &inventory.Description, &inventory.Status); err!= nil {
		log.Fatal(err)
	}

	newInventory := entity.Inventory{
		Name: "Captain Marvel's Suit",
    	ItemCode: "CMSU-001",
    	Stock: 3,
    	Description: "Suit worn by Captain Marvel, providing enhanced durability and flight capability",
    	Status: "active",
	}

	query = "UPDATE inventory SET item_name = ?, item_code = ?, stock = ?, description = ?, status = ? WHERE id =?"
	_, err := h.Exec(query, newInventory.Name, newInventory.ItemCode, newInventory.Stock, newInventory.Description, newInventory.Status, inventory.ID)
	if err!= nil {
		log.Fatal(err)
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status: "success",
        Message: "Inventory updated successfully",
	})
}

func (h *NewInventoryHandler) DeleteInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paramsId := p.ByName("id")

	query := "DELETE FROM inventory WHERE id = ?"
	_, err := h.Exec(query, paramsId)
	if err!= nil {
		log.Fatal(err)
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status: "success",
        Message: "Inventory deleted successfully",
	})
}