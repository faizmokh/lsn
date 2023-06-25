package gen

import (
	"bytes"
	_ "embed"
	"errors"
	"text/template"
)

//go:embed templates/agpl.txt
var agpl string

//go:embed templates/apache.txt
var apache string

//go:embed templates/boost.txt
var boost string

//go:embed templates/gpl-3.0.txt
var gpl string

//go:embed templates/lgpl.txt
var lgpl string

//go:embed templates/mit.txt
var mit string

//go:embed templates/mozilla.txt
var mozilla string

//go:embed templates/unlicense.txt
var unlicense string

type LicenseData struct {
	Year   int16
	Author string
}

func GetLicense(name string, data *LicenseData) (string, error) {
	var licenseTemplate string
	licenseData := *data
	switch name {
	case "agpl":
		if licenseData.Author == "" {
			return "", errors.New("GNU AGPLv3 license require --author flag to be set")
		}
		licenseTemplate = agpl
	case "apache":
		if licenseData.Author == "" {
			return "", errors.New("apache license 2.0 require --author flag to be set")
		}
		licenseTemplate = apache
	case "boost":
		licenseTemplate = boost
	case "gpl":
		if licenseData.Author == "" {
			return "", errors.New("GNU GPLv3 license require --author flag to be set")
		}
		licenseTemplate = gpl
	case "lgpl":
		licenseTemplate = lgpl
	case "mit":
		if licenseData.Author == "" {
			return "", errors.New("MIT license require --author flag to be set")
		}
		licenseTemplate = mit
	case "mozilla":
		licenseTemplate = mozilla
	case "unlicense":
		licenseTemplate = unlicense
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("license").Parse(licenseTemplate))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, licenseData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
