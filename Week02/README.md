### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

不应该，因为wrap抛出给上层后，上层必须处理这个error，当dao层换成其他工具包操作数据库(`如gorm返回ErrRecordNotFound`),返回的不是sql.ErrNoRows,
此时上层所有处理过sql.ErrNoRows的，必须全部重新处理dao层返回的新error

由于sql.ErrNoRows这个错误是可预知的(即需要判断有没有数据)  
我的想法是使用一个sentinel error，dao层遇到sql.ErrNoRows时，返回这个sentinel error。以后无论dao层如果切换数据库工具，上层都无需再处理这个error

对sql.ErrNoRows进行wrap操作个人感觉没有任何意义，毕竟不会去log这个error