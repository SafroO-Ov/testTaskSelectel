package rulesTest

import (
	"log"
	"log/slog"
)

func testEnglish() {
	// корректные — только английский

	log.Println("starting server")          // ok
	slog.Println("failed to connect to db") // ok
	log.Println("user authenticated")       // ok

	// некорректные — есть кириллица
	log.Println("запуск сервера")     // want "only English letters"
	log.Println("ошибка подключения") // want "only English letters"
	log.Println("starting сервер")    // want "only English letters"
	log.Println("user пароль reset")  // want "only English letters"
}
