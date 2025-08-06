package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WIN_X     = 500 * 2
	WIN_Y     = 250 * 2
	TITLE     = "go-snake"
	TILE_SIZE = 50
)

type World struct {
	tileSize int32
	width    int32
	heigth   int32
}

type Snake struct {
	pos rl.Vector2
}

type Game struct {
	world World
	snake Snake
	score int
}

var game = Game{}

func initGame() {
	game.world = World{
		width:  WIN_X / TILE_SIZE,
		heigth: WIN_Y / TILE_SIZE,
	}
}

func drawBorder() {
	for y := range game.world.heigth {
		for x := range game.world.width {
			if y == 0 || x == 0 || x == game.world.width-1 || y == game.world.heigth-1 {
				rl.DrawRectangle(x*TILE_SIZE, y*TILE_SIZE, TILE_SIZE, TILE_SIZE, rl.Gray)
			}
		}
	}
}

func main() {
	initGame()

	rl.InitWindow(WIN_X, WIN_Y, TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// logic

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// drawing

		drawBorder()

		rl.EndDrawing()
	}
}
