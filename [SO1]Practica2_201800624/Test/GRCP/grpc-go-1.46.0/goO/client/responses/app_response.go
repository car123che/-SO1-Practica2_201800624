package responses

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	Status   int        `json:"status"`
	Message  string     `json:"message"`
	Data     *fiber.Map `json:"data"`
	Nombrevm string     `json:"nombreVM"`
	Endpoint string     `json:"endpoint"`
	Fecha    time.Time  `json:"fecha"`
}

type Process struct {
	Status   int        `json:"status"`
	Message  string     `json:"message"`
	Data     *fiber.Map `json:"data"`
	Nombrevm string     `json:"nombreVM"`
	Endpoint string     `json:"endpoint"`
	Fecha    time.Time  `json:"fecha"`
}

type Inicio struct{
	Status   int        `json:"status"`
	Message  string     `json:"message"`
	Bienvenido string    `json:"Bienvenido"`
 



}