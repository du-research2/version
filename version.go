package version

import (
    "fmt"
    "github.com/teris-io/shortid"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Demo for Tzahi zzz, %v. Welcome Back Again! - with pr", name)
    fmt.Printf(shortid.Generate())
    return message
}


