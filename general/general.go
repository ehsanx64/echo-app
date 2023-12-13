package general

import (
	"fmt"
	"strings"
)

func getWelcomeMessage(name string) string {
	var out string = "User"

	if name != "" {
		if len(name) > 2 {
			out = strings.Title(name)
		} else {
			out = strings.ToUpper(name)
		}
	}

	return fmt.Sprintf("Welcome %s!!!\n", out)
}
