package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/middleware"
	"github.com/pddzl/kubeops/server/middleware/log"
	"github.com/pddzl/kubeops/server/router"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.New()
	Router.Use(log.GinLogger(), log.GinRecovery(true))

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.KOP_CONFIG.Local.Path, http.Dir(global.KOP_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.KOP_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.KOP_LOG.Info("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.KOP_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	kubernetesRouter := router.RouterGroupApp.Kubernetes
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		// system
		systemRouter.InitApiRouter(PrivateGroup)                // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)               // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)               // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)             // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)             // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)          // 注册角色路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录

		// example
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

		// kubernetes
		kubernetesRouter.InitNamespaceRouter(PrivateGroup)  // namespace
		kubernetesRouter.InitNodeRouter(PrivateGroup)       // node
		kubernetesRouter.InitPodRouter(PrivateGroup)        // pod
		kubernetesRouter.InitReplicaSetRouter(PrivateGroup) // replicaSet
		kubernetesRouter.InitDeploymentRouter(PrivateGroup) // deployment
		kubernetesRouter.InitDaemonSetRouter(PrivateGroup)  //  daemonSet
		kubernetesRouter.InitServicesRouter(PrivateGroup)   // services
		kubernetesRouter.InitIngressRouter(PrivateGroup)    // Ingress
	}

	global.KOP_LOG.Info("router register success")
	return Router
}
