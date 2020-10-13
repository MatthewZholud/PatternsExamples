package message

import "fmt"
// PlatformB - Definition of message platform B
type PlatformB struct {
	Name string
	// Add other fields for platform B
}
// SendMessage - Send a message
func (platform PlatformB) SendMessage(message string) {
	fmt.Printf("Sending message '%s' through platform B!", message)
}
