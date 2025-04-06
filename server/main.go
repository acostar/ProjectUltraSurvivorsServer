package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

type gameRun struct {
	UserID		string `json:"userID"`
	Log		string `json:"log"`
	ReachedLevel	int `json:"reachedLevel"`
	EnemiesKilled	int `json:"enemiesKilled"`
	Score		int `json:"score"`
}

type gameRunLeaderBoard struct {
	UserID		string `json:"userID"`
	ReachedLevel	int `json:"reachedLevel"`
	EnemiesKilled	int `json:"enemiesKilled"`
	Score		int `json:"score"`
}

var gameRuns = []gameRun{
	{ UserID: "PollaGorda123", Log: "[LevelLogger] LEVEL0-----", ReachedLevel: 4, EnemiesKilled: 5, Score: 512 },
	{ UserID: "PutoAmoMCPene", Log: "[LevelLogger] LEVEL0-----", ReachedLevel: 5, EnemiesKilled: 7, Score: 1024 },
	{ UserID: "MarioKartA90Pavos", Log: "[LevelLogger] LEVEL0-----", ReachedLevel: 6, EnemiesKilled: 10, Score: 2048 },
}

var leaderBoard []gameRunLeaderBoard

func main() {
	router := gin.Default()
	router.GET("/leaderBoard", getLeaderBoard)
	router.POST("/gameRun", postGameRun)

	leaderBoard = appendGameRunToLeaderBoard(gameRuns[0])

	router.Run("localhost:8000")
}

func getLeaderBoard(c *gin.Context) {
	updateLeaderBoard()
	c.IndentedJSON(http.StatusOK, leaderBoard) 
}

func postGameRun(c *gin.Context) {
	log.SetPrefix("PostGameRun: ")
	log.SetFlags(0)

	var newGameRun gameRun

	if err := c.BindJSON(&newGameRun); err != nil { 
		log.Fatal(err)
		return
	}

	gameRuns = append(gameRuns, newGameRun)
	c.IndentedJSON(http.StatusCreated, newGameRun)
}

func updateLeaderBoard() {
	for _, gameRun := range gameRuns {
		userFound := false
		for i, gameRunLeaderBoard := range leaderBoard {
			if gameRun.UserID == gameRunLeaderBoard.UserID {
				if gameRun.ReachedLevel > gameRunLeaderBoard.ReachedLevel {
					leaderBoard[i] = updateGameRunLeaderBoard(gameRunLeaderBoard, gameRun)
				}
				userFound = true
				break
			}
		}
		
		if !userFound { leaderBoard = appendGameRunToLeaderBoard(gameRun) }
	}
}

func appendGameRunToLeaderBoard(gameRun gameRun) []gameRunLeaderBoard {
	return append(leaderBoard, gameRunLeaderBoard{
		UserID: gameRun.UserID,
		ReachedLevel: gameRun.ReachedLevel,
		EnemiesKilled: gameRun.EnemiesKilled,
		Score: gameRun.Score,
	})
}

func updateGameRunLeaderBoard(actualGameRunLeaderBoard gameRunLeaderBoard, gameRun gameRun) gameRunLeaderBoard {
	return gameRunLeaderBoard{
		UserID: gameRun.UserID,
		ReachedLevel: gameRun.ReachedLevel,
		EnemiesKilled: gameRun.EnemiesKilled,
		Score: gameRun.Score,
	}
}
