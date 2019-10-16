// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package routers

import (
	"trism/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/query", &controllers.QueryController{}, "*:Index")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/class",
			beego.NSInclude(
				&controllers.ClasssController{},
			),
		),

		beego.NSNamespace("/exam",
			beego.NSInclude(
				&controllers.BaseExamController{},
			),
		),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.BaseStudentController{},
			),
		),

		beego.NSNamespace("/query",
			beego.NSInclude(
				&controllers.QueryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
