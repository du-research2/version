package version

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Demo for Tzahi z, %v. Welcome Back Again!", name)
    return message
}

