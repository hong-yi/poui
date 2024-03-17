package tui

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

func CreateMultiSelectForm(title string, choices []string, selectedValues *[]string) *huh.Form {
	selectionGroup := huh.NewOptions[string](choices...)
	selectionGroup[0] = selectionGroup[0].Selected(true)

	tfeUrlsSelect := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(fmt.Sprintf("%d %s", len(choices), title)).
				Options(selectionGroup...).
				Value(selectedValues),
		))

	return tfeUrlsSelect
}
