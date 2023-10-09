package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.gom/Shimizu1111/go-layered-architecture/presentation"
	"github.gom/Shimizu1111/go-layered-architecture/repository"
	"github.gom/Shimizu1111/go-layered-architecture/usecase"
)

func main() {
	dsn := "root:password@tcp(localhost:3321)/user-test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	userRepo := &repository.UserRepositoryImpl{DB: db}

	userUsecase := &usecase.UserUsecase{Repo: userRepo}

	// ルーティングの設定
	http.HandleFunc("/user/", presentation.GetUserHandler(userUsecase))

	// サーバーの起動
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		return
	}
}
