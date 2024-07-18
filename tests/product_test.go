package main

import (
	"ecommerce-backend/router"
	"net/http"
	"net/http/httptest"
	"testing"

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

    assert.Equal(t, http.StatusOK, w.Code)
    assert.JSONEq(t, `{"message": "Get all products"}`, w.Body.String())
}

// func TestGetProduct(t *testing.T) {
//     r := router.SetupRouter()
//     w := httptest.NewRecorder()
//     req, _ := http.NewRequest("GET", "/products/1", nil)
//     r.ServeHTTP(w, req)
	
//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.JSONEq(t, `{"message": "Get a single product"}`, w.Body.String())
// }

// func TestCreateProduct(t *testing.T) {
//     r := router.SetupRouter()
//     w := httptest.NewRecorder()
//     req, _ := http.NewRequest("POST", "/products", nil)
//     r.ServeHTTP(w, req)

//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.JSONEq(t, `{"message": "Create a new product"}`, w.Body.String())
// }

// func TestUpdateProduct(t *testing.T) {
//     r := router.SetupRouter()
//     w := httptest.NewRecorder()
//     req, _ := http.NewRequest("PUT", "/products/1", nil)
//     r.ServeHTTP(w, req)
// 	fmt.Printf("%+v\n", w)
//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.JSONEq(t, `{"message": "Update a product"}`, w.Body.String())
// }

func TestDeleteProduct(t *testing.T) {
    r := router.SetupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/products/1", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.JSONEq(t, `{"message": "Product deleted"}`, w.Body.String())
}