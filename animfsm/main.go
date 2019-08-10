package main

import (
	"fmt"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var pl *Player

func run() {
	pl = MakePlayer()
	done := pl.FSM.RunMachine()
	defer close(done)

	cfg := pixelgl.WindowConfig{
		Title:  "FSM",
		Bounds: pixel.R(0, 0, 256, 256),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	playerInfo := text.New(pixel.V(0, win.Bounds().Max.Y-10), atlas)
	playerInfo.Color = colornames.Black

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		playerInfo.Clear()
		fmt.Fprintf(playerInfo, "Horizontal speed: %v\n", pl.HorizontalMove)
		fmt.Fprintf(playerInfo, "Jumping: %v\n", pl.Jump)
		fmt.Fprintf(playerInfo, "Current state: %v\n", pl.FSM.GetCurrentState())

		win.Clear(colornames.Skyblue)

		if win.Pressed(pixelgl.KeyLeft) {
			pl.HorizontalMove = math.Max(-1, pl.HorizontalMove-(dt*10))
		} else if win.Pressed(pixelgl.KeyRight) {
			pl.HorizontalMove = math.Min(1, pl.HorizontalMove+(dt*10))
		} else {
			pl.HorizontalMove = 0
		}

		if win.JustPressed(pixelgl.KeySpace) {
			pl.Jump = true
		}

		if pl.CurrentAvatar != nil {
			sprite := pixel.NewSprite(*pl.CurrentAvatar, (*pl.CurrentAvatar).Bounds())

			sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()).ScaledXY(win.Bounds().Center(), pixel.V(4, 4)))
		}

		playerInfo.Draw(win, pixel.IM)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
