package dico

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var err error
var application Application

// GetDico chargement du dictionnaire de l'application
func GetDico() Application {
	if application.Title == "" {
		log.Println("Chargement du dictionnaire...")
		jsonFile, err := ioutil.ReadFile("resources/dico.json")
		if err != nil {
			log.Print("Unable to read dico file", err)
			return Application{}
		}
		err = json.Unmarshal(jsonFile, &application)
		if err != nil {
			log.Print("Invalid JSON file", err)
			return Application{}
		}
	}
	return application
}

// Application le dictionnaire de l'application
type Application struct {
	Title         string            `json:"title"`
	Info          string            `json:"info"`
	IconFile      string            `json:"icon_file"`
	Basename      string            `json:"basename"`
	DataDirectory string            `json:"data_directory"`
	Constants     map[string]string `json:"constants"`
	Tables        []Table           `json:"tables"`
}

// Table Table de l'application
type Table struct {
	ID       string    `json:"id"`
	Key      string    `json:"key"`
	Elements []Element `json:"elements"`
	Views    []View    `json:"views"`
	Forms    []Form    `json:"forms"`
}

// View Vue d'une table
type View struct {
	ID         string    `json:"id"`
	Deletable  bool      `json:"deletable"`
	Elements   []Element `json:"elements"`
	FormAdd    string    `json:"form_add"`
	FormEdit   string    `json:"form_edit"`
	Hide       bool      `json:"hide"`
	Limit      int       `json:"limit"`
	OrderBy    string    `json:"order_by"`
	SQLFooter  string    `json:"sql_footer"`
	SQLWhere   string    `json:"sql_where"`
	Searchable bool      `json:"searchable"`
	Title      string    `json:"title"`
}

// Form formulaire
type Form struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Elements []Element `json:"elements"`
}

// Element Rubrique de l'application
type Element struct {
	ID            string            `json:"id"`
	Args          map[string]string `json:"args"`
	BackgroundSQL string            `json:"background_sql"`
	ColEditable   bool              `json:"col_editable"`
	ColWith       int               `json:"col_width"`
	ColorSQL      string            `json:"color_sql"`
	Default       string            `json:"default"`
	DefaultSQL    string            `json:"default_sql"`
	Display       string            `json:"display"`
	Editable      bool              `json:"editable"`
	Formulas      string            `json:"formulas"`
	Height        int               `json:"height"`
	Help          string            `json:"help"`
	HelpSQL       string            `json:"help_sql"`
	Items         map[string]string `json:"items"`
	Jointure      map[string]string `json:"jointure"`
	LabelLong     string            `json:"label_long"`
	LabelShort    string            `json:"label_short"`
	Params        map[string]string `json:"params"`
	ReadOnly      bool              `json:"read_only"`
	Refresh       bool              `json:"refresh"`
	Required      bool              `json:"required"`
	SQLText       string            `json:"sql_text"`
	Searchable    bool              `json:"searchable"`
	Sortable      bool              `json:"sortable"`
	Type          string            `json:"type"`
}
