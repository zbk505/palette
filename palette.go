package palette

import (
	"fmt"
	"math"
)

type HSV struct {
	H    uint8
	S, V float64
}
type RGB struct {
	R, G, B uint8
}

// Converts an RGB color into an HSV color
func NewHSV(rgb RGB) (*HSV, error) {
	var R, G, B = rgb.R, rgb.G, rgb.B

	Ri := float64(R) / (255)
	Gi := float64(G) / (255)
	Bi := float64(B) / (255)

	RGB := []float64{Ri, Gi, Bi}

	Cmax, err := max(RGB)
	Cmin, err := min(RGB)
	if err != nil {
		return nil, fmt.Errorf("Error: invalid color supplied")
	}

	delta := Cmax - Cmin

	H, err := getHue(Ri, Gi, Bi, Cmax, delta)
	S := getSaturation(Cmax, delta)
	V := Cmax * 100

	if err != nil {
		return nil, err
	}
	return &HSV{uint8(math.Round(H)), S, V}, nil
}

// max returns the maximum value within a supplied array of float64
func max(arr []float64) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Supplied array is empty")
	}

	max := float64(math.MinInt)
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max, nil
}

// min returns the minimum value within a supplied array of float64
func min(arr []float64) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("Supplied array is empty")
	}

	min := float64(math.MaxInt)
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, nil
}
func getHue(R float64, G float64, B float64, Cmax float64, delta float64) (float64, error) {
	H := float64(0)
	if delta == 0 {
		return 0, nil
	} else if Cmax == R {
		H = 60 * math.Mod(((G-B)/delta), 6)
		return H, nil
	} else if Cmax == G {
		H = 60 * (((B - R) / delta) + 2)
		return H, nil
	} else if Cmax == B {
		H = 60 * (((R - G) / delta) + 4)
		return H, nil
	}
	return H, fmt.Errorf("Something bad happened while getting the Hue value")
}
func getSaturation(Cmax float64, delta float64) float64 {
	if Cmax == 0 {
		return float64(0)
	} else {
		return float64(delta/Cmax) * 100
	}
}

func main() {

	var c RGB = RGB{255, 30, 8}
	fmt.Printf("RGB Color: R: %d G: %d B: %d\n", c.R, c.G, c.B)
	hsv, err := NewHSV(c)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("HSV Color: H: %d S: %.1f V: %.1f\n", hsv.H, hsv.S, hsv.V)
}
