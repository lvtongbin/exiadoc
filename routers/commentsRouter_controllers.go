package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["trism/controllers:BaseExamController"] = append(beego.GlobalControllerRouter["trism/controllers:BaseExamController"],
		beego.ControllerComments{
			Method: "SimulateExam",
			Router: `/simulateExam`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:BaseStudentController"] = append(beego.GlobalControllerRouter["trism/controllers:BaseStudentController"],
		beego.ControllerComments{
			Method: "AddStudents",
			Router: `/addStudents`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:ClasssController"] = append(beego.GlobalControllerRouter["trism/controllers:ClasssController"],
		beego.ControllerComments{
			Method: "InitBaseClass",
			Router: `/initBaseClass`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:QueryController"] = append(beego.GlobalControllerRouter["trism/controllers:QueryController"],
		beego.ControllerComments{
			Method: "GetSummary",
			Router: `/getSummary`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:QueryController"] = append(beego.GlobalControllerRouter["trism/controllers:QueryController"],
		beego.ControllerComments{
			Method: "QueryExam",
			Router: `/queryExam`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:QueryController"] = append(beego.GlobalControllerRouter["trism/controllers:QueryController"],
		beego.ControllerComments{
			Method: "SelectClass",
			Router: `/selectClass`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["trism/controllers:QueryController"] = append(beego.GlobalControllerRouter["trism/controllers:QueryController"],
		beego.ControllerComments{
			Method: "SelectSchoolAndExam",
			Router: `/selectSchoolAndExam`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
