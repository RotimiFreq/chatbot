package main

import (
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"
	"net/http"
	"strconv"
)

// A simple football statisctis bot using functions

// Functions that consumes API endpoints

func main() {
	footballStat()
}

// creating some global variables.

var i int
var clubName string
var playerName string
var att int
var season string
var leID int
var leagueID int
var gog string
var ID_Team int
var Name_Team string

func footballStat() {

	//
	fmt.Printf("This is samuel's bot \n")
	fmt.Printf("What do you want to do today? \n")
	print("Press 1 for club stats.\nPress 2 for player stats.\n ")
	print("Input your option:")
	fmt.Scan(&i)

	// conditionals to check for either club or player statistics

	if i == 1 {
		println("Which club do you want to check thier statistics:?")
		fmt.Scan(&clubName)

		// we neeed to get the league of the team

		print("Which leaugue is the club playing?\n")

		print("1.EPL\n 2.Laliga\n 3.Seria-A\n 4.Bundesliga\n 5.Ligue-1\n 6.Eredivise\n 7.Premira-Liga\n")
		print("Input your option:")
		fmt.Scan(&leID)

		//we use the switch statement to assign the league id
		switch {
		case leID == 1:
			leagueID = 39

		case leID == 2:
			leagueID = 140

		case leID == 3:
			leagueID = 135

		case leID == 4:
			leagueID = 78

		case leID == 5:
			leagueID = 61

		case leID == 6:
			leagueID = 88

		case leID == 7:
			leagueID = 94
		}

		println("Which season do you want to check?")
		fmt.Scan(&season)

		// Getting data from user on which statistics is to be pulled

		println("Press 1 for attacking statistics and 2 for defensive stats:")
		fmt.Scan(&att)

		// calling  the neccesary functions after getting the data needed.

		// This function is to get the team Id in the database
		teamID()

		//This function gets the needed statistics
		apiFootball()

	}
	// condition for player stats

	if i == 2 {

		// getting data from users

		println("Which player statistics do you want to check :?")
		fmt.Scan(&playerName)

		println("Which club does the player plays in :?")
		fmt.Scan(&clubName)

		print("Which leaugue is the club playing?\n")
		print("1.EPL \n 2.Laliga\n 3.Seria-A\n 4.Bundesliga\n 5.Ligue-1\n 6.Eredivise\n 7.Premira-Liga\n")
		print("Input your option:")
		fmt.Scan(&leID)

		//we use the switch statement to assign the league id
		switch {
		case leID == 1:
			leagueID = 39

		case leID == 2:
			leagueID = 140

		case leID == 3:
			leagueID = 135

		case leID == 4:
			leagueID = 78

		case leID == 5:
			leagueID = 61

		case leID == 6:
			leagueID = 88

		case leID == 7:
			leagueID = 94
		}

		println("Which season do you want to check?")
		fmt.Scan(&season)

		teamID()

		playerstat()

	}
	defer footballStat()
}

// create a struct that would contain the json file
// to be returned after the get request

type JsonReturn struct {
	// an array of structs
	JSResponse []JSResponse `json:"response"`
}
type JSResponse struct {
	TeamData  theTeam  `json:"team"`
	VenueData theVenue `json:"venue"`
}

type theTeam struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}

type theVenue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	Image    string `json:"image"`
}

func teamID() {

	// concantenate the values from the user into the url

	url := "https://api-football-v1.p.rapidapi.com/v3/teams?name=" + clubName + "&league=" + strconv.Itoa(leagueID) + "&season=" + season

	// create a variable to request function
	req, _ := http.NewRequest("GET", url, nil)

	// make a http request to the server

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "fbcea68969msh557e2ac40ae66b4p1db147jsn1540f6b1e58c")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// close until the function is done
	defer res.Body.Close()
	//using the OS package to get our json file
	body, err := ioutil.ReadAll(res.Body)

	// log error if anything fails
	if err != nil {
		log.Fatal(err)
	}

	// converting the body in to byte data
	body_byte := []byte(body)

	// let create an array for the above struct declared

	var JsonArray JsonReturn

	// We unmarshal(turn over) the data from json file into the struct created

	err = json.Unmarshal(body_byte, &JsonArray)

	if err != nil {

		fmt.Println(err)
	}

	// create new variable to access the struct

	teamid := JsonArray.JSResponse[0]
	teamid2 := teamid.TeamData

	// get needed field and assign them to global variables
	Name_Team = teamid2.Name
	ID_Team = teamid2.ID

	fmt.Println(Name_Team)

}

// same process for the first functions
type returnArray struct {
	Response Response `json:"response"`
}

