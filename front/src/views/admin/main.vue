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
    v-model:visible="processModalVisible"
    title="商品步骤"
    @ok="updateProcess"
    @cancel="handleCancel"
  >
    <div class="process" style="margin-top: 20px;margin-bottom: 20px">
      <a-steps>
        <a-step title="商品上链" />
        <a-step 
          v-for="(step, index) in modalData.Process" 
          :key="index" 
          :title="step.Operation"
          @click="openPopup(index)"
        />
      </a-steps>
    </div>
  </a-modal>

  <a-modal
    v-model:visible="popupVisible"
    :title="PopupData.Operation"
    @ok="closePopup"
    @cancel="closePopup"
  >
    <p>时间: {{ PopupData.Date }}</p>
    <img class="popup-image" :src="PopupData.ImageId" alt="步骤图片">
  </a-modal>

  <div class="app">
    <div class="flex flex-col h-screen">
      <header class="flex items-center h-14 px-4 border-b lg:h-20 xl:px-6">
        <nav class="hidden gap-4 text-sm font-semibold lg:flex">
          <a class="text-gray-1200 dark:text-gray-400" style="font-size:20px" rel="ugc">
            水果溯源系统
          </a>
          <a class="text-gray-1000 dark:text-gray-400" rel="ugc">
            <RouterLink to='/main'>水果审核</RouterLink>
          </a>
          <a class="text-gray-500 dark:text-gray-400" href="#" rel="ugc">
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
            :pagination="{
              total: total,
              current: page,
              pageSize: 6,
              onChange: handlePageChange
            }"
          >
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'product'">
                <div class="flex items-center gap-3">
                  <div class="avatar">
                    <div class="mask mask-squircle w-12 h-12">
                      <img :src="record.ImageId" alt="商品图片"/>
                    </div>
                  </div>
                  <div>
                    <div class="font-bold">{{ record.Name }}</div>
                  </div>
                </div>
              </template>

              <template v-else-if="column.key === 'status'">
                <span :class="[
                  'status-text',
                  record.Status === 1 ? 'status-pending' : '',
                  record.Status === 3 ? 'status-rejected' : ''
                ]">
                  {{ getStatusText(record.Status) }}
                </span>
              </template>

              <template v-else-if="column.key === 'action'">
                <a-space>
                  <a-button type="link" @click="openProcess(record.Id)">查看步骤</a-button>
                  <a-button type="primary" @click="updateGoodsStatus(record.Id, 2)">通过</a-button>
                  <a-button danger @click="updateGoodsStatus(record.Id, 3)">拒绝</a-button>
                </a-space>
              </template>
            </template>
          </a-table>
        </div>
      </main>
    </div>
  </div>

  <div class="modal" role="dialog" id="lookProcess">
    <div class="modal-box">
      <h3 class="font-bold text-lg">商品步骤</h3>
      <div class="process" style="margin-top: 20px;margin-bottom: 20px">
        <ul class="steps">
          <li class="step step-primary">商品上链</li>
          <template v-for="(step, index) in modalData.Process" :key="index">
            <li class="step step-primary" @click="openPopup(index)">{{ step.Operation }}</li>
          </template>
        </ul>
      </div>
      <div class="modal-action">
        <a href="#" class="btn">关闭</a>
      </div>
    </div>
  </div>

  <div v-if="showPopup" class="popup-container">
    <h2 style="font-weight: bolder">操作: {{ PopupData.Operation }}</h2>
    <h2 style="font-weight: bolder">时间: {{ PopupData.Date }}</h2>
    <img class="mask mask-squircle" :src="PopupData.ImageId">
    <button class="btn" @click="closePopup">关闭</button>
  </div>

</template>


<script>
import { Table, Button, Modal, Steps, Space, message } from 'ant-design-vue'

