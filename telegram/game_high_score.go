package telegram

/*
GameHighScore represents one row of the high scores table for a game.
*/
type GameHighScore struct {
	Position int64 `json:"position"` // Position in high score table for the game
	User     *User `json:"user"`     // User
	Score    int64 `json:"score"`    // Score
}
