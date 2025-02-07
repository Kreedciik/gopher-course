package handler

import (
	"farmish/model"
)

func (h *Handler) Dashboard(request string, response *model.Dashboard) error {
	dashboard, err := h.services.Dashboard.GetDashboardData()
	if err != nil {
		return err
	}
	*response = dashboard
	return nil
}