export default {
  name: 'main',
  components: {
    [Table.name]: Table,
    [Button.name]: Button,
    [Modal.name]: Modal,
    [Steps.name]: Steps,
    [Steps.Step.name]: Steps.Step,
    [Space.name]: Space,
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      page: 1,
      pageSize: 6,
      total: 0,
      items: [],
      modalData: {},
      processModalVisible: false,
      popupVisible: false,
      PopupData: {},
      columns: [
        {
          title: '商品',
          key: 'product',
          dataIndex: 'product',
        },
        {
          title: '商品简介',
          dataIndex: 'Description',
          key: 'description',
        },
        {
          title: '商品库存',
          dataIndex: 'Stock',
          key: 'stock',
        },
        {
          title: '商品单价',
          dataIndex: 'Price',
          key: 'price',
        },
        {
          title: '商品单位',
          dataIndex: 'UintGoods',
          key: 'unit',
        },
        {
          title: '商品产地',
          dataIndex: 'Origin',
          key: 'origin',
        },
        {
          title: '商品销量',
          dataIndex: 'Sales',
          key: 'sales',
        },
        {
          title: '状态',
          key: 'status',
          dataIndex: 'Status'
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
      showPopup: false,
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.fetchGoods()
  },
  methods: {
    fetchGoods() {
      this.$axios.get("/main/get_unchecked_goods", {
        params: {
          page: this.page,
          pageSize: this.pageSize
        }
      }).then((res) => {
        if (res.data.code == 200) {
          this.items = res.data.products
          this.total = res.data.total
        } else {
          message.error("获取商品列表失败: " + res.data.message)
        }
      }).catch((err) => {
        console.log('获取商品失败:', err)
      })
    },

    handlePageChange(newPage) {
      this.page = newPage
      this.fetchGoods()
    },

    updateGoodsStatus(id, status) {
      this.$axios.post("/main/update_check_status", {
        id: id,
        status: status,
      }).then((res) => {
        if (res.data.code == 200) {
          message.success("修改商品状态成功: " + res.data.transactionHash)
          this.fetchGoods()
        } else {
          message.error("修改商品状态失败: " + res.data.message)
        }
      })
    },

    openProcess(id) {
      const item = this.items.find(item => item.Id === id)
      if (item) {
        this.modalData = {...item}
        window.location.href = '#lookProcess'
      }
    },

    openPopup(index) {
      this.PopupData = this.modalData.Process[index]
      this.showPopup = true
    },

    closePopup() {
      this.showPopup = false
    },

    getStatusText(status) {
      var kind = ["进展中", "待审核", "上架", "拒绝"]
      var index = parseInt(status, 10)
      if (index < 0 || index >= kind.length) {
        return "未知"
      }
      return kind[index]
    },

    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      message.success('退出登录成功')
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
    },
    updateProcess() {
      this.processModalVisible = false
      this.modalData = {}
    },
  }
}
</script>
<style>
.ant-table-wrapper {
  margin: 20px 0;
  background: #ffffff;
}

.ant-btn {
  margin-right: 8px;
}

.text-gray-1200, 
.text-gray-1000,
.text-gray-500,
.font-bold,
th,
td {
  color: #2F3529 !important;
}

.ant-btn-primary {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

.ant-btn-primary:hover {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

.ant-btn-dangerous {
  background-color: #FF6E2F !important;
  border-color: #FF6E2F !important;
  color: white !important;
}

.ant-btn-dangerous:hover {
  background-color: #f26529 !important;
  border-color: #f26529 !important;
}

.ant-btn-link {
  color: #FF6E2F !important;
}

.ant-btn-link:hover {
  color: #ff8952 !important;
}

.app {
  background: #f5f5f5;
}

.box {
  background: #ffffff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.ant-table-thead > tr > th {
  background: #f8f9f8 !important;
  color: #2F3529 !important;
}

.ant-table-tbody > tr > td {
  background: #ffffff;
}

.ant-tag-green {
  color: #f8b811 !important;
  background: #f6ffed !important;
  border-color: #f8b811 !important;
}

.ant-tag-orange {
  color: #FF6E2F !important;
  background: #fff7e6 !important;
  border-color: #FF6E2F !important;
}

.ant-steps .ant-steps-item-finish .ant-steps-item-icon {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

.ant-steps .ant-steps-item-finish .ant-steps-item-title {
  color: #2F3529 !important;
}

.ant-modal-title,
.ant-modal-content {
  color: #2F3529 !important;
}

.ant-pagination-item-active {
  border-color: #f8b811 !important;
}

.ant-pagination-item-active a {
  color: #f8b811 !important;
}

main {
  background: #f5f5f5;
}

header {
  background: #ffffff;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

footer {
  background: #ffffff;
  box-shadow: 0 -1px 3px rgba(0,0,0,0.1);
}

.popup-image {
  max-width: 100%;
  height: auto;
  margin-top: 16px;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  overflow: hidden;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 导航链接激活状态 */
.router-link-active {
  color: #f8b811 !important;
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

/* 状态文本颜色 */
.ant-table-tbody > tr > td .status-pending {
  color: #FF6E2F !important;
}

.ant-table-tbody > tr > td .status-rejected {
  color: #FF6E2F !important;
  font-weight: bold;
}

/* 分页器样式 */
.ant-pagination {
  margin-top: 20px !important;
  text-align: center !important;
}

.ant-pagination-prev .ant-pagination-item-link,
.ant-pagination-next .ant-pagination-item-link {
  color: #FF6E2F !important;
}

.ant-pagination-prev:hover .ant-pagination-item-link,
.ant-pagination-next:hover .ant-pagination-item-link {
  border-color: #FF6E2F !important;
  color: #FF6E2F !important;
}

.ant-pagination-item:hover {
  border-color: #FF6E2F !important;
}

.ant-pagination-item:hover a {
  color: #FF6E2F !important;
}

/* 表格鼠标悬停效果 */
.ant-table-tbody > tr:hover > td {
  background-color: #fff7f2 !important; /* 橙色的浅色背景 */
}

/* 修改表格中的状态显示 */
.status-text {
  padding: 4px 8px;
  border-radius: 4px;
}

.status-pending {
  background-color: #fff7e6;
  border: 1px solid #FF6E2F;
  color: #FF6E2F;
}

.status-rejected {
  background-color: #fff1f0;
  border: 1px solid #FF6E2F;
  color: #FF6E2F;
}

/* 模态框确认按钮 */
.ant-modal-footer .ant-btn-primary {
  background-color: #FF6E2F !important;
  border-color: #FF6E2F !important;
}

.ant-modal-footer .ant-btn-primary:hover {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

.ant-steps .ant-steps-item-process .ant-steps-item-icon {
  background-color: #f8b811 !important;
  border-color: #f8b811 !important;
}

.ant-steps .ant-steps-item-wait .ant-steps-item-icon {
  border-color: #f8b811 !important;
  color: #f8b811 !important;
}

.popup-container {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 9999;
  text-align: center;
}

.popup-container img {
  max-width: 100%;
  height: auto;
  margin: 20px 0;
}

.popup-container button {
  margin-top: 20px;
}

/* 步骤条样式 */
.steps {
  display: flex;
  overflow-x: auto;
  padding: 10px 0;
}

.step {
  cursor: pointer;
  padding: 10px;
  margin: 0 5px;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.step:hover {
  background-color: rgba(128, 189, 68, 0.1);
}

.step-primary {
  color: #f8b811;
}
</style>