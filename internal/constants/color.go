package constants

type ColorNotification int

const (
	ColorSuccess ColorNotification = 0x00ff00
	ColorInfo    ColorNotification = 0x0000ff
	ColorWarning ColorNotification = 0xffff00
	ColorError   ColorNotification = 0xff0000
	ColorDefault ColorNotification = 0xffffff
)
