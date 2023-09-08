package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		speed    = flag.Float64("speed", 0.3125, "playback speed")
		width    = flag.Int("width", 640, "output width")
		filename = flag.String("file", "", "input file")
	)

	flag.Parse()

	if *filename == "" {
		fmt.Println("Please provide an input file using the -file flag.")
		os.Exit(1)
	}

	speedStr := fmt.Sprintf("setpts=%f*PTS", *speed)
	resizeStr := fmt.Sprintf("fps=10,scale=%d:-1:flags=lanczos,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse", *width)

	cmd1 := exec.Command("ffmpeg", "-i", *filename, "-filter:v", speedStr, "output.mp4")
	cmd1.Run()

	cmd2 := exec.Command("ffmpeg", "-i", "output.mp4", "-vf", resizeStr, "-loop", "0", "output.gif")
	cmd2.Run()

	cmd3 := exec.Command("gifsicle", "-i", "output.gif", "--optimize=3", "--colors", "256", "-O3", "--lossy=80", "-o", "final_output.gif")
	cmd3.Run()

	os.Remove("output.mp4")
	os.Remove("output.gif")
}
