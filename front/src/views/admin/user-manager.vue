<template>
  <div class="modal" role="dialog" id="my_modal_8">
    <div class="modal-box">
      <h3 class="font-bold text-lg">修改用户信息</h3>
      <div class="grid gap-2" style="margin-top:20px">
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">
            用户名
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="username" disabled/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="password">
            密码
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="password" placeholder="输入你的密码" type="password" v-model="password"/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="security-question">
            问题
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="security-question" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保"/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="answer">
            答案
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="answer" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"/>
        </div>
      </div>
      <div class="modal-action">
        <a class="btn" @click="updateUserInfo">确定</a>
      </div>
    </div>
  </div>
  <div class="app">
    <div class="flex flex-col h-screen">
      <header class="flex items-center h-14 px-4 border-b lg:h-20 xl:px-6">
        <nav class="hidden gap-4 text-sm font-semibold lg:flex">
          <a class="text-gray-1200 dark:text-gray-400" style="font-size:20px" rel="ugc">
            水果溯源系统
          </a>
          <a class="text-gray-500 dark:text-gray-400" rel="ugc">
            <RouterLink to='/main'>水果审核</RouterLink>
          </a>
          <a class="text-gray-1000 dark:text-gray-400" href="#" rel="ugc">
            <RouterLink to='/user-manager'>用户管理</RouterLink>
          </a>
          <a class="text-gray-500 dark:text-gray-400" href="#" rel="ugc">
            <RouterLink to='/kind-manager'>水果分类管理</RouterLink>
          </a>
          <a class="text-gray-500 dark:text-gray-400" href="#" rel="ugc">
            <RouterLink to='/order-manager'>订单管理</RouterLink>
          </a>
        </nav>

        <div class="dropdown dropdown-end ml-auto">
          <div tabindex="0" role="button" class="btn btn-ghost rounded-btn dark:text-gray-400 ml-auto"
               v-text="username">
          </div>
          <ul tabindex="0" class="menu dropdown-content z-[1] p-2 shadow bg-base-100 rounded-box w-52 mt-4">
            <li @click="openModel"><a>修改用户信息</a></li>
            <li @click="logout"><a>退出登录</a></li>
          </ul>
        </div>
      </header>
      <main class="flex-1 grid w-full min-h-0 px-4 md:px-6">
        <div class="box">
          <a-table
            :columns="columns"
            :data-source="items"
            :pagination="false"
          >
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'status'">
                <span :class="[
                  'status-text',
                  record.disabled ? 'status-disabled' : 'status-enabled'
                ]">
                  {{ returnDisable(record.disabled) }}
                </span>
              </template>
              <template v-else-if="column.key === 'action'">
                <a-button 
                  type="primary" 
                  danger 
                  v-show="!record.disabled" 
                  @click="updateUserState(record.username, true)"
                >
                  停用账号
                </a-button>
                <a-button 
                  type="primary"
                  v-show="record.disabled" 
                  @click="updateUserState(record.username, false)"
                >
                  启用账号
                </a-button>
              </template>
            </template>
          </a-table>
          
          <div class="join" style="margin-left:45%;margin-right:45%;padding-bottom:20px">
            <button class="join-item btn" @click="handlePageChange(page - 1)">«</button>
            <button class="join-item btn">Page {{ page }}</button>
            <button class="join-item btn" @click="handlePageChange(page + 1)">»</button>
          </div>
        </div>
      </main>
    </div>
  </div>

</template>


<script>
import { Table, Button, message } from 'ant-design-vue'

