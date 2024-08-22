package common

type AdminUserException string

const (
	MailExistsError     = "用户邮箱已经存在！"
	UserMailExistsError = "用户名已经存在！"
	MailOrNameNullError = "用户名或密码不能为空!"
	AdminUserNotExists  = "用户不存在!"
	PassWordNotMatch    = "两次密码输入不正确!"

	MenuNameNullError    = "菜单名不能为空！"
	MenuNameExistsError  = "菜单名称已经存在！"
	UpMenuNotExistsError = "上级菜单不存在！"
	SecondError          = "菜单目前只开放两级，请选择一级菜单！"
	MenuNotExists        = "菜单不存在"
	DeleteChildMenuError = "请先删除菜单下的子菜单！"

	ResourceNameOrPathNullError = "资源名和路径不能为空！"
	ResourceTypeNotExistsError  = "资源分类不存在！"
	ResourcePathExistsError     = "此资源路径已经存在系统！"
	ResourceTypeNameNullError   = "资源分类名不能为空！"
	ResourceTypeExistsError     = "资源分类已经存在！"
	ResourceNotExistsError      = "资源不存在！"
	DeleteResourceTypeError     = "资源分类下存在资源，请先删除资源再删除分类！"

	RoleNameNullError   = "角色名称不能为空！"
	RoleNameExistsError = "角色名称已经存在！"
	RoleNotExists       = "角色不存在！"
	RoleExists          = "角色已经存在！"

	LoginOutOfDate    = "登陆过期，请重新登陆！"
	UserAuthNotEnough = "当前用户角色权限不够！"
	ResourceNotAuth   = "当前资源未被授权，请联系管理员授权之后再使用！"
)
