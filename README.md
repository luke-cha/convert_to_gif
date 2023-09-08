# convert_to_gif
Convert MOV or MP4 to gif

### Support
Only for Mac

### Prerequisites
```bash
brew install ffmpeg gifsicle
```

### Build
```
go build
```

### Setup
```bin
sudo mv convert_gif /usr/local/go/bin
```

### Usage
```bash
convert_gif -file="Screen Recording 2023-09-08 at 4.24.23 PM.mov" -speed=0.5 -width=1200
```