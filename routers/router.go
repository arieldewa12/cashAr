package routers

import (
	"cashAr/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// nsv1 := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/bulk_upload",
	// 		beego.NSInclude(
	// 			&controllers.BulkUploadController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/",
	// 		beego.NSInclude(
	// 			&controllers.MainController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(nsv1)
	beego.Router("/", &controllers.DashboardController{})
	beego.Router("/bulk_upload", &controllers.BulkUploadController{})
}
