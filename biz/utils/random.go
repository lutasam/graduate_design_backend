package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateActiveCode generate an active code with 6 width in random
func GenerateActiveCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
