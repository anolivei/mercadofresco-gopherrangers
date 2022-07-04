package buyers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers/buyers"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/buyer/domain/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	handler "github.com/Gopher-Rangers/mercadofresco-gopherrangers/cmd/server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	URL = "/api/v1/buyers/"
)

type responseDataArray struct {
	Code int
	Data []domain.Buyer
}

type responseData struct {
	Code  int
	Data  domain.Buyer
	Error string
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

func createRequestTestIvalidToken(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "invalid_token")
	return req, httptest.NewRecorder()
}
func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		baseData := createBaseData()
		req, response := createRequestTest(http.MethodGet, URL, "")

		mockService.On("GetAll", context.Background()).Return(baseData, nil)
		buyerRouterGroup.GET("/", buyerHandler.GetAll)
		server.ServeHTTP(response, req)

		resp := responseDataArray{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, baseData, resp.Data)
	})
	t.Run("find_all_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		req, response := createRequestTestIvalidToken(http.MethodGet, URL, "")

		buyerRouterGroup.GET("/", buyerHandler.AuthToken, buyerHandler.GetAll)
		server.ServeHTTP(response, req)

		resp := responseDataArray{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
}

func createBaseData() []domain.Buyer {
	var buyers []domain.Buyer
	buyerOne := domain.Buyer{
		ID:           25735482,
		CardNumberId: "Card1",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyerTwo := domain.Buyer{
		ID:           25735582,
		CardNumberId: "Card2",
		FirstName:    "Victor",
		LastName:     "Beltramini",
	}
	buyers = append(buyers, buyerOne, buyerTwo)
	return buyers
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)
		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		req, response := createRequestTest(http.MethodDelete, URL+"1", "")

		mockService.On("Delete", context.Background(), 1).Return(nil)
		buyerRouterGroup.DELETE("/:id", buyerHandler.Delete)
		server.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("delete_non_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)
		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		req, response := createRequestTest(http.MethodDelete, URL+"1", "")

		mockService.On("Delete", context.Background(), 1).Return(fmt.Errorf("produto 1 não encontrado"))
		buyerRouterGroup.DELETE("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.Delete)
		server.ServeHTTP(response, req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("delete_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)
		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		req, response := createRequestTest(http.MethodDelete, URL+"non_number", "")

		buyerRouterGroup.DELETE("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.Delete)
		server.ServeHTTP(response, req)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("delete_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)
		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		req, response := createRequestTestIvalidToken(http.MethodDelete, URL+"1", "")

		buyerRouterGroup.DELETE("/:id", buyerHandler.AuthToken, buyerHandler.ValidateID, buyerHandler.Delete)
		server.ServeHTTP(response, req)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})
}

func TestStore(t *testing.T) {
	//t.Run("create_ok", func(t *testing.T) {
	//	mockService := mocks.NewService(t)
	//	buyerHandler := handler.NewBuyer(mockService)
	//
	//	server := gin.Default()
	//	buyerRouterGroup := server.Group(URL)
	//
	//	buyerData := createBaseData()
	//	expected := `{"id":25735482,"card_number_id":"Card1","first_name":"Victor","last_name":"Beltramini"}`
	//
	//	req, response := createRequestTest(http.MethodPost, URL, expected)
	//	mockService.On("Create", context.Background(), buyer.Buyer{
	//		ID:           25735482,
	//		CardNumberId: "Card1",
	//		FirstName:    "Victor",
	//		LastName:     "Beltramini",
	//	}).Return(buyerData[0], nil)
	//	buyerRouterGroup.POST("/", buyerHandler.Create)
	//	server.ServeHTTP(response, req)
	//
	//	resp := responseData{}
	//	json.Unmarshal(response.Body.Bytes(), &resp)
	//
	//	assert.Equal(t, http.StatusCreated, response.Code, resp.Code)
	//	assert.Equal(t, buyerData[0], resp.Data)
	//	assert.Equal(t, resp.Error, "")
	//})
	//t.Run("create_conflict", func(t *testing.T) {
	//	ctx := context.Background()
	//	mockService := mocks.NewService(t)
	//	buyerHandler := handler.NewBuyer(mockService)
	//
	//	server := gin.Default()
	//	buyerRouterGroup := server.Group(URL)
	//
	//	buyersData := createBaseData()
	//	expected := `{"id":25735482,"card_number_id":"Card1","first_name":"Victor","last_name":"Beltramini"}`
	//
	//	req, response := createRequestTest(http.MethodPost, URL, expected)
	//	mockService.On("Create", ctx, buyersData[0]).Return(buyer.Buyer{},
	//		fmt.Errorf("buyer with card_number_id %s already exists", buyersData[0].CardNumberId))
	//	buyerRouterGroup.POST("/", buyerHandler.AuthToken, buyerHandler.Create)
	//	server.ServeHTTP(response, req)
	//
	//	resp := responseData{}
	//	json.Unmarshal(response.Body.Bytes(), &resp)
	//
	//	assert.Equal(t, http.StatusNotFound, response.Code, resp.Code)
	//	assert.Equal(t, buyer.Buyer{}, resp.Data)
	//	assert.Equal(t, resp.Error, "buyer with card_number_id Card1 already exists")
	//})
	t.Run("create_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"Card1","first_name":"Victor Hugoo","last_name":"Beltramini"}`

		req, response := createRequestTestIvalidToken(http.MethodPost, URL, expected)
		buyerRouterGroup.POST("/", buyerHandler.AuthToken, buyerHandler.Create)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "invalid token")
	})
	t.Run("create_wrong_body", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"","first_name":"Victor Hugoo","last_name":"Beltramini"}`

		req, response := createRequestTest(http.MethodPost, URL, expected)
		buyerRouterGroup.POST("/", buyerHandler.AuthToken, buyerHandler.Create)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "Validation error")
	})
}

func TestGetById(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyersData := createBaseData()
		req, response := createRequestTest(http.MethodGet, URL+"1", "")

		mockService.On("GetById", context.Background(), 1).Return(buyersData[0], nil)
		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.GetBuyerById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, response.Code, resp.Code)
		assert.Equal(t, buyersData[0], resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := domain.Buyer{}
		req, response := createRequestTest(http.MethodGet, URL+"25735483", "")

		mockService.On("GetById", context.Background(), 25735483).Return(buyerData, fmt.Errorf("buyer with id %d not founded", 25735483))
		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.GetBuyerById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, response.Code, resp.Code)
		assert.Equal(t, buyerData, resp.Data)
		assert.Equal(t, resp.Error, "buyer with id 25735483 not founded")
	})
	t.Run("find_by_id_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := domain.Buyer{}
		req, response := createRequestTestIvalidToken(http.MethodGet, URL+"1", "")

		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.GetBuyerById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
		assert.Equal(t, buyerData, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
	t.Run("find_by_id_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := domain.Buyer{}
		req, response := createRequestTest(http.MethodGet, URL+"sdadas", "")

		buyerRouterGroup.GET("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.GetBuyerById)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, buyerData, resp.Data)
		assert.Equal(t, resp.Error, "Id need to be a valid integer")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := domain.Buyer{
			ID:           25735482,
			CardNumberId: "Card1231",
			FirstName:    "Victor Hugoo",
			LastName:     "Beltramini",
		}
		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"Victor Hugoo","last_name":"Beltramini"}`

		req, response := createRequestTest(http.MethodPatch, URL+"25735482", expected)
		mockService.On("Update", context.Background(), buyerData).Return(buyerData, nil)
		buyerRouterGroup.PATCH("/:id", buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, response.Code, resp.Code)
		assert.Equal(t, buyerData, resp.Data)
		assert.Equal(t, resp.Error, "")
	})
	t.Run("update_invalid_first_name", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"","last_name":"Beltramini"}`

		req, response := createRequestTest(http.MethodPut, URL+"25735482", expected)
		buyerRouterGroup.PUT("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "first_name is mandatory")
	})
	t.Run("update_invalid_last_name", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"Victor Hugo","last_name":""}`

		req, response := createRequestTest(http.MethodPut, URL+"25735482", expected)
		buyerRouterGroup.PUT("/:id", buyerHandler.ValidateID, buyerHandler.AuthToken, buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "last_name is mandatory")
	})
	t.Run("update_conflict", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		buyerData := domain.Buyer{
			ID:           25735482,
			CardNumberId: "Card1231",
			FirstName:    "Victor Hugoo",
			LastName:     "Beltramini",
		}
		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"Victor Hugoo","last_name":"Beltramini"}`
		req, response := createRequestTest(http.MethodPut, URL+"25735482", expected)
		mockService.On("Update", context.Background(), buyerData).Return(domain.Buyer{}, fmt.Errorf("buyer with card_number_id %s already exists", buyerData.CardNumberId))
		buyerRouterGroup.PUT("/:id", buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusConflict, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "buyer with card_number_id Card1231 already exists")
	})
	t.Run("update_invalid_token", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"Victor Hugoo","last_name":"Beltramini"}`
		req, response := createRequestTestIvalidToken(http.MethodPatch, URL+"1", expected)
		buyerRouterGroup.PATCH("/:id", buyerHandler.AuthToken, buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusUnauthorized, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, handler.ERROR_TOKEN)
	})
	t.Run("update_id_non_number", func(t *testing.T) {
		mockService := mocks.NewService(t)
		buyerHandler := buyers.NewBuyer(mockService)

		server := gin.Default()
		buyerRouterGroup := server.Group(URL)

		expected := `{"id":25735482,"card_number_id":"Card1231","first_name":"Victor Hugoo","last_name":"Beltramini"}`
		req, response := createRequestTest(http.MethodPatch, URL+"non_number", expected)
		buyerRouterGroup.PATCH("/:id", buyerHandler.ValidateID, buyerHandler.Update)
		server.ServeHTTP(response, req)

		resp := responseData{}
		json.Unmarshal(response.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, response.Code, resp.Code)
		assert.Equal(t, domain.Buyer{}, resp.Data)
		assert.Equal(t, resp.Error, "Id need to be a valid integer")
	})
}
