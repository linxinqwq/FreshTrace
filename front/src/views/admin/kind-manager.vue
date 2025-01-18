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
  <a-modal
    v-model:visible="kindModalVisible"
    title="添加水果种类"
    @ok="addFruit"
    @cancel="handleCancel"
  >
    <a-form :model="formState">
      <a-form-item label="水果种类名">
        <a-input v-model:value="kind" placeholder="请输入商品种类(不能重复)" />
      </a-form-item>
    </a-form>
  </a-modal>
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
          <a class="text-gray-500 dark:text-gray-400" href="#" rel="ugc">
            <RouterLink to='/user-manager'>用户管理</RouterLink>
          </a>
          <a class="text-gray-1000 dark:text-gray-400" href="#" rel="ugc">
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
      <main class="flex-1 p-4">
        <a-button 
          type="primary"
          style="margin-bottom: 16px;"
          @click="showKindModal"
        >
          添加种类
        </a-button>
        <a-table
          :columns="columns"
          :data-source="items"
          :pagination="{
            total: total,
            current: page,
            pageSize: pageSize,
            onChange: handlePageChange,
            showTotal: total => `共 ${total} 条`
          }"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
              <a-button type="primary" danger @click="removeFruit(record.ID)">
                删除种类
              </a-button>
            </template>
          </template>
        </a-table>
      </main>
    </div>
  </div>

</template>


<script>
import { Table, Button, Modal, Form, Input, message } from 'ant-design-vue'

export default {
  name: 'kind-manager',
  components: {
    [Table.name]: Table,
    [Button.name]: Button,
    [Modal.name]: Modal,
    [Form.name]: Form,
    [Form.Item.name]: Form.Item,
    [Input.name]: Input,
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      allItems: [],
      items: [],
      page: 1,
      pageSize: 6,
      total: 0,
      kind: "",
      kindModalVisible: false,
      columns: [
        {
          title: '种类ID',
          dataIndex: 'ID',
          key: 'id',
        },
        {
          title: '种类名',
          dataIndex: 'kind',
          key: 'kind',
        },
        {
          title: '创建时间',
          dataIndex: 'CreatedAt',
          key: 'createdAt',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.fetchAllKinds()
  },
  methods: {
    fetchAllKinds() {
      this.$axios.get("/main/get_all_fruit_kind").then((res) => {
        if (res.data.code == 200) {
          console.log('获取到的所有种类数据:', res.data)
          this.allItems = res.data.fruits
          this.total = this.allItems.length
          this.updatePageData()
        } else {
          message.error("获取种类列表失败: " + res.data.message)
        }
      }).catch((err) => {
        console.log('获取种类失败:', err)
      })
    },

    updatePageData() {
      const start = (this.page - 1) * this.pageSize
      const end = start + this.pageSize
      this.items = this.allItems.slice(start, end)
    },

    handlePageChange(newPage) {
      console.log('页码变化:', newPage)
      this.page = newPage
      this.updatePageData()
    },

    showKindModal() {
      this.kindModalVisible = true
    },

    handleCancel() {
      this.kindModalVisible = false
      this.kind = ''
    },

    addFruit() {
      this.$axios.post("/main/add_fruit_kind", {
        kind: this.kind
      }).then((res) => {
        if (res.data.code == 200) {
          message.success("添加水果种类成功")
          this.kindModalVisible = false
          this.kind = ''
          this.fetchAllKinds()
        } else {
          message.error("添加水果种类失败: " + res.data.message)
        }
      }).catch((err) => {
        console.log(err)
      })
    },

    removeFruit(id) {
      Modal.confirm({
        title: '确认删除',
        content: '确定要删除这个种类吗？',
        onOk: () => {
          this.$axios.post("/main/remove_fruit_kind", {
            ID: id
          }).then((res) => {
            if (res.data.code == 200) {
              message.success("删除水果种类成功")
              this.fetchAllKinds()
            } else {
              message.error("删除水果种类失败: " + res.data.message)
            }
          }).catch((err) => {
            console.log(err)
          })
        }
      })
    },

    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录成功')
      setTimeout(() => {
        this.$router.push('/')
      }, 1000); // 1秒后执行
      // 跳转到首页
    },
    openModel() {
      window.location.href = '#my_modal_8';
    },
    updateUserInfo() {
      this.$axios.post('/user/update_user_info', {
        username: this.username,
        password: this.password,
        question: this.question,
        answer: this.answer
      }).then((res) => {
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
      window.location.href = '#';
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
:deep(.ant-btn-dangerous) {
  background-color: #FF6E2F !important;
  border-color: #FF6E2F !important;
  color: white !important;
}

:deep(.ant-btn-dangerous:hover) {
  background-color: #f26529 !important;
  border-color: #f26529 !important;
}

/* 表格样式 */
:deep(.ant-table-tbody > tr:hover > td) {
  background-color: #fff7f2 !important;
}

/* 分页器样式 */
:deep(.ant-pagination-item-active) {
  border-color: #f8b811 !important;
}

:deep(.ant-pagination-item-active a) {
  color: #f8b811 !important;
}

/* 模态框样式 */
:deep(.ant-modal-footer .ant-btn-primary) {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

:deep(.ant-modal-footer .ant-btn-primary:hover) {
  background-color: #e6a600 !important;
  border-color: #e6a600 !important;
}

/* 添加按钮样式 */
.add-button {
  width: 200px;
  height: 50px;
  border-radius: 25px;
  font-size: 16px;
  font-weight: 500;
  letter-spacing: 2px;
  text-transform: uppercase;
  background-color: #f8b811;
  border-color: #f8b811;
}

.add-button:hover {
  background-color: #e6a600;
  border-color: #e6a600;
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