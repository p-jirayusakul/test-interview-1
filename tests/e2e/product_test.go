package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"my_project/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductE2E_TableDriven(t *testing.T) {
	if os.Getenv("DATABASE_URL") == "" {
		t.Setenv("DATABASE_URL", "postgresql://postgres:1234@localhost:5432/stock_db")
	}
	if os.Getenv("PORT") == "" {
		t.Setenv("PORT", "8080")
	}

	router := setupRouter(t)

	tests := []struct {
		name     string
		method   string
		path     string
		body     any
		wantCode int
		wantBody string
	}{
		{
			name:     "invalid id on get",
			method:   http.MethodGet,
			path:     "/api/v1/product/invalid-uuid",
			wantCode: http.StatusBadRequest,
			wantBody: `{"successful":false,"errorCode":"INVALID_INPUT"}`,
		},
		{
			name:     "not found on get",
			method:   http.MethodGet,
			path:     "/api/v1/product/019cfeaf-43bc-72d4-ae1b-d4667a111111",
			wantCode: http.StatusNotFound,
			wantBody: `{"successful":false,"errorCode":"NOT_FOUND"}`,
		},
		{
			name:   "create product",
			method: http.MethodPost,
			path:   "/api/v1/product/",
			body: map[string]any{
				"name":        "Apple",
				"description": "A sweet fruit",
				"salePrice":   1.99,
				"price":       2.99,
			},
			wantCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			var reqBody []byte
			var err error
			if tt.body != nil {
				reqBody, err = json.Marshal(tt.body)
				require.NoError(t, err)
			}

			req, err := http.NewRequest(tt.method, tt.path, bytes.NewReader(reqBody))
			require.NoError(t, err)

			if tt.body != nil {
				req.Header.Set("Content-Type", "application/json")
			}

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, w.Body.String())
			}
		})
	}
}

func setupRouter(t *testing.T) *gin.Engine {
	t.Helper()

	server, err := bootstrap.NewServer()
	require.NoError(t, err)

	r, ok := server.Handler.(*gin.Engine)
	require.True(t, ok, "server.Handler must be *gin.Engine")

	return r
}
