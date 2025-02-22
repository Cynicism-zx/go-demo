package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).PaddingTop(2).PaddingLeft(4).Width(22)

func main() {
	fmt.Println(style.Render("Hello, kitty."))
}
