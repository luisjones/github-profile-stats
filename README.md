# GitHub Profile

This project demonstrates showcasing your GitHub profile through an HTTP server using Go.
Golang rewrite of [Jurredr's WidgetBox](https://github.com/Jurredr/github-widgetbox). (Visual output the same and using the same gradient colours)
Icons are from [Simple Icons](https://simpleicons.org/), [DevIcon](https://devicon.dev/) and [Iconify](https://iconify.design/).

## Installation

1. **Clone the repository:**
2. **Compile Go Binary**
- Navigate to the project directory.
- Run the Go's build command to compile the code into an executable file.
  ```
  go build
  ```
3. **Run HTTP server:**
- Navigate to the project directory, and run the executable which will start at `http://localhost:8081` by default. 
- Alternatively, Run the Go program by running:
  ```
  go run .
  ```

## To-Do
- [ ] Profile Statistics by calling GitHub API
- [ ] Rounded background SVG
- [ ] API Route Parameters
- [ ] Store colours in a local SQLite database
- [ ] WithRow to accept a Title removing the need to call WithParagraph seperately
