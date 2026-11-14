package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type DevicesController struct {
	basicController
}

func (c *DevicesController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/devices/list", "HandleList")
	b.Handle("POST", "/devices/update", "HandleUpdate")
	b.Handle("POST", "/devices/delete", "HandleDelete")
	b.Handle("GET", "/companies/list", "HandleCompanyList")
	b.Handle("GET", "/companies/tree", "HandleCompanyTree")
	b.Handle("POST", "/companies/add", "HandleCompanyAdd")
	b.Handle("POST", "/companies/delete", "HandleCompanyDelete")
	b.Handle("POST", "/departments/add", "HandleDepartmentAdd")
	b.Handle("POST", "/departments/delete", "HandleDepartmentDelete")
}

func (c *DevicesController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	hostname := c.Ctx.URLParamDefault("hostname", "")
	username := c.Ctx.URLParamDefault("username", "")
	rustdesk_id := c.Ctx.URLParamDefault("rustdesk_id", "")
	department := c.Ctx.URLParamDefault("department", "")
	company := c.Ctx.URLParamDefault("company", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Device{})

		if hostname != "" {
			q.Where("hostname LIKE ? OR alias LIKE ?", "%"+hostname+"%", "%"+hostname+"%")
		}
		if username != "" {
			q.Where("username LIKE ?", "%"+username+"%")
		}
		if rustdesk_id != "" {
			q.Where("rustdesk_id LIKE ?", "%"+rustdesk_id+"%")
		}
		if department != "" && department != "Tümü" {
			q.Where("department = ?", department)
		}
		if company != "" && company != "Tümü" {
			q.Where("company = ?", company)
		}
		q.Desc("is_online").Desc("updated_at")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	deviceList := make([]model.Device, 0)

	err := pagination.Paginate(query, &model.Device{}, &deviceList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, a := range deviceList {
		dept := a.Department
		if dept == "" {
			dept = "Genel"
		}
		comp := a.Company
		if comp == "" {
			comp = "Centeks Group"
		}
		list = append(list, iris.Map{
			"id":          a.Id,
			"alias":       a.Alias,
			"company":     comp,
			"department":  dept,
			"rustdesk_id": a.RustdeskId,
			"hostname":    a.Hostname,
			"username":    a.Username,
			"uuid":        a.Uuid,
			"version":     a.Version,
			"os":          a.Os,
			"memory":      a.Memory,
			"is_online":   a.IsOnline,
			"created_at":  a.CreatedAt.Format(config.TimeFormat),
			"updated_at":  a.UpdatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

type DeviceUpdateForm struct {
	Id         int    `json:"id"`
	Alias      string `json:"alias"`
	Company    string `json:"company"`
	Department string `json:"department"`
}

func (c *DevicesController) HandleUpdate() mvc.Result {
	var form DeviceUpdateForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}
	if form.Id == 0 {
		return c.Error(nil, "Geçersiz cihaz ID")
	}

	if form.Company != "" && form.Company != "Tümü" {
		var existingComp model.Company
		has, _ := c.Db.Where("name = ?", form.Company).Get(&existingComp)
		if !has {
			c.Db.Insert(&model.Company{Name: form.Company})
		}
		if form.Department != "" && form.Department != "Tümü" {
			var existingDept model.Department
			hasD, _ := c.Db.Where("company_name = ? AND name = ?", form.Company, form.Department).Get(&existingDept)
			if !hasD {
				c.Db.Insert(&model.Department{
					CompanyName: form.Company,
					Name:        form.Department,
				})
			}
		}
	}

	_, err := c.Db.Where("id = ?", form.Id).Cols("alias", "company", "department").Update(&model.Device{
		Alias:      form.Alias,
		Company:    form.Company,
		Department: form.Department,
	})
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "Cihaz başarıyla güncellendi")
}

type DeviceDeleteForm struct {
	Id int `json:"id"`
}

func (c *DevicesController) HandleDelete() mvc.Result {
	var form DeviceDeleteForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}
	if form.Id == 0 {
		return c.Error(nil, "Geçersiz cihaz ID")
	}

	_, err := c.Db.Where("id = ?", form.Id).Delete(&model.Device{})
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "Cihaz silindi")
}

func (c *DevicesController) HandleCompanyList() mvc.Result {
	var companies []model.Company
	c.Db.Find(&companies)

	names := []string{"Centeks Group"}
	seen := map[string]bool{"Centeks Group": true}

	for _, comp := range companies {
		if !seen[comp.Name] && comp.Name != "" {
			names = append(names, comp.Name)
			seen[comp.Name] = true
		}
	}

	var devCompanies []string
	c.Db.Table("device").Distinct("company").Where("company != ''").Find(&devCompanies)
	for _, dc := range devCompanies {
		if !seen[dc] && dc != "" {
			names = append(names, dc)
			seen[dc] = true
		}
	}

	return c.Success(names, "ok")
}

type CompanyTreeItem struct {
	Name        string   `json:"name"`
	Note        string   `json:"note"`
	Departments []string `json:"departments"`
	DeviceCount int64    `json:"device_count"`
}

func (c *DevicesController) HandleCompanyTree() mvc.Result {
	var companies []model.Company
	c.Db.Find(&companies)

	hasCenteks := false
	for _, comp := range companies {
		if comp.Name == "Centeks Group" || comp.Name == "Centeks" {
			hasCenteks = true
			break
		}
	}
	if !hasCenteks {
		companies = append([]model.Company{{Name: "Centeks Group"}}, companies...)
	}

	var devCompanies []string
	c.Db.Table("device").Distinct("company").Where("company != ''").Find(&devCompanies)
	existingNames := map[string]bool{}
	for _, comp := range companies {
		existingNames[comp.Name] = true
	}
	for _, dc := range devCompanies {
		if !existingNames[dc] && dc != "" {
			companies = append(companies, model.Company{Name: dc})
			existingNames[dc] = true
		}
	}

	var departments []model.Department
	c.Db.Find(&departments)
	deptMap := map[string][]string{}
	for _, d := range departments {
		deptMap[d.CompanyName] = append(deptMap[d.CompanyName], d.Name)
	}

	if len(deptMap["Centeks Group"]) == 0 && len(deptMap["Centeks"]) == 0 {
		deptMap["Centeks Group"] = []string{"Bilgi İşlem", "Muhasebe", "Kalite", "Fabrika", "Sunucular"}
	}

	tree := make([]CompanyTreeItem, 0)
	for _, comp := range companies {
		count, _ := c.Db.Where("company = ?", comp.Name).Count(&model.Device{})
		depts := deptMap[comp.Name]
		if depts == nil {
			depts = []string{}
		}

		var devDepts []string
		c.Db.Table("device").Distinct("department").Where("company = ? AND department != ''", comp.Name).Find(&devDepts)
		seenDept := map[string]bool{}
		for _, d := range depts {
			seenDept[d] = true
		}
		for _, dd := range devDepts {
			if !seenDept[dd] && dd != "" {
				depts = append(depts, dd)
				seenDept[dd] = true
			}
		}

		tree = append(tree, CompanyTreeItem{
			Name:        comp.Name,
			Note:        comp.Note,
			Departments: depts,
			DeviceCount: count,
		})
	}

	return c.Success(tree, "ok")
}

type CompanyAddForm struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

func (c *DevicesController) HandleCompanyAdd() mvc.Result {
	var form CompanyAddForm
	if err := c.Ctx.ReadJSON(&form); err != nil || form.Name == "" {
		return c.Error(nil, "Firma adı boş olamaz")
	}

	var existing model.Company
	has, _ := c.Db.Where("name = ?", form.Name).Get(&existing)
	if !has {
		c.Db.Insert(&model.Company{
			Name: form.Name,
			Note: form.Note,
		})
	}
	return c.Success(nil, "Firma başarıyla kaydedildi")
}

type CompanyDeleteForm struct {
	Name string `json:"name"`
}

func (c *DevicesController) HandleCompanyDelete() mvc.Result {
	var form CompanyDeleteForm
	if err := c.Ctx.ReadJSON(&form); err != nil || form.Name == "" {
		return c.Error(nil, "Firma adı boş olamaz")
	}
	if form.Name == "Centeks Group" || form.Name == "Centeks" {
		return c.Error(nil, "Ana firma silinemez")
	}

	c.Db.Where("name = ?", form.Name).Delete(&model.Company{})
	c.Db.Where("company_name = ?", form.Name).Delete(&model.Department{})
	return c.Success(nil, "Firma ve bağlı departmanları silindi")
}

type DepartmentAddForm struct {
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
}

func (c *DevicesController) HandleDepartmentAdd() mvc.Result {
	var form DepartmentAddForm
	if err := c.Ctx.ReadJSON(&form); err != nil || form.Name == "" || form.CompanyName == "" {
		return c.Error(nil, "Firma ve departman adı boş olamaz")
	}

	var existing model.Department
	has, _ := c.Db.Where("company_name = ? AND name = ?", form.CompanyName, form.Name).Get(&existing)
	if !has {
		c.Db.Insert(&model.Department{
			CompanyName: form.CompanyName,
			Name:        form.Name,
		})
	}
	return c.Success(nil, "Departman başarıyla eklendi")
}

type DepartmentDeleteForm struct {
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
}

func (c *DevicesController) HandleDepartmentDelete() mvc.Result {
	var form DepartmentDeleteForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}
	c.Db.Where("company_name = ? AND name = ?", form.CompanyName, form.Name).Delete(&model.Department{})
	return c.Success(nil, "Departman silindi")
}
