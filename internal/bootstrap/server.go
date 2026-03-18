package bootstrap

import (
	"context"
	"fmt"
	"my_project/docs"
	apphttp "my_project/internal/delivery/http"
	"my_project/internal/infrastructure/repository/postgres"
	"my_project/internal/infrastructure/repository/product"
	"my_project/internal/usecase"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	port int
}

// @BasePath /api/v1
func NewServer() (*http.Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	dbConn, err := newDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	const basePath = "/api/v1"
	router := gin.Default()

	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	group := router.Group(basePath)

	// init routes
	initProductsHandler(group, dbConn)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}

func newDatabase(addr string) (*pgxpool.Pool, error) {
	conn, err := postgres.NewPool(context.Background(), postgres.Config{
		DSN:             addr,
		MaxConns:        5,
		MinConns:        1,
		MaxConnLifetime: 0,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create repository pool: %w", err)
	}

	return conn, nil
}

func initProductsHandler(group *gin.RouterGroup, dbConn *pgxpool.Pool) {
	productRepo := product.NewProductRepository(dbConn)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := apphttp.NewProductHandler(productUseCase)
	apphttp.BindProductRoutes(group, productHandler)
}
