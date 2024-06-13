package handler

import (
	"net/http"

	"github.com/NachoxMacho/commnsense/types"
	"github.com/NachoxMacho/commnsense/view/home"
)


func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	test := []types.SearchData{
		{
			Text: "Hello",
			Checked: false,
		},
	}
	return home.Index(test).Render(r.Context(), w)
}