type Response struct {
	League struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
		Season  int    `json:"season"`
	}
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	}
	Form     string `json:"form"`
	Fixtures struct {
		Played struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		}
		Wins struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		}
		Draws struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		}
		Loses struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		}
	}
	Goals struct {
		For struct {
			Total struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			}
			Average struct {
				Home  string `json:"home"`
				Away  string `json:"away"`
				Total string `json:"total"`
			}
			Minute struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"46-60"`
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"61-75"`
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				}
			}
		}
		Against struct {
			Total struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			}
			Average struct {
				Home  string `json:"home"`
				Away  string `json:"away"`
				Total string `json:"total"`
			}
			Minute struct {
				Zero15 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				One630 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Three145 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Four660 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Six175 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Seven690 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				Nine1105 struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				}
				One06120 struct {
					Total      interface{} `json:"total"`
					Percentage interface{} `json:"percentage"`
				}
			}
		}
	}
	Biggest struct {
		Streak struct {
			Wins  int `json:"wins"`
			Draws int `json:"draws"`
			Loses int `json:"loses"`
		}
		Wins struct {
			Home string `json:"home"`
			Away string `json:"away"`
		}
		Loses struct {
			Home string `json:"home"`
			Away string `json:"away"`
		}
		Goals struct {
			For struct {
				Home int `json:"home"`
				Away int `json:"away"`
			}
			Against struct {
				Home int `json:"home"`
				Away int `json:"away"`
			}
		}
	}
	CleanSheet struct {
		Home  int `json:"home"`
		Away  int `json:"away"`
		Total int `json:"total"`
	}
	FailedToScore struct {
		Home  int `json:"home"`
		Away  int `json:"away"`
		Total int `json:"total"`
	}
	Penalty struct {
		Scored struct {
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
		}
		Missed struct {
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
		}
		Total int `json:"total"`
	}
	Lineups []struct {
		Formation string `json:"formation"`
		Played    int    `json:"played"`
	}
	Cards struct {
		Yellow struct {
			Zero15 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			One630 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			Three145 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			Four660 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			Six175 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			Seven690 struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			}
			Nine1105 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			One06120 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
		}
		Red struct {
			Zero15 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			One630 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			Three145 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			Four660 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			Six175 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			Seven690 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			Nine1105 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
			One06120 struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
			}
		}
	}
}

