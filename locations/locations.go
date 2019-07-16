package locations

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "strings"
)

type PathBotDirection struct {
    Direction string `json:"direction"`
}

type PathBotLocation struct {
    Status            string   `json:"status"`
    Message           string   `json:"message"`
    Exits             []string `json:"exits"`
    Description       string   `json:"description"`
    MazeExitDirection string   `json:"mazeExitDirection"`
    MazeExitDistance  int      `json:"mazeExitDistance"`
    LocationPath      string   `json:"locationPath"`
}

const baseURL = "https://api.noopschallenge.com"

func Start() PathBotLocation {
    return apiPost("/pathbot/start", strings.NewReader("{}"))
}

func Go(location PathBotLocation, direction PathBotDirection) (*PathBotLocation, error) {
    body, err := json.Marshal(direction)
    if err != nil {return nil, err}
    newLocation := apiPost(location.LocationPath, bytes.NewBuffer(body))
    return &newLocation, nil
}

func apiPost(path string, body io.Reader) PathBotLocation {
    res, err := http.Post(baseURL+path, "application/json", body)
    if err != nil { panic(err.Error()) }
    return parseResponse(res)
}

func parseResponse(res *http.Response) (loc PathBotLocation) {
    body, err := ioutil.ReadAll(res.Body)
    if err != nil { panic(err.Error()) }
    err = json.Unmarshal(body, &loc)
    if err != nil { panic(err.Error()) }
    return loc
}