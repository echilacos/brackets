package brackets

type Color string

const (
	ColorBlack      = Color("\033[1;30m%s\033[0m")
	ColorRed        = Color("\033[1;31m%s\033[0m")
	ColorGreen      = Color("\033[1;32m%s\033[0m")
	ColorGreenBold  = Color("\033[1;32m\033[1m%s\033[0m\033[0m")
	ColorYellow     = Color("\033[1;33m%s\033[0m")
	ColorYellowBold = Color("\033[1;33m\033[1m%s\033[0m\033[0m")
	ColorPurple     = Color("\033[1;34m%s\033[0m")
	ColorMagenta    = Color("\033[1;35m%s\033[0m")
	ColorTeal       = Color("\033[1;36m%s\033[0m")
	ColorWhite      = Color("\033[1;37m%s\033[0m")
)

func (c Color) GetString() string {
	return string(c)
}
