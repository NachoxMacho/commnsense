package handler

import (
	"net/http"

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
	return home.Index(test, "").Render(r.Context(), w)
}
