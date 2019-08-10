package main

import (
	"image/gif"
	"log"
	"os"
	"time"

	"github.com/arowshot/fsm"
	"github.com/arowshot/fsm/autofsm"
	"github.com/faiface/pixel"
)

// Player represents a player's avatar in the game
type Player struct {
	HorizontalMove float64
	Jump           bool
	FSM            fsm.FSM
	CurrentAvatar  *pixel.Picture
}

// MakePlayer creates a new player and adds all animation states and transitions
func MakePlayer() *Player {
	player := Player{}

	idleState := autofsm.AutoState{Action: player.showGif(&player, "idle.gif")}              // Create idle state
	player.FSM.AddState("idle", idleState.ToState())                                         // Add to fsm
	idleState.AddTransition("jump", func() bool { return player.Jump })                      // Transition to the jump state if the jump boolean is true
	idleState.AddTransition("walkLeft", func() bool { return player.HorizontalMove < -0.1 }) // Tansition to walkLeft when moving to the left
	idleState.AddTransition("walkRight", func() bool { return player.HorizontalMove > 0.1 }) // Tansition to walkRight when moving to the right

	jumpState := autofsm.AutoState{Action: func(fsm *fsm.FSM) { // Create jumping state
		time.Sleep(2 * time.Second) // Wait 2 seconds
		player.Jump = false         // Turn jump off
	}}
	player.FSM.AddState("jump", jumpState.ToState()) // Add to fsm
	jumpState.AddTransition("idle", nil)             // Transition back to idle state after jump is finished

	walkLeftState := autofsm.AutoState{Action: player.showGif(&player, "run.gif")}           // Create WalkLeftState with the ShowGif action
	player.FSM.AddState("walkLeft", walkLeftState.ToState())                                 // Add to fsm
	walkLeftState.AddTransition("idle", func() bool { return player.HorizontalMove > -0.1 }) // Return to idle if not moving left anymore

	walkRightState := autofsm.AutoState{Action: player.showGif(&player, "run.gif")}          // Create WalkRightState with the ShowGif action
	player.FSM.AddState("walkRight", walkRightState.ToState())                               // Add to fsm
	walkRightState.AddTransition("idle", func() bool { return player.HorizontalMove < 0.1 }) // Return to idle if not moving right anymore

	player.FSM.SetState(idleState.ToState())

	return &player
}

// showGif will return an action that will display one cycle of an animated gif
func (pl *Player) showGif(ply *Player, path string) func(*fsm.FSM) {
	file, err := os.Open(path) // Open the gif
	if err != nil {
		log.Fatalln(err) // Exit if there's an error
	}
	defer file.Close() // Close the file when we're done

	gifs, err := gif.DecodeAll(file) // Decode the gif
	if err != nil {
		log.Fatalln(err) // Exit if there's an error
	}

	images := []pixel.Picture{}
	delay := []int{}

	for i, srcImg := range gifs.Image {
		images = append(images, pixel.PictureDataFromImage(srcImg))
		delay = append(delay, gifs.Delay[i])
	}

	currentFrame := 0
	return func(fsm *fsm.FSM) {
		ply.CurrentAvatar = &images[currentFrame]                               // Set current avatar to the current frame
		delayMillis := time.Duration(delay[currentFrame]*10) * time.Millisecond // Multiple the duration by 10 to get milliseconds
		time.Sleep(delayMillis)                                      // Wait for frame delay

		currentFrame = (currentFrame + 1) % len(images) // Advance the current frame by 1
	}
}
