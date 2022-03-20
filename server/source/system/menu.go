package system

import (
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "odometer"}},
		// k8s Cluster
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "cluster", Name: "cluster", Component: "view/cluster/index.vue", Sort: 2, Meta: system.Meta{Title: "集群", Icon: "house"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "namespace", Name: "namespace", Component: "view/cluster/namespace/list.vue", Sort: 1, Meta: system.Meta{Title: "命名空间", Icon: "wind-power"}},
		{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "namespace/detail", Name: "namespace_detail", Component: "view/cluster/namespace/detail.vue", Sort: 1, Meta: system.Meta{Title: "命名空间详情", Icon: "wind-power"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "node", Name: "node", Component: "view/cluster/node/list.vue", Sort: 2, Meta: system.Meta{Title: "节点", Icon: "monitor"}},
		{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "node/detail", Name: "node_detail", Component: "view/cluster/node/detail.vue", Sort: 2, Meta: system.Meta{Title: "节点详情", Icon: "monitor"}},
		// k8s workloads
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "workloads", Name: "workloads", Component: "view/workloads/index.vue", Sort: 3, Meta: system.Meta{Title: "工作负载", Icon: "help-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "pod", Name: "pod", Component: "view/workloads/pod/list.vue", Sort: 3, Meta: system.Meta{Title: "Pods", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/detail", Name: "pod_detail", Component: "view/workloads/pod/detail.vue", Sort: 3, Meta: system.Meta{Title: "Pod详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/log", Name: "pod_log", Component: "view/workloads/pod/log.vue", Sort: 3, Meta: system.Meta{Title: "Pod日志", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/terminal", Name: "pod_terminal", Component: "view/workloads/pod/terminal.vue", Sort: 3, Meta: system.Meta{Title: "Pod终端", Icon: "ship"}},
		// superAdmin
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "12", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "pie-chart"}},
		// example
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "18", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: "媒体库", Icon: "upload"}},
		// person
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message"}},
	}
	if err := global.KOP_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.KOP_DB.Where("path = ?", "dashboard").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
