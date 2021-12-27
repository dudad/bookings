package handlers

import (
	"net/http"

	"github.com/dudad/bookings/pkg/config"
	"github.com/dudad/bookings/pkg/models"
	"github.com/dudad/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	aMap := map[string]string{}
	aMap["test"] = "Some new data"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	aMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: aMap})
}
