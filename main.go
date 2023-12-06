package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
        "profile-svg/svg"
)

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
	body := svg.Title("Skills", 44) +
                svg.Paragraph("Languages", 70) + 
                svg.Paragraph("Frameworks", 200) +
		Language(10, 90, "#70D0ED", "#00CDBF", "1", svg.Read("./icons/languages/Go.svg")) +
		Language(110, 90, "#FFE693", "#E4A125", "2", svg.Read("./icons/languages/JavaScript.svg")) +
		Language(210, 90, "#AFD3FC", "#2E79C7", "3", svg.Read("./icons/languages/TypeScript.svg")) +
                Language(310, 90, "#AFD3FC", "#2E79C7", "4", svg.Read("./icons/languages/Python.svg")) +
                Language(10, 210, "#AFD3FC", "#2E79C7", "5", svg.Read("./icons/frameworks/Svelte.svg"))

	footer := "</svg>"
	svg := header + body + footer
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
		%s
	</svg>
	`, x_offset, y_offset, id, gradient_from, gradient_to, id, id, svg)
}