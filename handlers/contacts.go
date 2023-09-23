package handlers

import (
	"html/template"
	"net/http"
	"strings"
)

type Contacts = struct {
	Name  string
	Email string
}

type ContactsTemplateData = struct {
	Contacts []Contacts
	Q        string
}

var contactsDb = []Contacts{{"Seagin", "seagin@me.com"}, {"Sausage", "sausage@me.com"}, {"Snow", "snow@me.com"}}

func filterContacts(q string) []Contacts {
	if len(q) > 0 {
		filteredContacts := []Contacts{}

		for _, c := range contactsDb {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(q)) {
				filteredContacts = append(filteredContacts, c)
			}
		}

		return filteredContacts
	}
	return contactsDb
}

func RenderContactsPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl", "templates/contacts.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := r.URL.Query().Get("q")
	contacts := filterContacts(q)
	data := ContactsTemplateData{contacts, q}

	err = t.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleContact(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	// as of now there is only get
	default:
		RenderContactsPage(w, r)
	}
}
