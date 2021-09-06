package player

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const leaderboardURL = "http://binod-web.herokuapp.com/leaderboard"

type leaderboardResp struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    map[int]playerEntry `json:"data"`
}

type playerEntry struct {
	Username string `json:"username"`
	Binods   int    `json:"binods"`
}

func getLeaderboard() (*leaderboardResp, error) {
	r, err := http.Get(leaderboardURL)
	if err != nil {
		return &leaderboardResp{}, err
	}

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return &leaderboardResp{}, err
	}

	lb := &leaderboardResp{}
	if err := json.Unmarshal(bodyBytes, lb); err != nil {
		return &leaderboardResp{}, err
	}

	return lb, nil
}

// DisplayLeaderboard renders a colorful table for the binod leaderboard.
func DisplayLeaderboard() {
	color.Cyan("Fetching the leaderboard...")

	leaderboard, err := getLeaderboard()
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Printf("Unable to fetch the leaderboard! The following error occurred: \n%v.", err.Error())
		return
	}

	if !leaderboard.Success {
		color.New(color.FgRed).Add(color.Underline).Printf("Unable to fetch the leaderboard! The following error occurred: \n%v.", leaderboard.Message)
		return
	}

	if len(leaderboard.Data) == 0 {
		color.Yellow("No entries in the binod leaderboard so far!")
		return
	}

	color.Green("Leaderboard fetched successfully.\n")
	data := [][]string{}
	for i := 0; i < len(leaderboard.Data); i++ {
		data = append(data,
			[]string{
				fmt.Sprintf("%d", i+1),
				leaderboard.Data[i+1].Username,
				fmt.Sprintf("%d", leaderboard.Data[i+1].Binods),
			})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Player", "Binod count"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgYellowColor,
			tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor, tablewriter.FgWhiteColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Italic, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	table.AppendBulk(data)
	table.Render()
}
