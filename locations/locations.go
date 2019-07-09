package locations

type PathBotDirection struct {
    Direction string 'json:"direction"'
}

type PathBotLocation struct {
    Status string 'json:"status"'
    Message string 'json:"message"'
    Exits []string 'json:"exits"'
    Description string 'json:"description"'
    MazeExitDirection string 'json:"mazeExitDirection"'
    MazeExitDistance int 'json:"mazeExitDistance"'
    LocationPath string 'json:"locationPath"'
}

const baseURL = "https://api.noopschallenge.com"

func RequestLocation() (location PathBotLocation) {
    return
}
