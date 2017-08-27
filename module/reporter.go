package module

type ModuleResult struct {
	//接口响应时间
	RT int
}

type RoutineResult struct {
	//此用户线程是否完成
	IsFinish bool
	//此用户线程失败位置
	FailPos int
}

type Reporter struct {
	//用户线程数
	UserNum int
	//每个模块的测试结果
	ModuleResult []ModuleResult
	//每条用例的测试结果
	RoutineResult []RoutineResult
	//成功的用户线程数
	SuccessUserNum int
	//失败的用户线程数
	FailUserNum int
}
