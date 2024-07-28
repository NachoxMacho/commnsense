package handler

import (
	"net/http"

	"github.com/NachoxMacho/commnsense/pkg/opnsense"
	"github.com/NachoxMacho/commnsense/types"
	"github.com/NachoxMacho/commnsense/view/home"
)


func HandleNewHomeIndex(c Config) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {

		data, err := getInterfaceSearchList(c)
		if err != nil {
			return err
		}

		return home.Index(data, "").Render(r.Context(), w)
	}
}

func getInterfaceSearchList(c Config) ([]types.SearchData, error) {
		interfaces, err := opnsense.GetInterfaces(c.OpnSense)
		if err != nil {
			return nil, err
		}


		input := make([]types.SearchData, 0, len(interfaces))
		for _, l := range interfaces {
			d := types.SearchData{
				Checked: false,
				Text:    l.Device,
			}
			input = append(input, d)
		}

		input[0].Checked = true
		return input, nil
}
