package routers

import (
	"cashAr/controllers"

	"github.com/astaxie/beego"
)

func init() {
	nsv1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/bulk_upload",
			beego.NSInclude(
				&controllers.BulkUploadController{},
			),
		),
	)
	beego.AddNamespace(nsv1)
}
