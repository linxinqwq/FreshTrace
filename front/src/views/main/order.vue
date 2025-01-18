<template>
  <div class="order-page">
    <NavHeader 
      :username="username"
      :identity="identity"
      @open-model="openModel"
      @logout="logout"
    />

    <div class="modal" role="dialog" id="my_modal_8">
      <div class="modal-box">
        <h3 class="font-bold text-lg">修改用户信息</h3>
        <div class="grid gap-2" style="margin-top:20px">
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">用户名</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="username" disabled>
          </div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="password">密码</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="password" placeholder="输入你的密码" type="password" v-model="password"></div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">问题</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保"></div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">答案</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"></div>
        </div>
        <div class="modal-action">
          <a class="btn" @click="updateUserInfo">确定</a>
        </div>
      </div>
    </div>
    <div class="app">
      <div class="flex flex-col min-h-screen">
        <div class="header-section" style="background: linear-gradient(to right, rgba(248, 184, 17, 0.8), rgba(248, 184, 17, 0.4)); height: 200px; margin-bottom: 30px;">
      <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; padding: 0 40px;">
        <h1 style="font-size: 36px; font-weight: bold; margin-bottom: 10px; color: #333;">我的订单</h1>
        <p style="color: #666; font-size: 16px;">查看订单详情，追踪溯源信息</p>
      </div>
    </div>
        <main class="flex-1 grid w-full min-h-0 px-4 md:px-6">
          <div class="box">
            <div class="filter-section" style="margin: 20px 0; padding: 15px; background: #f8f9fa; border-radius: 8px;">
              <div style="display: flex; gap: 20px; align-items: center;">
                <div class="date-filter" style="display: flex; align-items: center; gap: 10px;">
                  <span>时间范围：</span>
                  <input 
                    type="date" 
                    v-model="filters.startDate"
                    class="input input-bordered input-sm"
                    style="width: 150px;"
                  />
                  <span>至</span>
                  <input 
                    type="date" 
                    v-model="filters.endDate"
                    class="input input-bordered input-sm"
                    style="width: 150px;"
                  />
                </div>

                <div class="status-filter" style="display: flex; align-items: center; gap: 10px;">
                  <span>审核状态：</span>
                  <select 
                    v-model="filters.status" 
                    class="select select-bordered select-sm"
                    style="width: 120px;"
                  >
                    <option value="">全部</option>
                    <option value="0">未审核</option>
                    <option value="1">已审核</option>
                  </select>
                </div>

                <button 
                  class="btn btn-sm"
                  style="background-color: #FFDE75; color: #333;"
                  @click="applyFilters"
                >
                  筛选
                </button>
                <button 
                  class="btn btn-sm"
                  style="background-color: #FF6E2F; color: white; border: none;"
                  @click="resetFilters"
                >
                  重置
                </button>
              </div>
            </div>

            <div class="overflow-x-auto">
              <table class="table">
                <!-- head -->
                <thead>
                <tr>
                  <th>商品</th>
                  <th>区块号</th>
                  <th>商品单价</th>
                  <th>购买数量</th>
                  <th>订购总价</th>
                  <th>购买时间</th>
                  <th>配送地址</th>
                  <th>状态</th>
                  <th>操作</th>
                  <th></th>
                </tr>
                </thead>
                <tbody>
                <!-- row 1 -->
                <tr v-for="(product, index) in order" :key="index">
                  <td>
                    <div class="flex items-center gap-3">
                      <div class="avatar">
                        <div class="mask mask-squircle w-12 h-12">
                          <img :src="product.image" alt="Avatar Tailwind CSS Component"/>
                        </div>
                      </div>
                      <div>
                        <div class="font-bold">{{ product.name }}</div>
                        <div class="text-sm opacity-50">{{ product.senders }}</div>
                      </div>
                    </div>
                  </td>
                  <td>
                    <span class="badge badge-ghost badge-sm">{{ product.blockNumber }}</span>
                  </td>
                  <td>{{ product.price }}</td>
                  <td>{{ product.number }}</td>
                  <td>{{ product.totalPrice }}</td>
                  <td>{{ product.createTime }}</td>
                  <td>{{ product.address }}</td>
                  <td>{{ showStatus(product.status) }}</td>
                  <th>
                    <button class="btn btn-ghost btn-xs" @click="navigateToProductDetail(product.id)">商品详情</button>
                  </th>
                </tr>
                </tbody>
              </table>
            </div>
            <div class="flex justify-center mt-8 mb-8">
              <div class="join">
                <button class="join-item btn" @click="prevPage"
                  style="background-color: #FF6E2F; color: white; border: none;">«</button>
                <button class="join-item btn"
                  style="background-color: #FF6E2F; color: white; border: none;">Page {{ page }}</button>
                <button class="join-item btn" @click="nextPage"
                  style="background-color: #FF6E2F; color: white; border: none;">»</button>
              </div>
            </div>
          </div>

        </main>
      </div>
    </div>
    <Footer />
  </div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
