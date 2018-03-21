// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/vmware/harbor/src/common"
)

// Label holds information used for a label
type Label struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	Name         string    `orm:"column(name)" json:"name"`
	Description  string    `orm:"column(description)" json:"description"`
	Color        string    `orm:"column(color)" json:"color"`
	Level        string    `orm:"column(level)" json:"-"`
	Scope        string    `orm:"column(scope)" json:"scope"`
	ProjectID    int64     `orm:"column(project_id)" json:"project_id"`
	CreationTime time.Time `orm:"column(creation_time)" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time)" json:"update_time"`
}

// TableName ...
func (l *Label) TableName() string {
	return "harbor_label"
}

// LabelQuery : query parameters for labels
type LabelQuery struct {
	Name      string
	Level     string
	Scope     string
	ProjectID int64
	Pagination
}

// Valid ...
func (l *Label) Valid(v *validation.Validation) {
	if len(l.Name) == 0 {
		v.SetError("name", "cannot be empty")
	}
	if len(l.Name) > 128 {
		v.SetError("name", "max length is 128")
	}

	if l.Scope != common.LabelScopeGlobal && l.Scope != common.LabelScopeProject {
		v.SetError("scope", fmt.Sprintf("invalid: %s", l.Scope))
	} else if l.Scope == common.LabelScopeProject && l.ProjectID <= 0 {
		v.SetError("project_id", fmt.Sprintf("invalid: %d", l.ProjectID))
	}
}

// ResourceLabel records the relationship between resource and label
type ResourceLabel struct {
	ID           int64     `orm:"pk;auto;column(id)"`
	LabelID      int64     `orm:"column(label_id)"`
	ResourceID   string    `orm:"column(resource_id)"`
	ResourceType string    `orm:"column(resource_type)"`
	CreationTime time.Time `orm:"column(creation_time)"`
	UpdateTime   time.Time `orm:"column(update_time)"`
}

// TableName ...
func (r *ResourceLabel) TableName() string {
	return "harbor_resource_label"
}