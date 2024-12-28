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

	rl.SetTargetFPS(60)
	g.SpawnPlayer()
	for i := 0; i < 5; i++ {
		g.SpawnEnemy()
	}

	enemySpwanTicker := time.NewTicker(5 * time.Second)
	for !rl.WindowShouldClose() {
		select {
		case <-enemySpwanTicker.C:
			// game.MoveBlockDown()
			g.SpawnEnemy()
		default:
			g.EM.Update()
			g.HandleInput()
			g.Movement()
			g.Collision()
			rl.BeginDrawing()
			rl.ClearBackground(rl.DarkGray)
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
	vx := float32(rl.GetRandomValue(-3, 3))
	vy := float32(rl.GetRandomValue(-3, 3))

	// generate random shape
	sides := rl.GetRandomValue(3, 10)
	// generate random color
	// generate random rotation
	rotation := float32(rl.GetRandomValue(1, 5))
	tr := CTransform{
		Pos:       Vec2{float32(x), float32(y)},
		Direction: Vec2{vx, vy},
		Scale:     Vec2{float32(scale), float32(scale)},
		Angle:     0, Rotation: rotation}
	e.Components.Transform = tr
	cr := rl.GetRandomValue(0, 255)
	cg := rl.GetRandomValue(0, 255)
	cb := rl.GetRandomValue(0, 255)
	br := rl.GetRandomValue(0, 255)
	bg := rl.GetRandomValue(0, 255)
	bb := rl.GetRandomValue(0, 255)

	r := rl.GetRandomValue(15, 25)
	sh := CShape{Sides: sides, Radius: float32(r), cr: uint8(cr), cg: uint8(cg), cb: uint8(cb), br: uint8(br), bg: uint8(bg), bb: uint8(bb)}
	e.Components.Shape = sh
	// life := CLifespan{Lifespan: 100, Remaining: 50}
	// e.Components.Lifespan = &lifes
}

func (g *Game) SpawnPlayer() {
	e := g.EM.AddEntity("player")
	tr := CTransform{Pos: Vec2{100, 100}, Direction: Vec2{3, 3}, Scale: Vec2{1, 1}, Angle: 0, Rotation: 0, Speed: 3}
	e.Components.Transform = tr
	sh := CShape{Sides: 8, Radius: 25, cr: 50, cg: 255, cb: 0}
	e.Components.Shape = sh
	in := CInput{Up: false, Down: false, Left: false, Right: false, Shoot: false}
	e.Components.Input = &in
}

// SpawnBullet
func (g *Game) SpawnBullet(p, t Vec2) {
	e := g.EM.AddEntity("bullet")
	dir := Vec2{t.X - p.X, t.Y - p.Y}
	dir.Normalize()
	tr := CTransform{Pos: p, Direction: dir, Scale: Vec2{1, 1}, Angle: 0, Rotation: 0, Speed: 6}
	e.Components.Transform = tr
	sh := CShape{Sides: 20, Radius: 5, cr: 255, cg: 255, cb: 255}
	e.Components.Shape = sh
	lf := CLifespan{Lifespan: 60, Remaining: 60}
	e.Components.Lifespan = &lf
}

func (g *Game) Render() {
	entities := g.EM.Entities
	for _, e := range entities {
		tr := e.Components.Transform
		sh := e.Components.Shape
		alfa := uint8(255)
		if e.Components.Lifespan != nil {
			life := e.Components.Lifespan
			alfa = uint8(float32(life.Remaining) / float32(life.Lifespan) * 256)
		}
		color := rl.Color{R: e.Components.Shape.cr, G: e.Components.Shape.cg, B: e.Components.Shape.cb, A: alfa}
		rl.DrawPoly(rl.NewVector2(tr.Pos.X, tr.Pos.Y), sh.Sides, sh.Radius, tr.Angle, color)
		bClolor := rl.Color{R: e.Components.Shape.br, G: e.Components.Shape.bg, B: e.Components.Shape.bb, A: alfa}
		rl.DrawPolyLinesEx(rl.NewVector2(tr.Pos.X, tr.Pos.Y), sh.Sides, sh.Radius, tr.Angle, 3, bClolor)
	}
}

func (g *Game) Movement() {
	p := g.EM.GetEntityByTag("player")
	disp := Vec2{0, 0}
	if p[0].Components.Input != nil {
		if p[0].Components.Input.Left {
			disp.X -= p[0].Components.Transform.Direction.X
		} else if p[0].Components.Input.Right {
			disp.X += p[0].Components.Transform.Direction.X
		}
		if p[0].Components.Input.Up {
			disp.Y -= p[0].Components.Transform.Direction.Y
		} else if p[0].Components.Input.Down {
			disp.Y += p[0].Components.Transform.Direction.Y
		}
		if disp.X != 0 || disp.Y != 0 {
			disp.Normalize()
			disp.Mul(p[0].Components.Transform.Speed)
			p[0].Components.Transform.Pos.Add(disp)
			if p[0].Components.Transform.Pos.X < p[0].Components.Shape.Radius {
				p[0].Components.Transform.Pos.X = p[0].Components.Shape.Radius
			}
			if p[0].Components.Transform.Pos.X > float32(rl.GetScreenWidth())-p[0].Components.Shape.Radius {
				p[0].Components.Transform.Pos.X = float32(rl.GetScreenWidth()) - p[0].Components.Shape.Radius
			}
			if p[0].Components.Transform.Pos.Y < p[0].Components.Shape.Radius {
				p[0].Components.Transform.Pos.Y = p[0].Components.Shape.Radius
			}
			if p[0].Components.Transform.Pos.Y > float32(rl.GetScreenHeight())-p[0].Components.Shape.Radius {
				p[0].Components.Transform.Pos.Y = float32(rl.GetScreenHeight()) - p[0].Components.Shape.Radius
			}
		}
	}
	e := g.EM.GetEntityByTag("enemy")
	for i := range e {
		e[i].Components.Transform.Angle += e[i].Components.Transform.Rotation
		e[i].Components.Transform.Pos.Add(e[i].Components.Transform.Direction)

		if e[i].Components.Transform.Pos.X < e[i].Components.Shape.Radius ||
			e[i].Components.Transform.Pos.X > float32(rl.GetScreenWidth())-e[i].Components.Shape.Radius {
			e[i].Components.Transform.Direction.X = -e[i].Components.Transform.Direction.X
		}

		if e[i].Components.Transform.Pos.Y < e[i].Components.Shape.Radius ||
			e[i].Components.Transform.Pos.Y > float32(rl.GetScreenHeight())-e[i].Components.Shape.Radius {
			e[i].Components.Transform.Direction.Y = -e[i].Components.Transform.Direction.Y
		}
	}
	b := g.EM.GetEntityByTag("bullet")
	for i := range b {
		// fmt.Println("Bullet", b[i].Components.Transform.Pos)
		vel := b[i].Components.Transform.Direction
		vel.Mul(b[i].Components.Transform.Speed)
		b[i].Components.Transform.Pos.Add(vel)
		if b[i].Components.Lifespan != nil {
			b[i].Components.Lifespan.Remaining--
			if b[i].Components.Lifespan.Remaining <= 0 {
				b[i].alive = false
			}
		}
	}
}

func (g *Game) Collision() {
	enemy := g.EM.GetEntityByTag("enemy")
	bullets := g.EM.GetEntityByTag("bullet")
	for i := 0; i < len(bullets); i++ {
		for j := 0; j < len(enemy); j++ {
			if bullets[i].alive && enemy[j].alive {
				tr1 := bullets[i].Components.Transform
				tr2 := enemy[j].Components.Transform
				sh1 := bullets[i].Components.Shape
				sh2 := enemy[j].Components.Shape
				dist := Vec2{tr1.Pos.X - tr2.Pos.X, tr1.Pos.Y - tr2.Pos.Y}
				// dist.Normalize()
				// dist.Mul(sh1.Radius + sh2.Radius)
				if dist.Length() <= sh1.Radius+sh2.Radius {
					bullets[i].alive = false
					enemy[j].alive = false
				}
			}
		}
	}
}

func (g *Game) HandleInput() {
	p := g.EM.GetEntityByTag("player")

	pos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		dir := Vec2{pos.X, pos.Y}
		g.SpawnBullet(p[0].Components.Transform.Pos, dir)
	}

	if rl.IsKeyDown(rl.KeyA) {
		p[0].Components.Input.Left = true
	}
	if rl.IsKeyDown(rl.KeyD) {
		p[0].Components.Input.Right = true
	}
	if rl.IsKeyDown(rl.KeyW) {
		p[0].Components.Input.Up = true
	}
	if rl.IsKeyDown(rl.KeyS) {
		p[0].Components.Input.Down = true
	}
	if rl.IsKeyUp(rl.KeyA) {
		p[0].Components.Input.Left = false
	}
	if rl.IsKeyUp(rl.KeyD) {
		p[0].Components.Input.Right = false
	}
	if rl.IsKeyUp(rl.KeyW) {
		p[0].Components.Input.Up = false
	}
	if rl.IsKeyUp(rl.KeyS) {
		p[0].Components.Input.Down = false
	}

}
