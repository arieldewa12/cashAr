package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["cashAr/controllers:BulkUploadController"] = append(beego.GlobalControllerRouter["cashAr/controllers:BulkUploadController"],
        beego.ControllerComments{
            Method: "BulkUploadCash",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
