package frontend

import "fmt"

func logger(event string) string {
	return fmt.Sprintf("INFO: %s", event)
}

func toastify(event string) string {
	return fmt.Sprintf("notified by toast ui: %s", event)
}
