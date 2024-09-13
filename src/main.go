package main

import (
	"flag"
	"main/database"
	"main/funcs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbFlag := flag.Bool("d", false, "use database for store data")
	flag.Parse()

	// Мапа для хранения сокращенных ссылок
	links := make(map[string]string)

	// Мапа-кэш для чтения из БД
	links_db := make(map[string]string)

	// POST эндпоинт для создания сокращенной ссылки
	app.Post("/shorten", func(c *fiber.Ctx) error {
		// Получение оригинальной ссылки из тела запроса
		originalLink := string(c.Body())

		// Генерация случайного ключа для сокращенной ссылки
		shortLink := funcs.GenerateRandomString(6)

		// Сохранение сокращенной ссылки в мапу
		if *dbFlag {
			database.AddSortLink(shortLink, originalLink)
		} else {
			links[shortLink] = originalLink
		}

		// Формирование URL сокращенной ссылки
		shortenedLink := c.BaseURL() + "/" + shortLink

		// Возвращение сокращенной ссылки в ответе
		return c.SendString(shortenedLink)
	})

	// GET эндпоинт для перенаправления на оригинальную ссылку
	app.Get("/:key", func(c *fiber.Ctx) error {
		var originalLink string
		var exists bool

		// Получение ключа сокращенной ссылки из URL
		shortLink := c.Params("key")

		if *dbFlag {
			originalLink, exists = links_db[shortLink]
			if !exists {
				var err error
				originalLink, err = database.GetOriginalLink(shortLink)
				if err != nil {
					return c.Status(fiber.StatusNotFound).SendString("Link not found or database credentials are invalid (database)")
				}
				links_db[shortLink] = originalLink
			}

		} else {
			originalLink, exists = links[shortLink]
			if !exists {
				return c.Status(fiber.StatusNotFound).SendString("Link not found (memory)")
			}
		}

		return c.SendString(originalLink)
	})

	// Запуск сервера
	app.Listen(":8080")
}
