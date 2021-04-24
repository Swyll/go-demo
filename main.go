package main

import (
	"go-demo/modules/error/dao"
	"go-demo/modules/error/server"
)

/**
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

如毛大所说，这种sentinel在基础库中最好不用包装之后抛给上层处理，
第一，因为数据层是该模块比较底层的一部分，如果抛到服务层去处理，那么代码重复性工作大
第二，对于没有查询出来的数据，对于服务来说不是什么关键性的错误，可以在dao层进行降级处理，如返回一个空的数据，等等
第三，我觉得错误Wrap是给我们更好的在运行环境上快速的定位到问题所在，那么对于这种不是代码设计的err不应该打印在日志中，以免干扰正常的日志
**/
func main() {
	s := server.NewServer(dao.NewEngine())

	info, err := s.GetStudentInfo("demo")
	if err != nil {
		return
	}

	server.GetLogger().Errorf("%+v\n", info)
}
