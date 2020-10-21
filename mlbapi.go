package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type mlbAPIResponse struct {
	Dates []mlbAPIDate
}

type mlbAPIDate struct {
	Games []mlbAPIGame
}

type mlbAPIGame struct {
	GameID     int    `json:"gamePk"`
	GameAPIURL string `json:"link"`
	Teams      mlbAPITeams
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

// todaysDate := time.Now().Format("01/02/2006")
// datePtr := flag.String("date", todaysDate, "the date to get MLB games for (defaults to today)")
// flag.Parse()

// printCurrentLiveGames(*datePtr)

func printCurrentLiveGames(date string) {
	apiURL := gamesAPIURL(date)

	mlbAPIClient := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, apiURL, nil)

	response, apiErr := mlbAPIClient.Do(req)
	if apiErr != nil {
		log.Fatal(apiErr)
	}

	responseBody, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	mlbAPIResponse := mlbAPIResponse{}
	jsonErr := json.Unmarshal(responseBody, &mlbAPIResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(mlbAPIResponse.Dates) == 0 {

	}
	for _, dates := range mlbAPIResponse.Dates {
		fmt.Println()
		fmt.Printf("MLB games on %s:\n", date)
		for _, game := range dates.Games {
			awayTeamInfo := game.Teams.AwayTeam
			homeTeamInfo := game.Teams.HomeTeam
			fmt.Printf("- %s (%d-%d)", awayTeamInfo.TeamInfo.Name, awayTeamInfo.Record.Wins, homeTeamInfo.Record.Wins)
			fmt.Print(" @ ")
			fmt.Printf("%s (%d-%d)", homeTeamInfo.TeamInfo.Name, homeTeamInfo.Record.Wins, awayTeamInfo.Record.Wins)
			fmt.Println()
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
