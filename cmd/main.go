package main

import (
	"log"
	"net/http"
	delivery "todo-api/internal/delivery/http"
	sql "todo-api/internal/repository/mysql"
	"todo-api/internal/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	// DB接続
	dsn := "root:password@tcp(localhost:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 依存関係初期化
	repo := sql.NewMySQLRepository(db)
	todoUseCase := usecase.NewTodoUseCase(repo)
	todoHandler := delivery.NewTodoHandler(todoUseCase)

	// Setup router
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Todo API"))
	})

	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
