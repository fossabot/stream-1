package trellogithub

import (
	"github.com/merico-dev/stream/internal/pkg/log"
)

// Reinstall remove and set up trello-github-integ workflows.
func Reinstall(options *map[string]interface{}) (bool, error) {
	gis, err := NewTrelloGithub(options)
	if err != nil {
		return false, err
	}

	api := gis.GetApi()
	log.Infof("api is %s", api.Name)
	ws := defaultWorkflows.GetWorkflowByNameVersionTypeString(api.Name)

	for _, w := range ws {
		err := gis.DeleteWorkflow(w)
		if err != nil {
			return false, err
		}

		if err := gis.renderTemplate(w); err != nil {
			return false, err
		}
		err = gis.AddWorkflow(w)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
