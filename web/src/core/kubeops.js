// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
      欢迎使用 KubeOps
      当前版本:V2.0.0
    `)
  }
}