import Footer from '@/components/Footer.vue'

export default {
  name: 'order',
  components: {
    NavHeader,
    Footer
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      page: 1,
      order: [],
      filters: {
        startDate: '',
        endDate: '',
        status: ''
      },
      originalOrders: [], // ��储原始订单数据
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.$axios.get("/main/get_buy_record", {
      params: {
        page: this.page
      }
    }).then((res) => {
      if (res.data.code == 200) {
        this.originalOrders = res.data.products; // 保存原始数据
        this.order = res.data.products; // 显示数据
      } else {
        this.$message.error("获取订单失败: " + res.data.message)
      }
      console.log(res);
    }).catch((res) => {
      console.log(res);
    })
  },
  methods: {
    showStatus(data) {
      if (data == 0) {
        return "未审核"
      }
      return "已审核"
    },
    // 退出登录
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
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.$axios.get("/main/get_buy_record", {
          params: {
            page: this.page,
          }
        }).then((res) => {
          if (res.data.code == 200) {
            this.order = res.data.products
          } else {
            this.$message.error("获取商品列表失败: " + res.data.message)
          }
          console.log(res);
        }).catch((res) => {
          console.log(res);
        })
      } else {
        this.$message.error('已经是第一页了')
        return
      }
    },
    nextPage() {
      // 这里您可能需要一个逻辑来检查是否已经是最后一页
      if (this.order.length == 6) {
        this.page += 1;
      } else {
        this.$message.error('已经是最后一页了')
        return
      }
      this.$axios.get("/main/get_buy_record", {
        params: {
          page: this.page,
        }
      }).then((res) => {
        if (res.data.code == 200) {
          this.order = res.data.products
        } else {
          this.$message.error("获取商品列表失败: " + res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
    },
    navigateToProductDetail(productId) {
      this.$router.push({path: `/buy-detail`, query: {id: productId}});
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
    // 应用筛选
    applyFilters() {
      let filteredOrders = [...this.originalOrders];

      // 时间筛选
      if (this.filters.startDate && this.filters.endDate) {
        const start = new Date(this.filters.startDate);
        const end = new Date(this.filters.endDate);
        // 将结束时间设置为当天的最后一刻
        end.setHours(23, 59, 59, 999);
        
        filteredOrders = filteredOrders.filter(order => {
          // 检查购买时间
          const orderDate = new Date(order.createTime);
          // 使用 >= 和 <= 来包含起始和束日期
          const isOrderDateInRange = orderDate >= start && orderDate <= end;

          return isOrderDateInRange;
        });
      }

      // 状态筛选
      if (this.filters.status !== '') {
        filteredOrders = filteredOrders.filter(order => 
          order.status.toString() === this.filters.status
        );
      }

      this.order = filteredOrders;
      this.page = 1; // 重置页码
    },

    // 重置筛选
    resetFilters() {
      this.filters = {
        startDate: '',
        endDate: '',
        status: ''
      };
      this.order = [...this.originalOrders];
      this.page = 1; // 重置页码
    },
  }
}
</script>

<style scoped>
.filter-section {
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.filter-section input[type="date"],
.filter-section select {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 4px 8px;
}

.filter-section input[type="date"]:focus,
.filter-section select:focus {
  border-color: #FFDE75;
  box-shadow: 0 0 0 2px rgba(255, 222, 117, 0.2);
}

.join-item.btn {
  min-width: 60px;
}

.join-item.btn:hover {
  opacity: 0.9;
}
</style>
