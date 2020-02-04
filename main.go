package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

type Payload struct {
	Token   string `json:"token"`
	HookUrl string `json:"response_url"`
	Actions []struct {
		Id string `json:"action_id"`
	} `json:"actions"`
}

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		var payload Payload
		formData := []byte(c.FormValue("payload"))

		if err := json.Unmarshal(formData, &payload); err != nil {
			return err
		}

		fmt.Println(payload.Token)
		fmt.Println(payload.HookUrl)
		fmt.Println(payload.Actions[0].Id)

		return c.JSON(200, echo.Map{})
	})

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
	}

	fmt.Println("hello")

	e.Logger.Fatal(e.Start(":" + port))
}
