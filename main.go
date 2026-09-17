package gradify

import (
	"fmt"
	"strings"
)

var (
	Candy     = []string{"ff6666", "ffcc99", "ff99cc", "cc99ff", "99ccff", "66ccff", "ff99cc", "ff6666"}
	Minty     = []string{"66ffcc", "ccff99", "99ffcc", "99ccff", "cc99ff", "ff99cc", "66ffcc"}
	Nostalgia = []string{"8611ec", "1ede8e", "5744f0"}
	Nectar    = []string{"ffb347", "ff6f91"}

	Error   = []string{"cc0000", "e60000", "ff3333", "ff6666", "ff9999"}
	Success = []string{"006600", "009900", "33cc33", "66ff66", "99ff99"}
	Warning = []string{"cc9900", "e6b800", "ffcc00", "ffdb4d", "fff2b3"}
	Info    = []string{"0055cc", "0077ff", "3399ff", "66b2ff", "99ccff"}

	Default    = []string{"8800fe", "00ffcc"}
	Hollow     = []string{"ff00ff", "ffff00"}
	Sunset     = []string{"ff6d00", "ffd100"}
	Aether     = []string{"8000ff", "0000ff"}
	Blossom    = []string{"ff69b4", "8040ff"}
	Frostbite  = []string{"00b8cc", "cfeef4"}
	Matrix     = []string{"39ff14"}
	Emberlight = []string{"ffb400", "ff6347"}
	Aurora     = []string{"00ffe7", "ff00c8"}
	Glacier    = []string{"74ebd5", "acb6e5"}
	Christmas  = []string{"e34234", "7cfc00"}
)

type Preset struct {
	Name        string
	Description string
	Hex         []string
}

type Color struct {
	R, G, B int
}

func Convert(h string) (c Color, err error) {
	switch len(h) {
	case 6:
		_, err = fmt.Sscanf(h, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(h, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid hex color")
	}
	return
}

func Colorize(text string, r, g, b int) string {
	fg := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return fg + text + "\x1b[0m"
}

func Algo(s, e float64, steps int) []int {
	delta := (e - s) / float64(steps-1)
	colors := []int{int(s)}
	err := 0.0

	for i := 0; i < steps-1; i++ {
		n := float64(colors[i]) + delta
		err = err + (n - float64(int(n)))
		if err >= 0.5 {
			n = n + 1.0
			err = err - 1.0
		}

		colors = append(colors, int(n))
	}
	return colors
}

func InterpolateColor(start, end Color, t float64) Color {
	lerp := func(a, b int, t float64) int {
		return int(float64(a) + t*float64(b-a))
	}
	return Color{
		R: lerp(start.R, end.R, t),
		G: lerp(start.G, end.G, t),
		B: lerp(start.B, end.B, t),
	}
}

func MakeGradient(colors []Color, n int) ([]int, []int, []int) {
	if len(colors) < 2 {
		r := []int{colors[0].R, colors[0].R}
		g := []int{colors[0].G, colors[0].G}
		b := []int{colors[0].B, colors[0].B}
		return r, g, b
	}

	var R, G, B []int
	segments := len(colors) - 1
	for i := 0; i < segments; i++ {
		steps := (n-1)*i/segments + 1
		gradientR := Algo(float64(colors[i].R), float64(colors[i+1].R), steps)
		gradientG := Algo(float64(colors[i].G), float64(colors[i+1].G), steps)
		gradientB := Algo(float64(colors[i].B), float64(colors[i+1].B), steps)
		R = append(R, gradientR...)
		G = append(G, gradientG...)
		B = append(B, gradientB...)
	}
	return R, G, B
}

func Gradient(text string, rgb []string) string {
	text = strings.TrimSpace(text)

	hexValues := strings.TrimSpace(strings.Join(rgb, " "))
	colorHexValues := strings.Split(hexValues, " ")

	colors := make([]Color, len(colorHexValues))
	for i, hex := range colorHexValues {
		colors[i], _ = Convert(hex)
	}

	n := len(text)
	r, g, b := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		// Calculate gradient color for each character
		segment := float64(len(colors) - 1)
		segmentIdx := float64(i) / float64(n-1) * segment
		idx1 := int(segmentIdx)
		idx2 := idx1 + 1

		if idx2 >= len(colors) {
			idx2 = len(colors) - 1
		}

		fraction := segmentIdx - float64(idx1)
		gradR := int(float64(colors[idx1].R)*(1-fraction) + float64(colors[idx2].R)*fraction)
		gradG := int(float64(colors[idx1].G)*(1-fraction) + float64(colors[idx2].G)*fraction)
		gradB := int(float64(colors[idx1].B)*(1-fraction) + float64(colors[idx2].B)*fraction)

		r[i], g[i], b[i] = gradR, gradG, gradB
	}

	coloredText := ""

	for i, t := range text {
		coloredText += Colorize(fmt.Sprint(string(t)), r[i], g[i], b[i])
	}

	return coloredText
}
