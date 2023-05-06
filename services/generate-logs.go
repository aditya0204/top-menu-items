package services

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"top-menu-items/models"
)

func generateRandomID(max int) int {
	return rand.Intn(max)
}
func GenerateLogs() {
	fileName := fmt.Sprintf("./logs/menu_logs_%s.json", time.Now().Format("02-01-2006"))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var logLines models.MenuLogs
	// Generate random data and write to the log file
	for i := 0; i < 100; i++ {
		eaterID := generateRandomID(1000)
		foodMenuID := generateRandomID(5)
		logLines.Lines = append(logLines.Lines, models.MenuLogLine{
			EaterId:    eaterID,
			FoodMenuId: foodMenuID,
		},
		)

	}

	jsonBytes, err := json.MarshalIndent(logLines, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(jsonBytes)

}
