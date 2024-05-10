package log

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func LogError(msg string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	fmt.Fprintln(os.Stderr, style.Render(msg))
}
