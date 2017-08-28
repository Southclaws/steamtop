# steamtop

[![Travis](https://img.shields.io/travis/Southclaws/steamtop.svg)](https://travis-ci.org/Southclaws/steamtop)

Super simply one-function package to grab the data from http://store.steampowered.com/stats.

```go
func main() {
    games, _ := GetSteamTopGames()

    // prints the top 10 games for today
    for _, game := range games[:10] {
		fmt.Printf("%s: %d players, %d peak\n", game.Name, game.CurrentPlayers, game.PeakPlayers)
    }
}
```
