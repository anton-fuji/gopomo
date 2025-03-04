package timer

import (
	"fmt"
	"time"
)

func Start(duration time.Duration) {
	fmt.Println("🧑🏼‍💻 Let's Started!!")
	for i := int(duration.Seconds()); i > 0; i-- {
		fmt.Printf("\nRemaining time")
	}
}
