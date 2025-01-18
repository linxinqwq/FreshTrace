<template>
  <div class="app">
    <div class="relative flex items-center min-h-screen px-4 py-12 sm:px-6 lg:px-8">
      <div class="mx-auto w-full max-w-sm space-y-4 bg-white shadow-lg rounded-lg p-6">
        <form class="space-y-4">
          <p class="text-lg font-semibold text-center text-gray-700">注册</p>
          <div class="grid gap-4">
            <div>
              <label for="identity" class="text-sm font-medium">选择身份</label>
              <select class="select select-bordered w-full" v-model="identity">
                <option disabled selected>选择身份</option>
                <option>用户</option>
                <option>供应商</option>
                <option>管理员</option>
              </select>
            </div>
            <div>
              <label for="username" class="text-sm font-medium">用户名</label>
              <input
                class="input-field"
                id="username"
                placeholder="输入你的用户名"
                v-model="username"
              />
            </div>
            <div>
              <label for="password" class="text-sm font-medium">密码</label>
              <input
                class="input-field"
                id="password"
                placeholder="输入你的密码"
                type="password"
                v-model="password"
              />
            </div>
            <div>
              <label for="rePassword" class="text-sm font-medium">确认密码</label>
              <input
                class="input-field"
                id="rePassword"
                placeholder="再次输入你的密码"
                type="password"
                v-model="rePassword"
              />
            </div>
          </div>
          <button class="btn-primary" type="button" @click="register">注册</button>
        </form>
        <div class="text-center">
          <a class="font-semibold underline" @click="gotoLogin" style="cursor: pointer;">登录账号</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      identity: "选择身份",
      username: '',
      password: '',
      rePassword: ''
    }
  },
  methods: {
    gotoLogin() {
      this.$router.push("/login")
    },
    register() {
      if (this.username.length < 6) {
        this.$message.error("用户名不能短于六位");
      } else if (this.identity == '选择身份') {
        this.$message.error("请选择身份");
      } else if (this.password.length < 6 || this.rePassword.length < 6) {
        this.$message.error("密码不能短于六位");
      } else if (this.password != this.rePassword) {
        this.$message.error("密码两次输入不相同")
      } else {
        this.$axios.post('/user/register', {
          username: this.username,
          password: this.password,
          identity: this.identity
        }).then((res) => {
          if (res.data.code == 200) {
            this.$message.success(res.data.message)
            setTimeout(() => {
              this.$router.push('/')
            }, 1000); // 1秒后执行
          } else {
            this.$message.error(res.data.message)
          }
        }).catch((res) => {
          console.log(res);
        })

      }
    },
  }
}
</script>

<style scoped>
.app {
  background-image: url('../../assets/bg-2.jpg'), linear-gradient(270deg, rgba(255, 126, 95, 0.5), rgba(254, 180, 123, 0.5));
  background-size: cover, 400% 400%;
  animation: gradient 2s ease;
  height: 100vh;
  width: 100%;
}

@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 50% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.input-field {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  transition: border-color 0.3s;
}

.input-field:focus {
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(248, 184, 17, 0.1);
  outline: none;
}

.btn-primary {
  background-color: #f8b811;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 5px;
  width: 100%;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-primary:hover {
  opacity: 0.9;
}

.select-bordered:focus {
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(248, 184, 17, 0.1);
}

.underline:hover {
  color: #f8b811;
}
</style>
