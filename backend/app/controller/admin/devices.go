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
}

func (c *DevicesController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	hostname := c.Ctx.URLParamDefault("hostname", "")
	username := c.Ctx.URLParamDefault("username", "")
	rustdesk_id := c.Ctx.URLParamDefault("rustdesk_id", "")
	department := c.Ctx.URLParamDefault("department", "")

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
		list = append(list, iris.Map{
			"id":          a.Id,
			"alias":       a.Alias,
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

	_, err := c.Db.Where("id = ?", form.Id).Cols("alias", "department").Update(&model.Device{
		Alias:      form.Alias,
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
