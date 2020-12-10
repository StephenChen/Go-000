学习笔记  

### 作业题目：  

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

  dao 层遇到 sql.ErrNoRows 时，应该 wrap 这个 error 抛送给 service，service 在使用 WithMessage 进行包装。
