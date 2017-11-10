package admin
import(
	"testblog/models"
	"html/template"
)

type ArticleController struct{
	BaseController
}

func (c *ArticleController) Get(){
	blog, err := models.FindAllBlog()
	c.CheckErr(err, "get all blog error")

	c.Data["Blog"] = blog
        c.Data["Pagetitle"] = "文章列表"
        c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Layout = "layout/admin/admin.tpl"
	c.TplName = "admin/article/index.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["HtmlHead"] = "admin/article/index-css.tpl"
        c.LayoutSections["Scripts"] = "admin/article/index-js.tpl"
}

func (c *ArticleController) AddArticle(){
	section, err := models.FindAllSection()
	c.CheckErr(err, "query section from database error")
	label, err := models.FindAllLabel()
	c.CheckErr(err, "query label from database error")

	c.Data["Sections"] = section
	c.Data["Labels"] = label
        c.Data["Pagetitle"] = "新增文章"
        c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
        c.Layout = "layout/admin/admin.tpl"
        c.TplName = "admin/article/add.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["Scripts"] = "admin/article/add-js.tpl"
}

func (c *ArticleController) EditArticle(){

        c.Data["Pagetitle"] = "编辑文章"
        c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
        c.Layout = "layout/admin/admin.tpl"
        c.TplName = "admin/article/edit.tpl"
        //c.LayoutSections = make(map[string]string)
        //c.LayoutSections["Scripts"] = "admin/article/add-js.tpl"
}

func (c *ArticleController) DelArticle(){
}
