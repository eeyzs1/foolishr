package sample

import (
	"math/rand"
	"time"
	// "sync"

	"github.com/google/uuid"
	"foolishr/src/pb"
)

var (
	randInst *rand.Rand  // 单例实例 
	// once     sync.Once   // 保证初始化仅执行一次 
)

// func GetRandInstance() *rand.Rand {
//     once.Do(func() {
//         source := rand.NewSource(time.Now().UnixNano())
//         randInst = rand.New(source)
//     })
//     return randInst 
// }

// 初始化单例,由于go的init只会在pkg import的之后被调用一次，此后所有import皆不再初始化，所以不sync
func init() {
		source := rand.NewSource(time.Now().UnixNano())
		randInst = rand.New(source)
}

func randomStringFrandomSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[randInst.Intn(n)]
}

func randomBool() bool {
	return randInst.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + randInst.Int()%(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + randInst.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + randInst.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch randInst.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if randInst.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomCPUBrand() string {
	return randomStringFrandomSet("Intel", "AMD")
}

func randomCPUName(br string) string {
	if br == "Intel" {
		return randomStringFrandomSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}

	return randomStringFrandomSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUBrand() string {
	return randomStringFrandomSet("Nvidia", "AMD")
}

func randomGPUName(br string) string {
	if br == "Nvidia" {
		return randomStringFrandomSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFrandomSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomLaptopBrand() string {
	return randomStringFrandomSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(br string) string {
	switch br {
	case "Apple":
		return randomStringFrandomSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFrandomSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFrandomSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
