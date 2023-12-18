package main

import (
	"log"
	"main/AddressBook/controllers/stdhttp"
	"main/AddressBook/gate/psg"
	"net/http"
)

func main() {

	db, err := psg.NewPsg("localhost:5432", "postgres", "123")
	if err != nil {
		log.Fatal("Error creating database connection:", err)
	}

	controller := stdhttp.NewController("localhost:8080", db)

	http.HandleFunc("/record/add", controller.RecordAdd)
	http.HandleFunc("/records/get", controller.RecordsGet)
	http.HandleFunc("/record/update", controller.RecordUpdate)
	http.HandleFunc("/record/delete", controller.RecordDeleteByPhone)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// // main.go
// package main

// import (
// 	"fmt"
// 	"myModules/AddressBook/controllers/stdhttp"
// 	"myModules/AddressBook/gate/psg"
// 	"net/http"
// )

// func main() {
// 	psgAddr := "postgresql://127.0.0.1:5432"
// 	psgLogin := "postgres"
// 	psgPassword := "123"

// 	db, err := psg.NewPsg(psgAddr, psgLogin, psgPassword)
// 	if err != nil {
// 		fmt.Println("Failed to connect to PostgreSQL:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// Порт, на котором будет запущен HTTP-сервер
// 	serverAddr := ":8080"

// 	// Создаем экземпляр HTTP-сервера с контроллером stdhttp.Controller
// 	controller := stdhttp.NewController(serverAddr, db)
// 	http.HandleFunc("/record/add", controller.RecordAdd)
// 	http.HandleFunc("/records/get", controller.RecordsGet)
// 	http.HandleFunc("/record/update", controller.RecordUpdate)
// 	http.HandleFunc("/record/deleteByPhone", controller.RecordDeleteByPhone)

// 	// Запускаем HTTP-сервер
// 	fmt.Printf("Server is listening on %s...\n", serverAddr)
// 	err = http.ListenAndServe(serverAddr, nil)
// 	if err != nil {
// 		fmt.Println("Failed to start server:", err)
// 	}
// }
