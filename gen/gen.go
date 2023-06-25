package gen

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed templates/apache.txt
var apache string

//go:embed templates/gpl-3.0.txt
var gpl string

//go:embed templates/mit.txt
var mit string

//go:embed templates/unlicense.txt
var unlicense string

type LicenseData struct {
	Year   string
	Author string
}

func GetLicense(name string) (string, error) {
	var licenseTemplate string
	switch name {
	case "apache":
		licenseTemplate = apache
	case "gpl":
		licenseTemplate = gpl
	case "mit":
		licenseTemplate = mit
	case "unlicense":
		licenseTemplate = unlicense
	}

	licenseData := &LicenseData{
		Year:   "2023",
		Author: "Faiz Mokhtar",
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("license").Parse(licenseTemplate))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, licenseData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
