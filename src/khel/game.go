package khel

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	EM           EntityManager
	Paused       bool
	FrameNo      int
	PlayerParams map[string]float64
	EntityParams map[string]float64
}

func NewGame() *Game {
	// read config
	return &Game{
		EM:           *NewEntityManager(),
		PlayerParams: map[string]float64{},
		EntityParams: map[string]float64{},
	}
}

func (g *Game) Run() {
	rl.InitWindow(501, 621, "Geometry")
	defer rl.CloseWindow()
	// rl.InitAudioDevice()
	// defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)
	g.SpawnPlayer()
	g.SpawnEnemy()
	ticker := time.NewTicker(1000 / 60 * time.Millisecond)
	for !rl.WindowShouldClose() {
		// rl.UpdateMusicStream(game.Music)
		select {
		case <-ticker.C:
			// game.MoveBlockDown()
		default:
			g.EM.Update()
			g.Movement()
			// game.HandleInput()
			rl.BeginDrawing()
			rl.ClearBackground(rl.DarkBlue)
			// game.Draw()
			g.Render()
			rl.EndDrawing()
			g.FrameNo++
		}
	}
}

func (g *Game) SpawnEnemy() {
	e := g.EM.AddEntity("enemy")
	// generate random scale
	scale := rl.GetRandomValue(20, 30)
	// generate random position
	x := rl.GetRandomValue(scale, int32(rl.GetScreenWidth())-scale)
	y := rl.GetRandomValue(scale, int32(rl.GetScreenHeight())-scale)
	// generate random velocity
	velocity := float32(rl.GetRandomValue(1, 10))

	// generate random shape
	sides := rl.GetRandomValue(3, 10)
	// generate random color
	// generate random rotation
	rotation := float32(rl.GetRandomValue(1, 5))
	tr := CTransform{
		Pos:      Vec2{float32(x), float32(y)},
		Velocity: Vec2{velocity, velocity},
		Scale:    Vec2{float32(scale), float32(scale)},
		Angle:    0, Rotation: rotation}
	e.Components.Transform = tr
	sh := CShape{Sides: sides, Radius: float32(scale)}
	e.Components.Shape = sh

}

func (g *Game) SpawnPlayer() {
	e := g.EM.AddEntity("player")
	tr := CTransform{Pos: Vec2{100, 100}, Velocity: Vec2{0, 0}, Scale: Vec2{1, 1}, Angle: 0, Rotation: 3}
	e.Components.Transform = tr
	sh := CShape{Sides: 4, Radius: 25}
	e.Components.Shape = sh
}

func (g *Game) Render() {
	entities := g.EM.Entities
	for _, e := range entities {
		tr := e.Components.Transform
		sh := e.Components.Shape
		// rl.DrawCircle(int32(tr.Pos.X), int32(tr.Pos.Y), 10, rl.Red)
		// draw polygon shape
		// fmt.Println(sh.Sides)
		rl.DrawPoly(rl.NewVector2(tr.Pos.X, tr.Pos.Y), sh.Sides, sh.Radius, tr.Angle, rl.Green)
	}
}

func (g *Game) Movement() {
	for i := range g.EM.Entities {
		g.EM.Entities[i].Components.Transform.Pos.X += g.EM.Entities[i].Components.Transform.Velocity.X
		g.EM.Entities[i].Components.Transform.Pos.Y += g.EM.Entities[i].Components.Transform.Velocity.Y
		g.EM.Entities[i].Components.Transform.Angle += g.EM.Entities[i].Components.Transform.Rotation
	}
}
