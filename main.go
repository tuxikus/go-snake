package main

import (
	"fmt"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIN_X     = 500 * 2
	WIN_Y     = 250 * 2
	TITLE     = "go-snake"
	TILE_SIZE = 25
	FPS       = 60
)

const (
	UP    = 0
	DOWN  = 1
	LEFT  = 2
	RIGHT = 3
)

type Direction int

type World struct {
	tileSize int32
	width    int32
	heigth   int32
}

type SnakeTile struct {
	pos       rl.Vector2
	direction Direction
}

type Snake struct {
	tiles  []SnakeTile
	length uint16
}

type Food struct {
	pos rl.Vector2
}

type Game struct {
	world World
	snake Snake
	score int
	food  Food
}

var game = Game{}

func initGame() {
	game.world = World{
		width:  WIN_X / TILE_SIZE,
		heigth: WIN_Y / TILE_SIZE,
	}

	game.snake.tiles = []SnakeTile{
		{
			direction: RIGHT,
			pos:       rl.Vector2{X: 3, Y: 3},
		},
		{
			direction: RIGHT,
			pos:       rl.Vector2{X: 2, Y: 3},
		},
		{
			direction: RIGHT,
			pos:       rl.Vector2{X: 1, Y: 3},
		},
	}

	game.snake.length = 3
	newFood()
}

func processInput() {
	if rl.IsKeyPressed(rl.KeyUp) {
		if game.snake.tiles[0].direction != DOWN {
			game.snake.tiles[0].direction = UP
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if game.snake.tiles[0].direction != UP {
			game.snake.tiles[0].direction = DOWN
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		if game.snake.tiles[0].direction != RIGHT {
			game.snake.tiles[0].direction = LEFT
		}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		if game.snake.tiles[0].direction != LEFT {
			game.snake.tiles[0].direction = RIGHT
		}
	}
}

func snakeMovement() {
	for i := 0; i < int(game.snake.length); i++ {
		switch game.snake.tiles[i].direction {
		case UP:
			game.snake.tiles[i].pos.Y--
		case DOWN:
			game.snake.tiles[i].pos.Y++
		case LEFT:
			game.snake.tiles[i].pos.X--
		case RIGHT:
			game.snake.tiles[i].pos.X++
		default:
			return
		}
	}
}

func drawSnake() {
	for i := 0; i < int(game.snake.length); i++ {
		rl.DrawRectangle(int32(game.snake.tiles[i].pos.X)*TILE_SIZE, int32(game.snake.tiles[i].pos.Y)*TILE_SIZE, TILE_SIZE, TILE_SIZE, rl.Green)
	}
}

// r r r
//
//	  u
//	r r
func updateDirection() {
	for i := game.snake.length - 1; i > 0; i-- {
		game.snake.tiles[i].direction = game.snake.tiles[i-1].direction
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

func drawFood() {
	rl.DrawRectangle(int32(game.food.pos.X)*TILE_SIZE,
		int32(game.food.pos.Y)*TILE_SIZE,
		TILE_SIZE,
		TILE_SIZE,
		rl.Red)
}

func newFood() {
	game.food.pos.X = float32(rand.IntN(int(game.world.width)-2) + 1)
	game.food.pos.Y = float32(rand.IntN(int(game.world.heigth)-2) + 1)
}

func checkCollision() bool {
	// snake
	for i := 1; i < int(game.snake.length); i++ {
		if game.snake.tiles[i].pos == game.snake.tiles[0].pos {
			return true
		}
	}

	// border
	if game.snake.tiles[0].pos.Y == 0 || game.snake.tiles[0].pos.X == 0 || int32(game.snake.tiles[0].pos.X) == game.world.width-1 || int32(game.snake.tiles[0].pos.Y) == game.world.heigth-1 {
		return true
	}

	// food
	if game.snake.tiles[0].pos.X == game.food.pos.X && game.snake.tiles[0].pos.Y == game.food.pos.Y {
		last := game.snake.tiles[len(game.snake.tiles)-1]
		newPos := last.pos

		switch last.direction {
		case UP:
			newPos.Y += 1
		case DOWN:
			newPos.Y -= 1
		case LEFT:
			newPos.X += 1
		case RIGHT:
			newPos.X -= 1
		}

		game.snake.tiles = append(game.snake.tiles, SnakeTile{
			direction: game.snake.tiles[len(game.snake.tiles)-1].direction,
			pos:       newPos,
		})

		game.snake.length++
		game.score += 10
		newFood()
	}

	return false
}

func drawScore() {
	rl.DrawText(fmt.Sprintf("Score: %d", game.score), TILE_SIZE, 1, 20, rl.Black)
}

func main() {
	initGame()

	rl.InitWindow(WIN_X, WIN_Y, TITLE)
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)

	fc := 0 // frame counter

	for !rl.WindowShouldClose() {
		fc++

		///////////////////////////////////////////////////////////////
		//                           Logic                           //
		///////////////////////////////////////////////////////////////
		processInput()

		if fc >= FPS*1/8 {
			snakeMovement()
			updateDirection()

			if checkCollision() {
				return
			}

			fc = 0
		}

		///////////////////////////////////////////////////////////////
		//                          Drawing                          //
		///////////////////////////////////////////////////////////////
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawBorder()
		drawSnake()
		drawFood()
		drawScore()

		rl.EndDrawing()
	}
}
