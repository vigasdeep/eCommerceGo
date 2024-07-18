package main

import (
	"ecommerce-backend/router"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	// "github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	// Update with your module name"
)

// type App struct {
// 	Router *gin.Engine
// }
func TestGetProducts(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/products", nil)
    r.ServeHTTP(w, req)
    // fmt.Printf("Response: %s\n", w.Body.String()) // Log the response body

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProduct(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/products/1", nil)
    r.ServeHTTP(w, req)
	
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Regexp(t, `{"id":[0-9]+,"name":"Product [A-Za-z0-9]+","price":[0-9]+}`, w.Body.String())
}

func TestCreateProduct(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
    productJSON := `{"name": "Random Product", "price": 123}`
    req, _ := http.NewRequest("POST", "/products", strings.NewReader(productJSON))
    req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    // assert.Regexp(t, `{"id": [0-9]+, "name": "Random Product", "price": 123}`, w.Body.String())
}

func TestUpdateProduct(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
	productJSON := `{"name": "Random Product", "price": 123}`

    req, _ := http.NewRequest("PUT", "/products/1", strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")
    r.ServeHTTP(w, req)

	// fmt.Printf("%+v\n", w)
    assert.Equal(t, http.StatusOK, w.Code)
    // assert.JSONEq(t, `{"message": "Update a product"}`, w.Body.String())
}

func TestDeleteProduct(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/products/2", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.JSONEq(t, `{"message": "Product deleted"}`, w.Body.String())
}