package svg

import (
    "fmt"
    "strings"
)

type TextCharacteristics struct {
	ID         string
	Colour      string
	YOffset    int16
	FontSize   int8
	FontFamily string
	FontWeight int16
	Text        string
}


func Title(text string, yoffset int16) string {
    return Text(TextCharacteristics{strings.ToLower(text), "#000000", yoffset, 42, "Roboto-Medium, Roboto, sans-serif", 500, text})
}

func Paragraph(text string, yoffset int16) string {
    return Text(TextCharacteristics{strings.ToLower(text), "#bfbfbf", yoffset, 24, "Roboto-Regular, Roboto, sans-serif", 500, text})
}


func Text(args TextCharacteristics) string {
	return fmt.Sprintf(`<text id="%s" fill="%s" transform="translate(10 %d)" font-size="%d" font-family="%s" font-weight="%d"><tspan x="0" y="0">%s</tspan></text>`, args.ID, args.Colour, args.YOffset, args.FontSize, args.FontFamily, args.FontWeight, args.Text)
}
