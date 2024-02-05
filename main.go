package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"profile-svg/svg"
	"strings"
)

type Language struct {
	Name         string
	GradientFrom string
	GradientTo   string
	SVG          string
}

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
func (b *Body) WithRow(content ...Language) *Body {
        rows := 90 + (len(b.Rows) * 120)
	//var numberOfRowElements = len(content)
	var lowerXCounter = 10
	var elements []*Element
	for elementNumber, element := range content {
                id := fmt.Sprintf("%d_%d", len(b.Rows), elementNumber)
		elements = append(elements, &Element{Content: AddLanguage(lowerXCounter, rows, element.GradientFrom, element.GradientTo, id, element.SVG)})
		lowerXCounter += 100
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

type IncomingRequest struct {
	Languages  []string
	Frameworks []string
}

func GenerateSVG(w http.ResponseWriter, req *http.Request) {
	//var requestData IncomingRequest
	//queryParameters := req.URL.Query()
	//languages := queryParameters.Get("languages")
	//requestData.Languages = strings.Split(languages, ",")

	w.Header().Set("Content-Type", "image/svg+xml")
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
			Language{"Go", "#70D0ED", "#00CDBF", svg.Read("./icons/languages/Go.svg")},
			Language{"JavaScript", "#FFE693", "#E4A125", svg.Read("./icons/languages/JavaScript.svg")},
                        Language{"TypeScript", "#AFD3FC", "#2E79C7", svg.Read("./icons/languages/TypeScript.svg")},
                        Language{"Python", "#AFD3FC", "#2E79C7", svg.Read("./icons/languages/Python.svg")},
                        Language{"Visual Basic", "#AFD3FC", "#2E79C7", svg.Read("./icons/languages/VisualBasic.svg")},
                        Language{"HTML5", "#FDBAA2", "#F1652A", svg.Read("./icons/languages/HTML.svg")},
			Language{"CSS3", "#A7C1FD", "#2865F0", svg.Read("./icons/languages/CSS.svg")},
		).
                WithRow(
                        Language{"Svelte", "#FFB7A6", "#F83A01", svg.Read("./icons/frameworks/Svelte.svg")},
                        Language{"Preact", "#FFB7A6", "#F83A01", svg.Read("./icons/frameworks/Preact.svg")},
                        Language{"React", "#FFB7A6", "#F83A01", svg.Read("./icons/frameworks/React.svg")},
                        Language{"Ionic", "#FFB7A6", "#F83A01", svg.Read("./icons/frameworks/Ionic.svg")},


		)
	io.WriteString(w, CreateSVG(1000, 1000, body.String()))
}

func AddLanguage(x_offset int, y_offset int, gradient_from string, gradient_to string, id string, svg string) string {
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

func CreateSVG(width int, height int, contents string) string {
	return fmt.Sprintf(`
	<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">
	%s
	</svg>
	`, width, height, contents)
}