export default {
  name: 'user-manager',
  components: {
    ATable: Table,
    AButton: Button,
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      items: [],
      page: 1,
      pageSize: 6,
      columns: [
        {
          title: '用户名',
          dataIndex: 'username',
          key: 'username',
          width: '200px'
        },
        {
          title: '身份',
          dataIndex: 'identity',
          key: 'identity',
          width: '150px'
        },
        {
          title: '状态',
          key: 'status',
          dataIndex: 'disabled',
          width: '100px'
        },
        {
          title: '操作',
          key: 'action',
          width: '120px',
          fixed: 'right'
        },
      ]
    }
  },

  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    
    const hide = message.loading('正在获取用户数据...', 0)
    this.$axios.get("/main/get_all_user_info", {
      params: {
        page: this.page
      }
    }).then((res) => {
      hide()
      if (res.data.code == 200) {
        this.items = res.data.users || []
      } else {
        message.error("获取用户列表失败: " + res.data.message)
      }
    }).catch((err) => {
      hide()
      console.error('获取用户失败:', err)
      message.error("获取用户失败，请稍后重试")
    })
  },

  methods: {
    handlePageChange(newPage) {
      if (newPage < this.page && this.page === 1) {
        message.error('已经是第一页了')
        return
      }
      if (newPage > this.page && this.items.length < 6) {
        message.error('已经是最后一页了')
        return
      }

      this.page = newPage
      const hide = message.loading('正在获取用户数据...', 0)
      this.$axios.get("/main/get_all_user_info", {
        params: {
          page: this.page
        }
      }).then((res) => {
        hide()
        if (res.data.code == 200) {
          this.items = res.data.users || []
        } else {
          message.error("获取用户列表失败: " + res.data.message)
        }
      }).catch((err) => {
        hide()
        console.error('获取用户失败:', err)
        message.error("获取用户失败，请稍后重试")
      })
    },

    returnDisable(bool) {
      return !bool ? "可用" : "不可用"
    },

    updateUserState(username, status) {
      const hide = message.loading('正在更新用户状态...', 0)
      this.$axios.post("/main/update_user_status", {
        username: username,
        status: status,
      }).then((res) => {
        hide()
        if (res.data.code == 200) {
          message.success("修改用户状态成功")
          const hideReload = message.loading('正在刷新用户数据...', 0)
          this.$axios.get("/main/get_all_user_info", {
            params: {
              page: this.page
            }
          }).then((res) => {
            hideReload()
            if (res.data.code == 200) {
              this.items = res.data.users || []
            }
          }).catch(() => {
            hideReload()
            message.error("刷新用户数据失败")
          })
        } else {
          message.error("修改用户状态失败: " + res.data.message)
        }
      }).catch((err) => {
        hide()
        console.error('修改用户状态失败:', err)
        message.error("修改用户状态失败，请稍后重试")
      })
    },

    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录成功')
      setTimeout(() => {
        this.$router.push('/')
      }, 1000)
    },

    openModel() {
      window.location.href = '#my_modal_8'
    },

    updateUserInfo() {
      this.$axios.post('/user/update_user_info', {
        username: this.username,
        password: this.password,
        question: this.question,
        answer: this.answer
      }).then((res) => {
        console.log(res)
      }).catch((res) => {
        console.log(res)
      })
      window.location.href = '#'
      this.password = ''
      this.question = ''
      this.answer = ''
    }
  }
}
</script>
<style scoped>
.app {
  background: #f5f5f5;
}

main {
  background: white;
  min-height: calc(100vh - 80px);
  padding: 20px;
}

:deep(.ant-table-wrapper) {
  margin: 20px 0;
}

:deep(.ant-btn) {
  margin-right: 8px;
}

/* 主按钮样式 */
:deep(.ant-btn-primary) {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

:deep(.ant-btn-primary:hover) {
  background-color: #e6a600 !important;
  border-color: #e6a600 !important;
}

/* 危险按钮样式 */
:deep(.ant-btn-primary.ant-btn-dangerous) {
  background-color: #FF6E2F !important;
  border-color: #FF6E2F !important;
}

:deep(.ant-btn-primary.ant-btn-dangerous:hover) {
  background-color: #f26529 !important;
  border-color: #f26529 !important;
}

/* 状态标签样式 */
.status-text {
  padding: 4px 8px;
  border-radius: 4px;
}

.status-disabled {
  background-color: #fff1f0;
  border: 1px solid #FF6E2F;
  color: #FF6E2F;
}

.status-enabled {
  background-color: #fff7f2;
  border: 1px solid #f8b811;
  color: #f8b811;
}

/* 分页器样式 */
:deep(.ant-pagination-item-active) {
  border-color: #f8b811 !important;
}

:deep(.ant-pagination-item-active a) {
  color: #f8b811 !important;
}

/* 导航链接样式 */
:deep(.router-link-active) {
  color: #f8b811 !important;
}

.text-gray-1200, 
.text-gray-1000,
.text-gray-500 {
  color: #2F3529 !important;
}

/* 下拉菜单样式 */
.dropdown-content {
  background: #ffffff;
}

.dropdown-content a {
  color: #2F3529;
}

.dropdown-content a:hover {
  color: #f8b811;
}

/* 分页按钮样式 */
.join {
  display: flex;
  justify-content: center;
  margin: 20px 0;
}

.join-item.btn {
  background-color: #ffffff;
  border: 1px solid #d9d9d9;
  color: #2F3529;
}

.join-item.btn:hover {
  border-color: #f8b811;
  color: #f8b811;
}

/* 布局样式 */
.box {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

header {
  background: #ffffff;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}
</style>