func apiFootball() {

	url := "https://api-football-v1.p.rapidapi.com/v3/teams/statistics?league=" + strconv.Itoa(leagueID) + "&season=" + season + "&team=" + strconv.Itoa(ID_Team)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "fbcea68969msh557e2ac40ae66b4p1db147jsn1540f6b1e58c")

	res, error := http.DefaultClient.Do(req)
	if error != nil {
		log.Fatal(error)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	stat_byte := []byte(body)

	var statsRtn returnArray

	error1 := json.Unmarshal(stat_byte, &statsRtn)

	if error1 != nil {
		log.Fatal(error1)

	}

	stats_return := statsRtn.Response
	form := stats_return.Form

	fixture_win := stats_return.Fixtures.Wins
	fixture_loss := stats_return.Fixtures.Loses
	fixtures_draw := stats_return.Fixtures.Draws

	goals := stats_return.Goals
	goals_for := goals.For
	goal_for_average := goals.For.Average
	goals_aganist := goals.Against
	goals_aganist_average := goals_aganist.Average
	biggest := stats_return.Biggest
	biggest_goal := biggest.Goals
	biggest_goal_for := biggest_goal.For
	biggest_goal_aganist := biggest_goal.Against
	biggest_loss := biggest.Loses
	biggest_win := biggest.Wins
	biggest_streak := biggest.Streak
	cleansheet := stats_return.CleanSheet
	failed_score := stats_return.FailedToScore
	penalties := stats_return.Penalty
	lineup_used := stats_return.Lineups

	Attacking_Statistics := func() {
		fmt.Printf("%s attacking statistics for the season %s \n ", Name_Team, season)
		fmt.Println("Lineup used  throughout the season(number of times used):\n", lineup_used)
		fmt.Println("Form for the season:\n", form)

		fmt.Println("Games won at home :\n", fixture_win.Home)
		fmt.Println("Games won Away:\n", fixture_win.Away)

		fmt.Println("Games loss at home :\n", fixture_loss.Home)
		fmt.Println("Games loss away :\n", fixture_loss.Away)

		fmt.Println("Games draw at home :\n", fixtures_draw.Home)
		fmt.Println("Games draw Away :\n", fixtures_draw.Away)
		fmt.Println("Goals scored at home:\n ", goals_for.Total.Home)
		fmt.Println("Goals scored away :\n", goals_for.Total.Away)

		fmt.Println("Goals average at home:\n", goal_for_average.Home)
		fmt.Println("Goals average away:\n", goal_for_average.Away)
		fmt.Println("Penalties scored: \n", penalties.Scored)
		fmt.Println("Penalties missed: \n ", penalties.Missed)

		fmt.Println("Number of games won consecutively:\n ", biggest_streak.Wins)
		fmt.Println("Number of games lost consecutively:\n ", biggest_streak.Loses)
		fmt.Println("Number of games drawn consecutively :\n", biggest_streak.Draws)

		fmt.Println("Biggest win at home ", biggest_win.Home)
		fmt.Println("Biggest win Away ", biggest_win.Away)

		fmt.Println("biggest goal margin at home", biggest_goal_for.Home)
		fmt.Println("biggest goal margin Away", biggest_goal_for.Away)

		fmt.Println("Number of games failed to score at home", failed_score.Home)
		fmt.Println("Number of games failed to score at away", failed_score.Away)

	}

	Defensive_Statictics := func() {

		fmt.Printf("%s Defensive statistics for the season %s :- \n", Name_Team, season)

		fmt.Println("Lineup used throughout the season(number of times used):\n", lineup_used)
		fmt.Println("Form for the season:\n", form)

		fmt.Println("Goals conceded at home:\n ", goals_aganist.Total.Home)
		fmt.Println("Goals conceded Away:\n ", goals_aganist.Total.Away)
		fmt.Println("Goals conceded average at home:\n", goals_aganist_average.Home)
		fmt.Println("Goals conceded average away:\n", goals_aganist_average.Away)
		fmt.Println("Biggest loss at Home:\n", biggest_loss.Home)
		fmt.Println("Biggest loss Away:\n", biggest_loss.Away)
		fmt.Println("Biggest loss margin at home: \n", biggest_goal_aganist.Home)
		fmt.Println("Biggest loss margin Away: \n", biggest_goal_aganist.Away)
		fmt.Println(" Number of Cleansheet at home:\n ", cleansheet.Home)
		fmt.Println(" Number of Cleansheet at away:\n ", cleansheet.Away)

	}

	switch {

	case att == 1:
		Attacking_Statistics()
	case att == 2:
		Defensive_Statictics()

	}

}

type playerstat_struct struct {
	JsonPlayerstat []JsonPlayerstat `json:"response"`
}

type JsonPlayerstat struct {
	Player struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Age       int    `json:"age"`
		Birth     struct {
			Date    string `json:"date"`
			Place   string `json:"place"`
			Country string `json:"country"`
		}
		Nationality string `json:"nationality"`
		Height      string `json:"height"`
		Weight      string `json:"weight"`
		Injured     bool   `json:"injured"`
		Photo       string `json:"photo"`
	}
	Statistics []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		}
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		}
		Games struct {
			Appearences int         `json:"appearences"`
			Lineups     int         `json:"lineups"`
			Minutes     int         `json:"minutes"`
			Number      interface{} `json:"number"`
			Position    string      `json:"position"`
			Rating      string      `json:"rating"`
			Captain     bool        `json:"captain"`
		}
		Substitutes struct {
			In    int `json:"in"`
			Out   int `json:"out"`
			Bench int `json:"bench"`
		}
		Shots struct {
			Total int `json:"total"`
			On    int `json:"on"`
		}
		Goals struct {
			Total    int         `json:"total"`
			Conceded interface{} `json:"conceded"`
			Assists  int         `json:"assists"`
			Saves    interface{} `json:"saves"`
		}
		Passes struct {
			Total    int `json:"total"`
			Key      int `json:"key"`
			Accuracy int `json:"accuracy"`
		}
		Tackles struct {
			Total         int `json:"total"`
			Blocks        int `json:"blocks"`
			Interceptions int `json:"interceptions"`
		}
		Duels struct {
			Total int `json:"total"`
			Won   int `json:"won"`
		}
		Dribbles struct {
			Attempts int         `json:"attempts"`
			Success  int         `json:"success"`
			Past     interface{} `json:"past"`
		}
		Fouls struct {
			Drawn     int `json:"drawn"`
			Committed int `json:"committed"`
		}
		Cards struct {
			Yellow    int `json:"yellow"`
			Yellowred int `json:"yellowred"`
			Red       int `json:"red"`
		}
		Penalty struct {
			Won      int         `json:"won"`
			Commited interface{} `json:"commited"`
			Scored   int         `json:"scored"`
			Missed   int         `json:"missed"`
			Saved    interface{} `json:"saved"`
		}
	}
}

