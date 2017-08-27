package steamtop

import (
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/pkg/errors"
	"github.com/yhat/scrape"
)

// Game represents a single game with the current player count and the day's peak player count
type Game struct {
	Name           string `json:"name"`
	CurrentPlayers int    `json:"current_players"`
	PeakPlayers    int    `json:"peak_players"`
}

// GetSteamTopGames returns slice of Games, same order they appear on site: most-to-least players
func GetSteamTopGames() (result []Game, err error) {
	response, err := http.Get("http://store.steampowered.com/stats/")
	if err != nil {
		return result, errors.Wrap(err, "failed to GET steam stats page")
	}

	if response.StatusCode > 299 {
		return result, errors.Errorf("response status was %s", response.Status)
	}

	root, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Tr {
			return scrape.Attr(n, "class") == "player_count_row"
		}
		return false
	}

	rows := scrape.FindAll(root, matcher)

	for _, row := range rows {
		columns := scrape.FindAllNested(row, scrape.ByTag(atom.Span))
		nameElement, ok := scrape.Find(row, scrape.ByTag(atom.A))
		if !ok {
			continue
		}

		game := Game{Name: scrape.Text(nameElement)}

		if len(columns) == 2 {
			game.CurrentPlayers, err = strconv.Atoi(strings.Replace(scrape.Text(columns[0]), ",", "", -1))
			if err != nil {
				return result, errors.Wrap(err, "bad number in stats table")
			}
			game.PeakPlayers, err = strconv.Atoi(strings.Replace(scrape.Text(columns[1]), ",", "", -1))
			if err != nil {
				return result, errors.Wrap(err, "bad number in stats table")
			}
		}

		result = append(result, game)
	}

	return result, nil
}
