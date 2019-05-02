package core

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"git.eaciitapp.com/sebar/knot"
)

// Mux = For handle mux
type Mux struct {
	http.ServeMux
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	getPath := r.URL.Path
	isDev := Configuration().GetString("environment") == "dev"

	publicPrefix := Configuration().GetString("public.prefix")
	publicPort := Configuration().GetString("public.port")

	apiPrefix := Configuration().GetString("api.prefix")

	if isDev {
		if strings.HasPrefix(getPath, publicPrefix) {
			publicURL, err := url.Parse(FullAddress(publicPort))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			httputil.NewSingleHostReverseProxy(publicURL).ServeHTTP(w, r)
			return
		}
	} else {
		if strings.HasPrefix(getPath, publicPrefix) {
			basePath, _ := os.Getwd()
			vueDistPath := filepath.Join(basePath, "views", "dist")

			if filepath.Ext(getPath) == "" {
				indexPath := filepath.Join(vueDistPath, "index.html")
				http.ServeFile(w, r, indexPath)
				return
			}

			http.StripPrefix(publicPrefix, http.FileServer(http.Dir(vueDistPath))).ServeHTTP(w, r)
			return
		}
	}

	// API
	if strings.HasPrefix(getPath, apiPrefix) {
		m.ServeMux.ServeHTTP(w, r)
		return
	}
	if getPath == "/" || getPath == "" {
		http.Redirect(w, r, publicPrefix, http.StatusTemporaryRedirect)
		return
	}
}

// HandleFunc = Standart handle with func(w http.ResponseWriter, r *http.Request)
func (m *Mux) HandleFunc(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	apiPrefix := Configuration().GetString("api.prefix")
	routeWithPrefix := apiPrefix + route
	m.ServeMux.HandleFunc(routeWithPrefix, handler)
}

// HandleFuncForKnot = Handle for knot with func(ctx *knot.WebContext)
func (m *Mux) HandleFuncForKnot(route string, ctxKnot *knot.WebContext) *knot.Application {
	app := knot.NewApp()
	// apiPrefix := Configuration().GetString("api.prefix")
	// routeWithPrefix := apiPrefix + route
	// m.ServeMux.HandleFunc(routeWithPrefix, func())
	// 	myApp := knot.NewApp()
	// 	myApp.Static("static", core.ViewsPath(""))
	// 	myApp.Static("assets", core.AssetsPath(""))
	// 	myApp.SetViewsPath(core.ViewsPath(""))
	// 	myApp.Register(&Public{}, "")
	// 	myApp.AddRoute("hello", func(ctx *knot.WebContext) {
	// 		ctx.Write([]byte("Hi from public => hello!"), http.StatusOK)
	// 	})
	return app
}
