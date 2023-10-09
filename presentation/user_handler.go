package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.gom/Shimizu1111/go-layered-architecture/usecase"
)

func GetUserHandler(usecase *usecase.UserUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/user/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		user, err := usecase.GetUserByID(id)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Failed to encode user", http.StatusInternalServerError)
			return
		}
	}
}
