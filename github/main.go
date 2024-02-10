package github

// GET /repos/{owner}/{repo}/languages

var endpoint = "https://api.github.com"

/*
{
  "title": "Language",
  "description": "Language",
  "type": "object",
  "additionalProperties": {
    "type": "integer"
  }
}
*/

type languageResponse struct {
    title string
    description string
    type string
    additionalProperties interface{}
}

// GET /user/repos
// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repositories-for-the-authenticated-user