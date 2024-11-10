package quotes

import (
	"github.com/hackebrot/turtle"
)

func GetEmoji(name string) string {
	emoji, ok := turtle.Emojis[name]
	if !ok {
		return "Error!"
	}
	return emoji.Char
}

// no electricity to charge laptop, worked on phone and read more on web development with GO
