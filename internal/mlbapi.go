package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

type mlbAPIResponse struct {
	Dates []mlbAPIDate
}

type mlbAPIDate struct {
	Games []mlbAPIGame
}

type mlbAPIGame struct {
	Description string
	GameID      int    `json:"gamePk"`
	GameAPIURL  string `json:"link"`
	Teams       mlbAPITeams
	Status      mlbAPIGameStatus
}

type mlbAPITeams struct {
	AwayTeam mlbAPITeam `json:"away"`
	HomeTeam mlbAPITeam `json:"home"`
}

type mlbAPITeam struct {
	Record   mlbAPITeamRecord `json:"leagueRecord"`
	Score    int
	TeamInfo mlbAPITeamInfo `json:"team"`
}

type mlbAPITeamInfo struct {
	ID   int
	Name string
}

type mlbAPITeamRecord struct {
	Wins   int
	Losses int
}

type mlbAPIGameStatus struct {
	AbstractGameState string
}

func getCurrentLiveGames(date string, logger *logrus.Logger) {
	apiURL := gamesAPIURL(date)

	mlbAPIClient := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, apiURL, nil)

	response, apiErr := mlbAPIClient.Do(req)
	if apiErr != nil {
		logger.Fatal(apiErr)
	}

	responseBody, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		logger.Fatal(readErr)
	}

	mlbAPIResponse := mlbAPIResponse{}
	jsonErr := json.Unmarshal(responseBody, &mlbAPIResponse)
	if jsonErr != nil {
		logger.Fatal(jsonErr)
	}

	if len(mlbAPIResponse.Dates) == 0 {
		// don't do anything?
	}

	for _, dates := range mlbAPIResponse.Dates {
		for _, game := range dates.Games {
			awayTeamInfo := game.Teams.AwayTeam
			homeTeamInfo := game.Teams.HomeTeam
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("%s - %s - ", game.Status.AbstractGameState, game.Description))
			builder.WriteString(fmt.Sprintf("%s (%d-%d)", awayTeamInfo.TeamInfo.Name, awayTeamInfo.Record.Wins, homeTeamInfo.Record.Wins))
			builder.WriteString(fmt.Sprint(" @ "))
			builder.WriteString(fmt.Sprintf("%s (%d-%d)", homeTeamInfo.TeamInfo.Name, homeTeamInfo.Record.Wins, awayTeamInfo.Record.Wins))

			logger.Println(builder.String())
		}
	}
}

func gamesAPIURL(date string) string {
	url := url.URL{
		Scheme: "https",
		Host:   "statsapi.mlb.com",
	}
	url.Path = path.Join(url.Path, "api/v1/schedule/games")
	q := url.Query()
	q.Add("sportId", "1")
	q.Add("date", date)
	url.RawQuery = q.Encode()

	return url.String()
}
