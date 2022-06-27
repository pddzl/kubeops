import { createApp } from 'vue'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import router from '@/router/index'
import { store } from '@/pinia'
import App from './App.vue'
const app = createApp(App)
app.config.productionTip = false

app.use(store).use(router).use(ElementPlus, { locale: zhCn }).mount('#app')

export default app
