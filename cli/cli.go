package cli

import (
	"fmt"
)

// beautifully print
func Brint(text string) {
	fmt.Printf("🟢 %s\n", text)
}

func BrintErr(text string) {
	fmt.Printf("🔴 %s\n", text)
}
