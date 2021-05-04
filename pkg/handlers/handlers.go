package handlers

import (
	"github.com/ho3einTry/bookings/pkg/Models"
	"github.com/ho3einTry/bookings/pkg/config"
	"github.com/ho3einTry/bookings/pkg/render"
	"net/http"
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

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.SessionManger.Put(r.Context(), "remote_ip", remoteIP)

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Home Page, again."

	//send the data to the template
	render.RenderTemplate(w, "home.page.gohtml", &Models.TemplateDate{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIp := m.App.SessionManger.GetString(r.Context(), "remote_ip")
	stringMap := map[string]string{
		"test":      "Hello About Page, again.",
		"remote_ip": remoteIp,
	}

	render.RenderTemplate(w, "about.page.gohtml", &Models.TemplateDate{
		StringMap: stringMap,
	})
}
