package biz

// UserDO  领域对象（ Domain Object， DO），微服务运行时核心业务对象的载体， DO 一般包括实体或值对象。
type UserDO struct {
	Id        int64
	Age       int32
	UserId    int64
	UserName  string
	Password  string
	Email     string
	Phone     string
	RoleName  string
	CreatedAt string
	UpdatedAt string
}
