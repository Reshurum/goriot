package goriot

import (
	"fmt"
	"strconv"
)

type MatchList struct {
	endIndex   int
	Matches    []MatchReference `json:"matches"`
	startIndex int
	totalGames int
}

type MatchReference struct {
	Champion   int64  `json:"champion"`
	Lane       string `json:"lane"`
	MatchID    int64  `json:"matchId"`
	PlatformID string `json:"platformId"`
	Queue      string `json:"queue"`
	Region     string `json:"region"`
	Role       string `json:"role"`
	Season     string `json:"season"`
	Timestamp  int64  `json:"timestamp"`
}

func MatchListBySummonerID(
	region string,
	summonerID int64,
	championIDs []int64,
	rankedQueues []string,
	seasons []string,
	beginTime int64,
	endTime int64,
	beginIndex int,
	endIndex int) (
	matchList MatchList, err error) {

	if !IsKeySet() {
		return matchList, ErrAPIKeyNotSet
	}

	// create a filter for specific champions
	championIDStr := intURLParameter(championIDs).String()

	// check to see if specific queues are being filtered
	rankedQueuesStr := strURLParameter(rankedQueues).String()

	seasonsStr := strURLParameter(seasons).String()

	beginTimeStr := ""
	if beginTime > -1 {
		beginTimeStr = strconv.FormatInt(beginTime, 10)
	}

	endTimeStr := ""
	if endTime > -1 {
		endTimeStr = strconv.FormatInt(endTime, 10)
	}

	// check for indexing
	beginIndexStr := ""
	if beginIndex > -1 {
		beginIndexStr = strconv.Itoa(beginIndex)
	}

	endIndexStr := ""
	if endIndex > -1 {
		endIndexStr = strconv.Itoa(endIndex)
	}

	// build argument string
	args := fmt.Sprintf(
		"api_key=%v&championIds=%v&rankedQueues=%v&seasons=%v&beginTime=%v&endTime=%v&beginIndex=%v&endIndex=%v",
		apikey,
		championIDStr,
		rankedQueuesStr,
		seasonsStr,
		beginTimeStr,
		endTimeStr,
		beginIndexStr,
		endIndexStr)

	// build url string, request, return payload
	url := fmt.Sprintf(
		"https://%v.%v/lol/%v/v2.2/matchlist/by-summoner/%d?%v",
		region,
		BaseAPIURL,
		region,
		summonerID,
		args)
	err = requestAndUnmarshal(url, &matchList)
	if err != nil {
		return matchList, err
	}

	return matchList, nil

}
