package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/NachoxMacho/commnsense/types"
	"github.com/NachoxMacho/commnsense/view/ui"
)

func HandleSearchData(w http.ResponseWriter, r *http.Request) error {

	searchString := ""
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		u, err := url.ParseQuery(string(b))
		if err != nil {
			return err
		}
		searchString = u.Get("search")
	}

	renderData := make([]types.SearchData, 0, len(test))
	for _, d := range test {
		if strings.Contains(strings.ToLower(d.Text), strings.ToLower(searchString)) {
			renderData = append(renderData, d)
		}
	}

	return ui.DropDownContent(renderData, false).Render(r.Context(), w)
}

func HandleDropDown(w http.ResponseWriter, r *http.Request) error {
	return ui.SearchDropDown(test, r.URL.Query().Get("selected")).Render(r.Context(), w)
}
