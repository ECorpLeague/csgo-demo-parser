package main

import (
	"fmt"
	"os"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
    "io/ioutil"
	"encoding/json"
)

type PlayerH struct {
    Player string `json:"player"`
    HealthDamage int `json:"healthDamage"`
}

func main() {

	f, err := os.Open("testdemo.dem")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	// Register handler on kill events
	// p.RegisterEventHandler(func(e events.Kill) {
	// 	var hs string
	// 	if e.IsHeadshot {
	// 		hs = " (HS)"
	// 	}
	// 	var wallBang string
	// 	if e.PenetratedObjects > 0 {
	// 		wallBang = " (WB)"
	// 	}
	// 	fmt.Printf("%s <%v%s%s> %s\n", e.Killer, e.Weapon, hs, wallBang, e.Victim)
	// })


	var playerHData []PlayerH

	p.RegisterEventHandler(func(e events.PlayerHurt) {

		fmt.Println(formatPlayer(e.Player))

		playerHData = append(playerHData, PlayerH{
            Player: formatPlayer(e.Player),
            HealthDamage: e.HealthDamage,
        })
	})


	// Parse to end
	err = p.ParseToEnd()

	b, _ := json.Marshal(playerHData)
		fmt.Println(string(b))

		_ = ioutil.WriteFile("file.json", b, 0644)

	if err != nil {
		panic(err)
	}
}

func formatPlayer(p *common.Player) string {
	if p == nil {
		return "?"
	}

	return p.Name
}