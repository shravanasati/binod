package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type LeaderBoardResp struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    map[int]PlayerEntry `json:"data"`
}

type PlayerEntry struct {
	Username string `json:"username"`
	Binods   int    `json:"binods"`
}

func getLeaderboard() (*LeaderBoardResp, error) {
	r, err := http.Get("http://binod-web.herokuapp.com/leaderboard")
	if err != nil {
		return &LeaderBoardResp{}, err
	}

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return &LeaderBoardResp{}, err
	}

	lb := &LeaderBoardResp{}
	if err := json.Unmarshal(bodyBytes, lb); err != nil {
		return &LeaderBoardResp{}, err
	}

	return lb, nil
}

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
