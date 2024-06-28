package handler

import (
	"net/http"

	"github.com/NachoxMacho/commnsense/pkg/opnsense/unbound"
	"github.com/NachoxMacho/commnsense/types"
	"github.com/NachoxMacho/commnsense/view/home"
)

var test = []types.SearchData{
	{
		Text:    "Hello",
		Checked: false,
	},
	{
		Text:    "Hello2",
		Checked: false,
	},
	{
		Text:    "Test",
		Checked: true,
	},
}

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	data, err := unbound.GetDHCPLeases()
	if err != nil {
		return err
	}

	input := []types.SearchData{}
	dataLoop: for _, l := range data {
		for _, i := range input {
			if i.Text == l.Interface {
				continue dataLoop
			}
		}
		d := types.SearchData{
			Checked: false,
			Text: l.Interface,
		}
		input = append(input, d)

	}

	input[0].Checked = true

	return home.Index(input, "").Render(r.Context(), w)
}
