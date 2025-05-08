package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	running bool = true
	bgColor = rl.NewColor(19,199,212, 0)
	grass rl.Texture2D
	player rl.Texture2D
	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerSpeed float32 = 3
	musicPaused bool
	music rl.Music
	cam rl.Camera2D
)
const(
	screenWidth = 700
	screenHeight = 600

)

func drawScene(){
	rl.ClearBackground(bgColor)
	rl.DrawTexture(grass, 100, 50, rl.White)
	rl.DrawTexturePro(player, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}
func input(){
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp){
		playerDest.Y-=playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown){
		playerDest.Y+=playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight){
		playerDest.X+=playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft){
		playerDest.X-=playerSpeed
	}
	if rl.IsKeyPressed(rl.KeyQ){
		musicPaused = !musicPaused
	}
}
func update(){
	running = !rl.WindowShouldClose()

	rl.UpdateMusicStream(music)
	if musicPaused{
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
}
func render(){
	rl.BeginDrawing()
	rl.BeginMode2D(cam)
	drawScene()
	rl.EndMode2D()
	rl.EndDrawing()
}

func init(){
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	
	grass = rl.LoadTexture("./res/Tilesets/Grass.png")
	player = rl.LoadTexture("./res/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0,0,48,48)
	playerDest = rl.NewRectangle(200,200,100,100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./res/music.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenWidth/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.0)
}

func quit (){
	rl.UnloadTexture(grass)
	rl.UnloadTexture(player)
	rl.CloseAudioDevice()
}

func main() {
	defer rl.CloseWindow()

	for running  {
		input()
		update()
		render()
	}

	quit()
}
