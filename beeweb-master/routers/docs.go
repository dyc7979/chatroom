// Copyright 2013 Beego Web authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

import (
	"strings"

	"github.com/beego/beeweb/models"
)

// DocsRouter serves about page.
type DocsRouter struct {
	baseRouter
}

// Get implemented Get method for DocsRouter.
func (this *DocsRouter) Get() {
	this.Data["IsDocs"] = true
	this.TplNames = "docs.html"

	reqUrl := this.Ctx.Request.URL.String()
	fullName := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
	if qm := strings.Index(fullName, "?"); qm > -1 {
		fullName = fullName[:qm]
	}

	this.Data["Sections"] = models.TplTree.Sections

	if fullName == "docs" {
		this.Redirect("/docs/Overview_Introduction", 302)
		return
	}

	df := models.GetDoc(fullName, this.Lang)
	if df == nil {
		this.Redirect("/docs/Overview_Introduction", 302)
		return
	}

	this.Data[fullName] = true

	// Set showed section.
	i := strings.Index(fullName, "_")
	if i > -1 {
		sec := fullName[:i]
		node := fullName[i+1:]
		this.Data["Section"] = getSection(sec)
		this.Data["Node"] = node
		this.Data["FullName"] = fullName
	}
	this.Data["Title"] = df.Title
	this.Data["Data"] = string(df.Data)
	this.Data["IsHasMarkdown"] = true
}

func getSection(name string) *models.Section {
	for _, sec := range models.TplTree.Sections {
		if sec.Name == name {
			return &sec
		}
	}
	return nil
}
