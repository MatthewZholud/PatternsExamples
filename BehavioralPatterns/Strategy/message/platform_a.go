package message

import "fmt"
// PlatformA - Definition of message platform A
type PlatformA struct {
	Name string
	// Add other fields for platform A
}
// SendMessage - Send a message
func (platform PlatformA) SendMessage(message string) {
	fmt.Printf("Sending message '%s' through platform A!", message)
}