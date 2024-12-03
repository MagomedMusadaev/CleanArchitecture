package main

import (
	routhttp "CleanArchitecture/api/http"
	"CleanArchitecture/internal/handlers"
	"CleanArchitecture/internal/repositories"
	"CleanArchitecture/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Инициализируем репозиторий (мнимая база данных)
	inMemoryRepo := repositories.NewInMemoryUserRepository()

	// Инициализируем сервис
	userService := services.NewUserService(inMemoryRepo)

	// Инициализируем обработчики
	userHandler := handlers.NewUserHandler(userService)

	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Регистрируем маршруты
	routhttp.RegisterUserRoutes(mux, userHandler)

	// Запуск HTTP-сервера
	port := 8086
	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
