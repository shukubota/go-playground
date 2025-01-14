package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	ShopName    string
	AccessToken string
}

type CustomerResponse struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	TotalSpent string `json:"total_spent"`
	Note       string `json:"note"`
}

type CustomerRequest struct {
	Customer Customer `json:"customer"`
}

type ShopifyClient struct {
	config Config
	client *http.Client
}

func NewShopifyClient(config Config) *ShopifyClient {
	return &ShopifyClient{
		config: config,
		client: &http.Client{},
	}
}

func (c *ShopifyClient) buildURL(path string) string {
	return fmt.Sprintf("https://%s.myshopify.com/admin/api/2025-01%s",
		c.config.ShopName, path)
}

func (c *ShopifyClient) sendRequest(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(method, c.buildURL(path), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Shopify-Access-Token", c.config.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req)
}

type Handler struct {
	client *ShopifyClient
}

func NewHandler(client *ShopifyClient) *Handler {
	return &Handler{client: client}
}

// 顧客一覧の取得
func (h *Handler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.client.sendRequest("GET", "/customers.json", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Shopify API error: %s", body), resp.StatusCode)
		return
	}

	var customers CustomerResponse
	if err := json.NewDecoder(resp.Body).Decode(&customers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// 顧客の詳細情報取得
func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")

	resp, err := h.client.sendRequest("GET", fmt.Sprintf("/customers/%s.json", customerID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Shopify API error: %s", body), resp.StatusCode)
		return
	}

	var customer CustomerRequest
	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// 顧客の作成
func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerReq CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customerReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.sendRequest("POST", "/customers.json", customerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Shopify API error: %s", body), resp.StatusCode)
		return
	}

	var customer CustomerRequest
	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

// 顧客情報の更新
func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")

	var customerReq CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customerReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.sendRequest("PUT", fmt.Sprintf("/customers/%s.json", customerID), customerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Shopify API error: %s", body), resp.StatusCode)
		return
	}

	var customer CustomerRequest
	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// 顧客の削除
func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")

	resp, err := h.client.sendRequest("DELETE", fmt.Sprintf("/customers/%s.json", customerID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Shopify API error: %s", body), resp.StatusCode)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	config := Config{
		ShopName:    os.Getenv("SHOPIFY_SHOP_NAME"),
		AccessToken: os.Getenv("SHOPIFY_ACCESS_TOKEN"),
	}

	shopifyClient := NewShopifyClient(config)
	handler := NewHandler(shopifyClient)

	r.Route("/customers", func(r chi.Router) {
		r.Get("/", handler.ListCustomers)
		r.Post("/", handler.CreateCustomer)
		r.Get("/{id}", handler.GetCustomer)
		r.Put("/{id}", handler.UpdateCustomer)
		r.Delete("/{id}", handler.DeleteCustomer)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
