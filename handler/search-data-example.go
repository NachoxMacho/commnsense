package handler

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/NachoxMacho/commnsense/types"
	"github.com/NachoxMacho/commnsense/view/ui"
)

func HandleSearchData(c Config) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {

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

		data, err := getInterfaceSearchList(c)
		if err != nil {
			return err
		}

		renderData := make([]types.SearchData, 0, len(data))
		for _, d := range data {
			if strings.Contains(strings.ToLower(d.Text), strings.ToLower(searchString)) {
				renderData = append(renderData, d)
			}
		}

		return ui.DropDownContent(renderData, false).Render(r.Context(), w)
	}
}

func HandleDropDown(c Config) httpHandler {

	return func(w http.ResponseWriter, r *http.Request) error {
		data, err := getInterfaceSearchList(c)
		if err != nil {
			return err
		}
		return ui.SearchDropDown(data, r.URL.Query().Get("selected")).Render(r.Context(), w)
	}
}
