package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"profile-svg/svg"
	"strings"
)

type Body struct {
	Title      string
	Paragraphs []string
	Rows       []*ElementRow
}

func NewBody() *Body {
	return &Body{}
}
func (b *Body) WithTitle(title string) *Body {
	b.Title = title
	return b
}
func (b *Body) WithParagraph(paragraph string) *Body {
	b.Paragraphs = append(b.Paragraphs, paragraph)
	return b
}
func (b *Body) WithRow(content ...string) *Body {
	var elements []*Element
	for _, f := range content {
		elements = append(elements, &Element{Content: f})
	}
	row := &ElementRow{Elements: elements}
	b.Rows = append(b.Rows, row)
	return b
}
func (b *Body) String() string {
	var body []string
	// First deal with the title
	if b.Title != "" {
		body = append(body, b.Title)
	}
	// Then deal with the paragraphs
	body = append(body, b.Paragraphs...)
	// Finally deal with the Rows of Elements
	for _, row := range b.Rows {
		var rows []string
		for _, element := range row.Elements {
			rows = append(rows, element.Content)
		}
		body = append(body, rows...)
	}
	// Turn the array of strings into a singular string
	return strings.Join(body, "")
}

type ElementRow struct {
	Elements []*Element
}

type Element struct {
	Content string
}

func NewElement(content string) *Element {
	return &Element{
		Content: content,
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(GenerateSVG))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Println("Started Webserver on port :8081")
}

func GenerateSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	header := `<svg xmlns="http://www.w3.org/2000/svg" width="1000" height="1000">`
	body := NewBody().
		WithTitle(
			svg.Title("Skills", 44),
		).
		WithParagraph(
			svg.Paragraph("Languages", 70),
		).
		WithParagraph(
			svg.Paragraph("Frameworks", 200),
		).
		WithRow(
			Language(10, 90, "#70D0ED", "#00CDBF", "1", svg.Read("./icons/languages/Go.svg")),
			Language(110, 90, "#FFE693", "#E4A125", "2", svg.Read("./icons/languages/JavaScript.svg")),
			Language(210, 90, "#AFD3FC", "#2E79C7", "3", svg.Read("./icons/languages/TypeScript.svg")),
			Language(310, 90, "#AFD3FC", "#2E79C7", "4", svg.Read("./icons/languages/Python.svg")),
		)

	footer := "</svg>"
	svg := header + body.String() + footer
	io.WriteString(w, svg)
}

func Language(x_offset int16, y_offset int16, gradient_from string, gradient_to string, id string, svg string) string {
	return fmt.Sprintf(`
	<svg width="80" height="80" x="%d" y="%d">
		<defs>
			<linearGradient id="linear-gradient%s" gradientUnits="objectBoundingBox">
				<stop offset="0" stop-color="%s" />
				<stop offset="1" stop-color="%s" />
			</linearGradient>
		</defs>
		<rect id="gradient-box%s" width="80" height="80" rx="20" fill="url(#linear-gradient%s)" />
                <g transform="translate(20, 20)">
		%s
                </g>
	</svg>
	`, x_offset, y_offset, id, gradient_from, gradient_to, id, id, svg)
}
