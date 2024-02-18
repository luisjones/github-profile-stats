package svg

import (
    "fmt"
    "strings"
)

type TextCharacteristics struct {
	ID         string
	Colour      string
	YOffset    int
	FontSize   int8
	FontFamily string
	FontWeight int16
	Text        string
}


func Title(text string, yoffset int) string {
    return Text(TextCharacteristics{strings.ToLower(text), "#000000", yoffset, 42, "Roboto-Medium, Roboto, sans-serif", 500, text})
}

func Paragraph(text string, yoffset int) string {
    return Text(TextCharacteristics{strings.ToLower(text), "#bfbfbf", yoffset, 24, "Roboto-Regular, Roboto, sans-serif", 500, text})
}


func Text(args TextCharacteristics) string {
	return fmt.Sprintf(`<text id="%s" fill="%s" transform="translate(10 %d)" font-size="%d" font-family="%s" font-weight="%d"><tspan x="0" y="0">%s</tspan></text>`, args.ID, args.Colour, args.YOffset, args.FontSize, args.FontFamily, args.FontWeight, args.Text)
}

// This function is used to center the text within an SVG
// Could do with optimising this function to produce a more accurate figure
func CenterText(args TextCharacteristics, word string) float64 {
        var xValue float64
        wordLength := len(word)
        if wordLength > 4 {
            xValue = 10 + (75 - float64(wordLength) * 7.5) / 2.5
        } else {
            xValue = (75 - float64(wordLength) * 7.5) / 2.5
        }
        return xValue
}
