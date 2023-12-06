package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		Language(10, 90, "#70D0ED", "#00CDBF", "1", ReadSvg("./languages/Go.svg")) +
		Language(110, 90, "#FFE693", "#E4A125", "2", ReadSvg("./languages/JavaScript.svg")) +
		Language(210, 90, "#AFD3FC", "#2E79C7", "3", ReadSvg("./languages/TypeScript.svg")) +
                Language(310, 90, "#AFD3FC", "#2E79C7", "3", ReadSvg("./languages/Python.svg")) +
                Language(10, 210, "#AFD3FC", "#2E79C7", "3", ReadSvg("./languages/Svelte.svg"))

	footer := "</svg>"
	svg := header + body + footer
	io.WriteString(w, svg)
}

func ReadSvg(file_name string) string {
	b, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
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

/*

    <svg width="80" height="80" x="210" y="90">
        <defs>
            <linearGradient id="linear-gradient3" gradientUnits="objectBoundingBox">
                <stop offset="0" stop-color="#AFD3FC" />
                <stop offset="1" stop-color="#2E79C7" />
            </linearGradient>
        </defs>
        <rect id="gradient-box3" width="80" height="80" rx="20" fill="url(#linear-gradient3)" />


        <g transform="translate(20, 20)">


            <svg fill="white" width="40" height="40" role="img" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <title>TypeScript</title>
                <path
                    d="M1.125 0C.502 0 0 .502 0 1.125v21.75C0 23.498.502 24 1.125 24h21.75c.623 0 1.125-.502 1.125-1.125V1.125C24 .502 23.498 0 22.875 0zm17.363 9.75c.612 0 1.154.037 1.627.111a6.38 6.38 0 0 1 1.306.34v2.458a3.95 3.95 0 0 0-.643-.361 5.093 5.093 0 0 0-.717-.26 5.453 5.453 0 0 0-1.426-.2c-.3 0-.573.028-.819.086a2.1 2.1 0 0 0-.623.242c-.17.104-.3.229-.393.374a.888.888 0 0 0-.14.49c0 .196.053.373.156.529.104.156.252.304.443.444s.423.276.696.41c.273.135.582.274.926.416.47.197.892.407 1.266.628.374.222.695.473.963.753.268.279.472.598.614.957.142.359.214.776.214 1.253 0 .657-.125 1.21-.373 1.656a3.033 3.033 0 0 1-1.012 1.085 4.38 4.38 0 0 1-1.487.596c-.566.12-1.163.18-1.79.18a9.916 9.916 0 0 1-1.84-.164 5.544 5.544 0 0 1-1.512-.493v-2.63a5.033 5.033 0 0 0 3.237 1.2c.333 0 .624-.03.872-.09.249-.06.456-.144.623-.25.166-.108.29-.234.373-.38a1.023 1.023 0 0 0-.074-1.089 2.12 2.12 0 0 0-.537-.5 5.597 5.597 0 0 0-.807-.444 27.72 27.72 0 0 0-1.007-.436c-.918-.383-1.602-.852-2.053-1.405-.45-.553-.676-1.222-.676-2.005 0-.614.123-1.141.369-1.582.246-.441.58-.804 1.004-1.089a4.494 4.494 0 0 1 1.47-.629 7.536 7.536 0 0 1 1.77-.201zm-15.113.188h9.563v2.166H9.506v9.646H6.789v-9.646H3.375z" />
            </svg>

        </g>


    </svg>


    <svg width="80" height="80" x="310" y="90">
        <defs>
            <linearGradient id="linear-gradient4" gradientUnits="objectBoundingBox">
                <stop offset="0" stop-color="#AFD3FC" />
                <stop offset="1" stop-color="#2E79C7" />
            </linearGradient>
        </defs>
        <rect id="gradient-box4" width="80" height="80" rx="20" fill="url(#linear-gradient4)" />


        <g transform="translate(20, 20)">


            <svg fill="white" width="40" height="40" role="img" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <title>Python</title>
                <path
                    d="M14.25.18l.9.2.73.26.59.3.45.32.34.34.25.34.16.33.1.3.04.26.02.2-.01.13V8.5l-.05.63-.13.55-.21.46-.26.38-.3.31-.33.25-.35.19-.35.14-.33.1-.3.07-.26.04-.21.02H8.77l-.69.05-.59.14-.5.22-.41.27-.33.32-.27.35-.2.36-.15.37-.1.35-.07.32-.04.27-.02.21v3.06H3.17l-.21-.03-.28-.07-.32-.12-.35-.18-.36-.26-.36-.36-.35-.46-.32-.59-.28-.73-.21-.88-.14-1.05-.05-1.23.06-1.22.16-1.04.24-.87.32-.71.36-.57.4-.44.42-.33.42-.24.4-.16.36-.1.32-.05.24-.01h.16l.06.01h8.16v-.83H6.18l-.01-2.75-.02-.37.05-.34.11-.31.17-.28.25-.26.31-.23.38-.2.44-.18.51-.15.58-.12.64-.1.71-.06.77-.04.84-.02 1.27.05zm-6.3 1.98l-.23.33-.08.41.08.41.23.34.33.22.41.09.41-.09.33-.22.23-.34.08-.41-.08-.41-.23-.33-.33-.22-.41-.09-.41.09zm13.09 3.95l.28.06.32.12.35.18.36.27.36.35.35.47.32.59.28.73.21.88.14 1.04.05 1.23-.06 1.23-.16 1.04-.24.86-.32.71-.36.57-.4.45-.42.33-.42.24-.4.16-.36.09-.32.05-.24.02-.16-.01h-8.22v.82h5.84l.01 2.76.02.36-.05.34-.11.31-.17.29-.25.25-.31.24-.38.2-.44.17-.51.15-.58.13-.64.09-.71.07-.77.04-.84.01-1.27-.04-1.07-.14-.9-.2-.73-.25-.59-.3-.45-.33-.34-.34-.25-.34-.16-.33-.1-.3-.04-.25-.02-.2.01-.13v-5.34l.05-.64.13-.54.21-.46.26-.38.3-.32.33-.24.35-.2.35-.14.33-.1.3-.06.26-.04.21-.02.13-.01h5.84l.69-.05.59-.14.5-.21.41-.28.33-.32.27-.35.2-.36.15-.36.1-.35.07-.32.04-.28.02-.21V6.07h2.09l.14.01zm-6.47 14.25l-.23.33-.08.41.08.41.23.33.33.23.41.08.41-.08.33-.23.23-.33.08-.41-.08-.41-.23-.33-.33-.23-.41-.08-.41.08z" />
            </svg>

        </g>


    </svg>


    <svg width="80" height="80" x="10" y="210">
        <defs>
            <linearGradient id="linear-gradient4" gradientUnits="objectBoundingBox">
                <stop offset="0" stop-color="#AFD3FC" />
                <stop offset="1" stop-color="#2E79C7" />
            </linearGradient>
        </defs>
        <rect id="gradient-box4" width="80" height="80" rx="20" fill="url(#linear-gradient4)" />


        <g transform="translate(20, 20)">


            <svg fill="white" width="40" height="40" role="img" viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg">
                <title>Svelte</title>
                <path
                    d="M10.354 21.125a4.44 4.44 0 0 1-4.765-1.767 4.109 4.109 0 0 1-.703-3.107 3.898 3.898 0 0 1 .134-.522l.105-.321.287.21a7.21 7.21 0 0 0 2.186 1.092l.208.063-.02.208a1.253 1.253 0 0 0 .226.83 1.337 1.337 0 0 0 1.435.533 1.231 1.231 0 0 0 .343-.15l5.59-3.562a1.164 1.164 0 0 0 .524-.778 1.242 1.242 0 0 0-.211-.937 1.338 1.338 0 0 0-1.435-.533 1.23 1.23 0 0 0-.343.15l-2.133 1.36a4.078 4.078 0 0 1-1.135.499 4.44 4.44 0 0 1-4.765-1.766 4.108 4.108 0 0 1-.702-3.108 3.855 3.855 0 0 1 1.742-2.582l5.589-3.563a4.072 4.072 0 0 1 1.135-.499 4.44 4.44 0 0 1 4.765 1.767 4.109 4.109 0 0 1 .703 3.107 3.943 3.943 0 0 1-.134.522l-.105.321-.286-.21a7.204 7.204 0 0 0-2.187-1.093l-.208-.063.02-.207a1.255 1.255 0 0 0-.226-.831 1.337 1.337 0 0 0-1.435-.532 1.231 1.231 0 0 0-.343.15L8.62 9.368a1.162 1.162 0 0 0-.524.778 1.24 1.24 0 0 0 .211.937 1.338 1.338 0 0 0 1.435.533 1.235 1.235 0 0 0 .344-.151l2.132-1.36a4.067 4.067 0 0 1 1.135-.498 4.44 4.44 0 0 1 4.765 1.766 4.108 4.108 0 0 1 .702 3.108 3.857 3.857 0 0 1-1.742 2.583l-5.589 3.562a4.072 4.072 0 0 1-1.135.499m10.358-17.95C18.484-.015 14.082-.96 10.9 1.068L5.31 4.63a6.412 6.412 0 0 0-2.896 4.295 6.753 6.753 0 0 0 .666 4.336 6.43 6.43 0 0 0-.96 2.396 6.833 6.833 0 0 0 1.168 5.167c2.229 3.19 6.63 4.135 9.812 2.108l5.59-3.562a6.41 6.41 0 0 0 2.896-4.295 6.756 6.756 0 0 0-.665-4.336 6.429 6.429 0 0 0 .958-2.396 6.831 6.831 0 0 0-1.167-5.168Z" />
            </svg>


        </g>


    </svg>
</svg>
*/
