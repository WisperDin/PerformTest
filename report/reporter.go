package report

type ModuleResult struct {
	//接口响应时间
	RT int
}

type RoutineResult struct {
	//此用户线程是否完成
	IsFinish bool
	//此用户线程失败位置
	FailPos int
	//每个模块的测试结果
	ModuleResult []ModuleResult
}

type Reporter struct {
	//用户线程数
	UserNum int
	//每条用例的测试结果
	RoutineResult []RoutineResult
	//对于所有用例,平均的接口测试结果
	AverModuleResult []ModuleResult
	//成功的用户线程数
	SuccessUserNum int
	//失败的用户线程数
	FailUserNum int
}
