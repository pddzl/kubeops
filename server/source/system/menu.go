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
		// 1 dashboard 仪表盘
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "odometer"}},
		// 2 k8s Cluster 集群
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "cluster", Name: "cluster", Component: "view/cluster/index.vue", Sort: 2, Meta: system.Meta{Title: "集群", Icon: "house"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "namespace", Name: "namespace", Component: "view/cluster/namespace/list.vue", Sort: 1, Meta: system.Meta{Title: "命名空间", Icon: "wind-power"}},
		{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "namespace/detail", Name: "namespace_detail", Component: "view/cluster/namespace/detail.vue", Sort: 2, Meta: system.Meta{Title: "命名空间详情", Icon: "wind-power"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "node", Name: "node", Component: "view/cluster/node/list.vue", Sort: 3, Meta: system.Meta{Title: "节点", Icon: "monitor"}},
		{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "node/detail", Name: "node_detail", Component: "view/cluster/node/detail.vue", Sort: 4, Meta: system.Meta{Title: "节点详情", Icon: "monitor"}},
		// 7 k8s workloads 工作负载
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "workloads", Name: "workloads", Component: "view/workloads/index.vue", Sort: 3, Meta: system.Meta{Title: "工作负载", Icon: "help-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "pod", Name: "pod", Component: "view/workloads/pod/list.vue", Sort: 1, Meta: system.Meta{Title: "Pods", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/detail", Name: "pod_detail", Component: "view/workloads/pod/detail.vue", Sort: 2, Meta: system.Meta{Title: "Pod详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/log", Name: "pod_log", Component: "view/workloads/pod/log.vue", Sort: 3, Meta: system.Meta{Title: "Pod日志", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "pod/terminal", Name: "pod_terminal", Component: "view/workloads/pod/terminal.vue", Sort: 4, Meta: system.Meta{Title: "Pod终端", Icon: "ship"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "job", Name: "job", Component: "view/workloads/job/list.vue", Sort: 5, Meta: system.Meta{Title: "Job", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "job/detail", Name: "job_detail", Component: "view/workloads/job/detail.vue", Sort: 6, Meta: system.Meta{Title: "Job详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "replicaSet", Name: "replicaSet", Component: "view/workloads/replicaSet/list.vue", Sort: 7, Meta: system.Meta{Title: "ReplicaSet", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "replicaSet/detail", Name: "replicaSet_detail", Component: "view/workloads/replicaSet/detail.vue", Sort: 8, Meta: system.Meta{Title: "ReplicaSet详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "deployment", Name: "deployment", Component: "view/workloads/deployment/list.vue", Sort: 9, Meta: system.Meta{Title: "Deployment", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "deployment/detail", Name: "deployment_detail", Component: "view/workloads/deployment/detail.vue", Sort: 10, Meta: system.Meta{Title: "Deployment详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: false, ParentId: "7", Path: "daemonSet", Name: "daemonSet", Component: "view/workloads/daemonSet/list.vue", Sort: 11, Meta: system.Meta{Title: "DaemonSet", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "7", Path: "daemonSet/detail", Name: "daemonSet_detail", Component: "view/workloads/daemonSet/detail.vue", Sort: 12, Meta: system.Meta{Title: "DaemonSet详情", Icon: "ship"}},
		// 20 service 服务
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "service", Name: "service", Component: "view/service/index.vue", Sort: 4, Meta: system.Meta{Title: "服务", Icon: "van"}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "services", Name: "services", Component: "view/service/services/list.vue", Sort: 1, Meta: system.Meta{Title: "Services", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "20", Path: "services/detail", Name: "services_detail", Component: "view/service/services/detail.vue", Sort: 2, Meta: system.Meta{Title: "Services详情", Icon: "ship"}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "ingress", Name: "ingress", Component: "view/service/ingress/list.vue", Sort: 3, Meta: system.Meta{Title: "Ingress", Icon: "ship"}},
		{MenuLevel: 0, Hidden: true, ParentId: "20", Path: "ingress/detail", Name: "ingress_detail", Component: "view/service/ingress/detail.vue", Sort: 4, Meta: system.Meta{Title: "Ingress详情", Icon: "ship"}},
		// 25 superAdmin 超级管理员
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 5, Meta: system.Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "25", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 5, Meta: system.Meta{Title: "操作历史", Icon: "pie-chart"}},
		// 31 example 示例文件
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "31", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 1, Meta: system.Meta{Title: "媒体库", Icon: "upload"}},
		// person 个人信息
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 7, Meta: system.Meta{Title: "个人信息", Icon: "message"}},
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
