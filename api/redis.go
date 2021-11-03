package api

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func SetRedisValue(c *fiber.Ctx) error {

	key := c.Query("key")
	value := c.Query("value")

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	return c.SendStatus(200)
}

func GetRedisValue(c *fiber.Ctx) error {

	key := c.Query("key")

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return c.SendString("Value is: " + val)
}
