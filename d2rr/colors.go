package d2rr

type Color string

const white1 Color = "ÿc0"
const white2 Color = "ÿc="
const gray1 Color = "ÿc5"
const gray2 Color = "ÿcK"
const gray3 Color = "ÿcI"
const black1 Color = "ÿc6"
const black2 Color = "ÿcM"
const lightred Color = "ÿcE"
const red1 Color = "ÿc1"
const red2 Color = "ÿcU"
const darkred Color = "ÿcS"
const orange1 Color = "ÿc@"
const orange2 Color = "ÿc8"
const orange3 Color = "ÿcJ"
const orange4 Color = "ÿcL"
const lightgold1 Color = "ÿc7"
const lightgold2 Color = "ÿcH"
const gold1 Color = "ÿc4"
const gold2 Color = "ÿcD"
const yellow1 Color = "ÿc9"
const yellow2 Color = "ÿcR"
const green1 Color = "ÿc2"
const green2 Color = "ÿcQ"
const green3 Color = "ÿcC"
const green4 Color = "ÿc<"
const darkgreen1 Color = "ÿcA"
const darkgreen2 Color = "ÿc:"
const turquoise Color = "ÿcN"
const skyblue Color = "ÿcT"
const lightblue1 Color = "ÿcF"
const lightblue2 Color = "ÿcP"
const blue1 Color = "ÿc3"
const blue2 Color = "ÿcB"
const lightpink Color = "ÿcG"
const pink Color = "ÿcO"
const purple Color = "ÿc;"

func ColorCode(color string) Color {
	switch c := color; c {
	case "white1":
		return white1
	case "white2":
		return white1
	case "gray1":
		return gray1
	case "gray2":
		return gray2
	case "gray3":
		return gray3
	case "black1":
		return black1
	case "black2":
		return black2
	case "lightred":
		return lightred
	case "red1":
		return red1
	case "red2":
		return red2
	case "darkred":
		return darkred
	case "orange1":
		return orange1
	case "orange2":
		return orange2
	case "orange3":
		return orange3
	case "orange4":
		return orange4
	case "lightgold1":
		return lightgold1
	case "lightgold2":
		return lightgold2
	case "gold1":
		return gold1
	case "gold2":
		return gold2
	case "yellow1":
		return yellow1
	case "yellow2":
		return yellow2
	case "green1":
		return green1
	case "green2":
		return green2
	case "green3":
		return green3
	case "green4":
		return green4
	case "darkgreen1":
		return darkgreen1
	case "darkgreen2":
		return darkgreen2
	case "turquoise":
		return turquoise
	case "skyblue":
		return skyblue
	case "lightblue1":
		return lightblue1
	case "lightblue2":
		return lightblue2
	case "blue1":
		return blue1
	case "blue2":
		return blue2
	case "lightpink":
		return lightpink
	case "pink":
		return pink
	case "purple":
		return purple
	default:
		return ""
	}
}

func (c Color) String() string {
	return string(c)
}
