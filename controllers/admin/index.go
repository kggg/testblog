package admin

import (
	
)

type IndexController struct {
	BaseController
}

func (c *BaseController) Get(){
        c.Layout = "layout/admin/admin.tpl"
        c.TplName = "admin/index.tpl"
}

