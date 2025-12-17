package product

import (
	"crobe-ecommerce/app/backend/pkg/database"
	"crobe-ecommerce/app/backend/pkg/entities"
	"testing"

	"github.com/joho/godotenv"
)

var productTest = entities.Product{
	Name:        "Test Product",
	Description: "This is a test product",
}

var productID string

func TestProduct_New(t *testing.T) {
	t.Log("TestProduct_New started ...")
	godotenv.Load("../../../../.env")
	db, client, err := database.DatabaseConnection()
	if err != nil {
		t.Fatal("Database Connection Error: ", err)
	}
	defer func() {
		if err = client.Disconnect(t.Context()); err != nil {
			t.Fatal(err)
		}
	}()

	productColl := db.Collection("products")
	productRepo := NewRepo(productColl)
	productServ := NewService(productRepo)

	newProduct, err := productServ.InsertProduct(&productTest)
	if err != nil {
		t.Fatal("Error inserting new product: ", err)
	}

	productID = newProduct.ID.Hex()

	t.Log("[Database Test] Sucessfully inserted new product with ID: ", productID)
}
func TestProduct_Get(t *testing.T) {
	t.Log("TestProduct_Get started...")

	db, client, err := database.DatabaseConnection()
	if err != nil {
		t.Fatal("Database Connection Error: ", err)
	}
	defer func() {
		if err = client.Disconnect(t.Context()); err != nil {
			t.Fatal(err)
		}
	}()

	productColl := db.Collection("products")
	productRepo := NewRepo(productColl)
	productServ := NewService(productRepo)

	_, err = productServ.GetProduct(productID)
	if err != nil {
		t.Fatal("Error getting product: ", err)
	}
	t.Log("[Database Test] Sucessfully fetched product!")
	godotenv.Load()

}
func TestProduct_Update(t *testing.T) {
	t.Log("TestProduct_Update in development...")
}
func TestProduct_Delete(t *testing.T) {
	t.Log("TestProduct_Delete started...")
	db, client, err := database.DatabaseConnection()
	if err != nil {
		t.Fatal("Database Connection Error: ", err)
	}
	defer func() {
		if err = client.Disconnect(t.Context()); err != nil {
			t.Fatal(err)
		}
	}()

	productColl := db.Collection("products")
	productRepo := NewRepo(productColl)
	productServ := NewService(productRepo)

	err = productServ.DeleteProduct(&entities.DeleteRequest{ID: productID})
	if err != nil {
		t.Fatal("Error deleting product: ", err)
	}
	t.Log("[Database Test] Sucessfully deleted product!")
}
