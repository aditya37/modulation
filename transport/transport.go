package transport

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func Print(value string) {
	u := uuid.New()
	key := fmt.Sprintf("%s:%s", u.String(), value)
	log.Println(key)
	log.Println(time.Now().String())
}
