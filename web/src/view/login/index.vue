<template>
  <div class="container">
    <div class="login-window">
      <div class="logo" />
      <el-form ref="loginRef" :model="loginForm" status-icon>
        <el-form-item prop="username">
          <el-input v-model.trim="loginForm.username" placeholder="账号" :prefix-icon="User" style="width: 300px" />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model.trim="loginForm.password"
            type="password"
            placeholder="密码"
            show-password
            :prefix-icon="Lock"
            style="width: 300px"
          />
        </el-form-item>
        <el-form-item>
          <div class="vPicBox">
            <el-input v-model="loginForm.captcha" placeholder="请输入验证码" style="width: 60%" />
            <div class="vPic">
              <img v-if="picPath" :src="picPath" alt="请输入验证码" @click="loginVerify()">
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="danger" style="width: 300px">登陆</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login'
}
</script>

<script setup>
import { captcha } from '@/api/user'
import { reactive, ref } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'

const loginRef = ref('null')
const loginForm = reactive({ username: '', password: '', captcha: '', captchaId: '' })

// 表单校验
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})

// 获取验证码
const picPath = ref('')

const loginVerify = () => {
  captcha({}).then((ele) => {
    rules.captcha[1].max = ele.data.captchaLength
    rules.captcha[1].min = ele.data.captchaLength
    picPath.value = ele.data.picPath
    loginForm.captchaId = ele.data.captchaId
  })
}
loginVerify()

</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  height: 100%;
  background-image: url('@/assets/login1.jpg');
  background-size: cover;
  display: flex;
  justify-content: center;
  align-items: center;
  .login-window {
    background-color: rgb(255, 255, 255);
    padding: 40px 40px 30px 40px;
    border-radius: 10px;
    box-shadow: 2px 3px 7px rgb(0 0 0 / 20%);
    .logo {
      box-sizing: border-box;
      margin-bottom: 20px;
      background-image: url('@/assets/logo.png');
      background-size: cover;
      background-repeat: no-repeat;
      height: 20px;
      width: 120px;
      position: relative;
      left: 50%;
      transform: translateX(-50%);
    }
    .vPicBox {
      display: flex;
      justify-content: space-between;
      width: 100%;
      .vPic {
        width: 33%;
        height: 38px;
        background: #ccc;
        img {
          width: 100%;
          height: 100%;
          vertical-align: middle;
        }
      }
    }
  }
}
</style>
