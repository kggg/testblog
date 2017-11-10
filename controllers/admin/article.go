package admin
import(
	"testblog/models"
)

type ArticleController struct{
	BaseController
}

func (c *ArticleController) Get(){
	blog, err := models.FindAllBlog()
	c.CheckErr(err, "get all blog error")
	c.Data["Blog"] = blog
	c.Layout = "layout/admin/admin.tpl"
	c.TplName = "admin/article/index.tpl"
}
