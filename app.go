package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

// setup redis
func setupRedis() *redis.Client {

	// get redis args from env
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")
	redis_db := os.Getenv("REDIS_DB_NUMBER")
	redis_max_retries := os.Getenv("REDIS_MAX_RETRIES")

	// convert string to number
	db, _ := strconv.Atoi(redis_db)
	max_retries, _ := strconv.Atoi(redis_max_retries)

	client := redis.NewClient(&redis.Options{
		Addr:       redis_host + ":" + redis_port,
		Password:   redis_password,
		DB:         db,
		MaxRetries: max_retries,
	})

	// Ping the Redis server to ensure a successful connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Redis connected successfully:", pong)

	return client
}

// load environment variable
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file: ", err)
	}
}

func main() {
	app := fiber.New()

	// load env variables
	loadEnv()

	// setup redis
	client := setupRedis()

	// ping
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi, from golang fiber")
	})

	app.Get("/customer-segments", func(c *fiber.Ctx) error {
		msisdn := c.Query("msisdn")
		if msisdn == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "msisdn query parameter is required",
			})
		}

		// TODO: Perform any operations needed using the msisdn query parameter.
		log.Printf("msisdn received: %s\n", msisdn)

		// Create a JSON object to store the response.
		response := make(map[string]interface{})

		// Lookup the MSISDN in Redis sets.
		// TODO: Need to define All set process
		for i := 1; i <= 20; i++ {
			key := fmt.Sprintf("customer_segment_%d", i)
			// customer_segment check
			exists, err := client.SIsMember(context.Background(), key, msisdn).Result()
			if err != nil {
				log.Fatal(err.Error())
			}
			if exists {
				// get customer segment wise feature
				feature_key := fmt.Sprintf("customer_segment_%d:feature", i)
				val := client.HGetAll(context.Background(), feature_key).Val()
				response[key] = val
			} else {
				response[key] = 0
			}

		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msisdn": msisdn,
			"data":   response,
		})
	})

	// listen app
	server_port := os.Getenv("SERVER_PORT")
	app.Listen(":" + server_port)
}
