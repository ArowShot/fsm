//go:generate go-bindata -o assets.go assets/
// This will package all of the asset files to be used in the binaties

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
	pl = MakePlayer() // Create the player
	done := pl.FSM.RunMachine() // Start the animation FSM
	defer close(done) // Stop the FSM when it's done

	cfg := pixelgl.WindowConfig{ // Specify some options for the game window
		Title:  "FSM",
		Bounds: pixel.R(0, 0, 256, 256),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg) // Create the game window using the config
	if err != nil {
		panic(err) // Exit if there's an error creating the window
	}

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII) // Create text atlas (basically a font)

	playerInfo := text.New(pixel.V(0, win.Bounds().Max.Y-10), atlas) // Create playerInfo text
	playerInfo.Color = colornames.Black // Set playerInfo text color to black

	last := time.Now() // Keep track of the last frame
	for !win.Closed() {
		dt := time.Since(last).Seconds() // Store how long the frame took
		last = time.Now() // Reset the time of the last frame

		playerInfo.Clear()// Clear the playerInfo text before updating it

		// Ouput information about the player to the playerInfo text
		fmt.Fprintf(playerInfo, "Horizontal speed: %v\n", pl.HorizontalMove)
		fmt.Fprintf(playerInfo, "Jumping: %v\n", pl.Jump)
		fmt.Fprintf(playerInfo, "Current state: %v\n", pl.FSM.GetCurrentState())

		win.Clear(colornames.Skyblue) // Clear the canvas

		if win.Pressed(pixelgl.KeyLeft) {
			pl.HorizontalMove = math.Max(-1, pl.HorizontalMove-(dt*10)) // Increase left movement up to 1 while the left arrow is being presse
		} else if win.Pressed(pixelgl.KeyRight) {
			pl.HorizontalMove = math.Min(1, pl.HorizontalMove+(dt*10)) // Increase right movement up to 1 while the right arrow is being pressed
		} else {
			pl.HorizontalMove = 0 // When no directional buttons are pressed reset the movement to 0
		}

		if win.JustPressed(pixelgl.KeySpace) {
			pl.Jump = true // Set the player's jump value to true when the spacd bar gets pressed
		}

		if pl.CurrentAvatar != nil { // Draw if the player's avatar exists
			sprite := pixel.NewSprite(*pl.CurrentAvatar, (*pl.CurrentAvatar).Bounds()) // Create the player sprite from the current player texture

			sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()).ScaledXY(win.Bounds().Center(), pixel.V(4, 4))) // Draw the player at the center scaled by 4
		}

		playerInfo.Draw(win, pixel.IM) // Draw the player information text

		win.Update() // Update the window
	}
}

func main() {
	pixelgl.Run(run) // Run the main function in the main thread (required by opengl)
}
