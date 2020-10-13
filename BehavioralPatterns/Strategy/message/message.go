package message

// MessagePlatform - Behavior for each platform
type MessagePlatform interface {
	SendMessage(message string)
}

// MessagePlatforms - The available platforms registry
var MessagePlatforms = map[string] MessagePlatform{
	"A": PlatformA{Name: "Platform A"},
	"B": PlatformB{Name: "Platform B"},
}
