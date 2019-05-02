package core

import (
	"fmt"
	"os"

	"git.eaciitapp.com/sebar/knot"

	tk "github.com/eaciit/toolkit"
	"github.com/spf13/viper"
)

var configuration *viper.Viper
var env string
var app tk.M
var api tk.M
var public tk.M
var database tk.M
var isDev bool

// Environment = get config environment
func Environment() string {
	return env
}

// IsDev = Environment Development
func IsDev() bool {
	return isDev
}

// IsProd = Environment Production
func IsProd() bool {
	return !isDev
}

// App = get config app
func App() tk.M {
	return app
}

// API = get config api
func API() tk.M {
	return api
}

// Public = get config public
func Public() tk.M {
	return public
}

// ViewsPath = Get path of view's files
func ViewsPath(view string) string {
	return func() string {
		d, _ := os.Getwd()
		v := ""
		if IsDev() {
			v = d + "/views/" + view
		} else {
			v = d + "/views/dist/" + view
		}
		return v
	}()
}

// AssetsPath = Get path of asset's files
func AssetsPath(asset string) string {
	return func() string {
		d, _ := os.Getwd()

		return d + "/assets/" + asset
	}()
}

// FullAddress = Get full address with protocol
func FullAddress(port string) string {
	portFix := API().GetString("port")
	if port != "" {
		portFix = port
	}
	address := App().GetString("host") + ":" + portFix
	return "http://" + address
}

// ReadConfig = for read json config
func ReadConfig() (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigType("json")
	conf.AddConfigPath("./config")
	conf.SetConfigName("config")
	err := conf.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to read configuration. Error : %s", err.Error())
	}
	return conf, nil
}

// Configuration = set configuration
func Configuration() *viper.Viper {
	if configuration != nil {
		return configuration
	}
	configuration, err := ReadConfig()
	if err != nil {
		os.Exit(0)
	}
	env = configuration.GetString("env")
	app, _ = tk.ToM(configuration.Get("app"))
	api, _ = tk.ToM(configuration.Get("api"))
	public, _ = tk.ToM(configuration.Get("public"))
	database, _ = tk.ToM(configuration.Get("database"))
	isDev = env == "dev" || env == "development"

	return configuration
}

// ResponseJSON = Set response json
func ResponseJSON(status, message string, data interface{}) tk.M {
	return tk.M{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

// ResponseJSONError = Set response json for error
func ResponseJSONError(err string) tk.M {
	return ResponseJSON("nok", err, nil)
}

// ResponseJSONOK = Set response json for ok
func ResponseJSONOK(data interface{}) tk.M {
	return ResponseJSON("ok", "", data)
}

// GetMethod = Get method request
func GetMethod(k *knot.WebContext) string {
	m := k.Request.Method
	return m
}

// GetParam = Get param value from query string
func GetParam(k *knot.WebContext, key string) string {
	return k.Request.URL.Query().Get(key)
}
