package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"example/database/postgres"
	"example/database/repository"
	"example/dataloaders"
	"example/domain"
	"example/generated"
	"example/models"
)

const defaultPort = "8080"

func graphqlHandler(DB *gorm.DB) gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{
		Domain: domain.Domain{
			TaskRepository:     repository.TaskRepository{DB: DB},
			EmployeeRepository: repository.EmployeeRepository{DB: DB},
			M2MRepository:      repository.M2MRepository{DB: DB},
			ComapnyRepository:  repository.ComapnyRepository{DB: DB},
		},
	}}))
	return func(ctx *gin.Context) {
		handlerFunc := dataloaders.DataLoaderMiddleware(DB, srv)
		handlerFunc.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playGroundHandler() gin.HandlerFunc {
	handlerFunc := playground.Handler("GraphQL playground", "/query")
	return func(ctx *gin.Context) {
		handlerFunc.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("unable to load .env file")
	}
}

func main() {
	dropCreate := flag.Bool("db_drop_create", false,
		"Specifies if database needs to drop  table before migration while starting application.")
	dbUp := flag.Bool("db_up", false,
		"Specifies if database needs migrations while starting application.")
	dbDown := flag.Bool("db_down", false,
		"Specifies if database needs to drop while starting application.")
	flag.Parse()
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err := postgres.New(dsn)
	if err != nil {
		panic(err)
	}
	if *dbDown || *dropCreate {
		log.Println("Droping All Tables")
		err = DB.Migrator().DropTable(&models.Company{}, &models.Employee{}, &models.Task{})
		if err != nil {
			panic(err)
		}
	}
	if *dbUp || *dropCreate {
		log.Println("Auto Migrating All Tables")
		err = DB.AutoMigrate(&models.Company{}, &models.Employee{}, &models.Task{})
		if err != nil {
			panic(err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	models.PageLimit, _ = strconv.Atoi(os.Getenv("PAGE_LIMIT"))
	models.OrderBy = os.Getenv("ORDER_BY")

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:" + port},
		AllowCredentials: true,
	}))
	fmt.Println("[GIN-debug] Blog Written on MEDIUM")
	engine.POST("/query", graphqlHandler(DB))
	engine.GET("/", playGroundHandler())
	go fmt.Printf("[GIN-debug] Connect to http://localhost:%s/ for GraphQL playground\n", port)
	if engine.Run() != nil {
		panic("unable to run gin engine")
	}
}
