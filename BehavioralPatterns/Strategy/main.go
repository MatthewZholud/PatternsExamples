//Стратегия — это поведенческий паттерн, выносит набор
//алгоритмов в собственные классы и делает их взаимозаменимыми.
// пример с навигатором

package main
import (
	"github.com/PatternExamples/BehavioralPatterns/Strategy/message"
	"log"
)
func main() {
	//messagePlatform := "A"
	messagePlatform := "B"
	//messagePlatform := "C"
	platform, exists := message.MessagePlatforms[messagePlatform]
	if !exists {
		log.Fatalf("Platform %s does not exist!", messagePlatform)
	}
	platform.SendMessage("I am the message!")
}