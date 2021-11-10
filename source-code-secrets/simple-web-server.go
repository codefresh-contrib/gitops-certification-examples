package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/ini.v1"
)

type configurationListHandler struct {
	appMode        string
	paypalURL      string
	paypalCertPath string
	dbCon          string
	dbUser         string
	dbPassword     string
}

func (h *configurationListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>I am a simple application looking for my secrets.</h1> <h2>My properties are:</h2><ul>")
	fmt.Fprintf(w, "<li>app_mode: "+h.appMode+"</li>")
	fmt.Fprintf(w, "<li>paypal_url: "+h.paypalURL+"</li>")
	fmt.Fprintf(w, "<li>paypal_cert: "+h.paypalCertPath+"</li>")
	fmt.Fprintf(w, "<li>db_con: "+h.dbCon+"</li>")
	fmt.Fprintf(w, "<li>db_user: "+h.dbUser+"</li>")
	fmt.Fprintf(w, "<li>db_password: "+h.dbPassword+"</li>")
	fmt.Fprintf(w, "</ul>")

	fmt.Fprintf(w, "<h2> Mysql URL </h2>")
	fmt.Fprintf(w, "<pre> %s</pre>", readFileToString(h.dbCon))

	fmt.Fprintf(w, "<h2> Mysql username </h2>")
	fmt.Fprintf(w, "<pre> %s</pre>", readFileToString(h.dbUser))

	fmt.Fprintf(w, "<h2> Mysql password </h2>")
	fmt.Fprintf(w, "<pre> %s</pre>", readFileToString(h.dbPassword))

	fmt.Fprintf(w, "<h2> Paypal cert </h2>")
	fmt.Fprintf(w, "<pre> %s</pre>", readFileToString(h.paypalCertPath))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")

}

func main() {

	cfg, err := ini.LooseLoad("/config/settings.ini", "settings.ini")
	if err != nil {
		fmt.Printf("Failed to read configuration file: %v", err)
	}

	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())

	clh := configurationListHandler{}
	clh.appMode = cfg.Section("").Key("app_mode").String()
	clh.paypalURL = cfg.Section("paypal").Key("paypal_url").String()
	clh.paypalCertPath = cfg.Section("paypal").Key("paypal_cert").String()
	clh.dbCon = cfg.Section("mysql").Key("db_con").String()
	clh.dbUser = cfg.Section("mysql").Key("db_user").String()
	clh.dbPassword = cfg.Section("mysql").Key("db_password").String()

	fmt.Println("Simple web server is starting now on port 8080...")

	http.Handle("/", &clh)
	http.HandleFunc("/health", healthHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server at port 8080: %v", err)
	}
}

func readFileToString(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "Could not read " + filename
	}
	return string(data)
}