func playerstat() {

	fmt.Println("this the player stat func")

	url := "https://api-football-v1.p.rapidapi.com/v3/players?team=" + strconv.Itoa(ID_Team) + "&league=" + strconv.Itoa(leagueID) + "&season=" + season + "&search=" + playerName

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "fbcea68969msh557e2ac40ae66b4p1db147jsn1540f6b1e58c")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	stat_byte_player := []byte(body)

	var Stat_Player playerstat_struct

	jsonerr := json.Unmarshal(stat_byte_player, &Stat_Player)
	if jsonerr != nil {
		log.Fatal(jsonerr)
	}

	player := Stat_Player.JsonPlayerstat[0]

	player_profile := player.Player
	player_statistics := player.Statistics

	playerfullname := player_profile.Firstname + player_profile.Lastname
	games := player_statistics[0]
	games_played := games.Games
	App := games_played.Appearences
	minute := games_played.Minutes
	position := games_played.Position
	ave_rating := games_played.Rating
	sub := games.Substitutes
	sub_in := sub.In
	sub_out := sub.Out
	bench := sub.Bench
	goals_full := games.Goals
	Goal := goals_full.Total
	assist := goals_full.Assists

	saves := goals_full.Saves
	concedded := goals_full.Conceded
	Pen := games.Penalty

	passing := games.Passes
	totalPass := passing.Total
	keyPass := passing.Key
	accuracy := passing.Accuracy

	penSco := Pen.Scored
	PenMissed := Pen.Missed
	penWon := Pen.Won
	penSave := Pen.Saved

	SS := games.Shots
	total_shot := SS.Total
	total_shot_on := SS.On
	Drib := games.Dribbles
	totslDrib := Drib.Attempts
	totslcomp := Drib.Success
	dribAganist := Drib.Past

	Tackles := games.Tackles
	tackmade := Tackles.Total
	tack_inter := Tackles.Interceptions
	tack_block := Tackles.Blocks

	fouls := games.Fouls

	commitedFoul := fouls.Committed

	duels := games.Duels
	duelsTotal := duels.Total
	duelsWon := duels.Won

	cards := games.Cards
	yellow := cards.Yellow
	yelred := cards.Yellowred
	red := cards.Red

	fmt.Println("player full name:\n", playerfullname)

	fmt.Println("Position played:\n", position)

	fmt.Println("Avarage rating :\n", ave_rating)

	fmt.Println("number of games played\n", App)

	fmt.Println("total minutes played:\n", minute)

	fmt.Println("number of games player was subtituted in:\n", sub_in)

	fmt.Println("number of games player was subtituted out:\n", sub_out)

	fmt.Println("number of games player was on the bench:\n", bench)

	fmt.Println("Total number of Goals scored :\n", Goal)

	fmt.Println("total number of shots:\n", total_shot)

	fmt.Println("total number of shots on target:\n", total_shot_on)

	fmt.Println("Penalties scored:\n", penSco)

	fmt.Println("Penalties missed:\n", PenMissed)

	fmt.Println("Penalties won:\n", penWon)

	fmt.Println("Penalties concedded:\n", Pen.Commited)

	fmt.Println("Penalties saved:\n", penSave)

	fmt.Println("Saves:\n", saves)

	fmt.Println("goals concedded:\n", concedded)

	fmt.Println("Total number of Assist :\n", assist)

	fmt.Println("total number of passes:\n", totalPass)

	fmt.Println("Key passes :\n", keyPass)

	fmt.Println("Accuracy of passes:\n", accuracy)

	fmt.Println("Dribbles Attempted:\n", totslDrib)

	fmt.Println("Dribbles completed:\n", totslcomp)

	fmt.Println("total nummber of times dribbles past:\n", dribAganist)

	fmt.Println("total tackles made:\n", tackmade)

	fmt.Println("interception completed:\n", tack_inter)

	fmt.Println("Blocks made:\n", tack_block)

	fmt.Println("total number of duels:\n", duelsTotal)

	fmt.Println("Duels Won:\n", duelsWon)

	fmt.Println("Fouls committed:\n", commitedFoul)

	fmt.Println("Number of yellow card:\n", yellow)

	fmt.Println("Number of second yellow:\n", yelred)

	fmt.Println("Red cards:\n", red)

}

//
//
//
//
//
//
//
//
// kindly go through my codes and tell me what i should improve on.
//
//
//
//
//
// Thank you.
