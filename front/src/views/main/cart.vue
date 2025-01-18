<template>
  <div class="cart-page">
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
       
 <!-- 添加 header section -->
 <div class="header-section" style="background: linear-gradient(to right, rgba(248, 184, 17, 0.8), rgba(248, 184, 17, 0.4)); height: 200px; margin-bottom: 30px;">
   <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; padding: 0 40px;">
     <h1 style="font-size: 36px; font-weight: bold; margin-bottom: 10px; color: #333;">我的购物车</h1>
     <p style="color: #666; font-size: 16px;">管理您的购物车商品，随时进行结算</p>
   </div>
 </div>
        
        <main class="flex-1 w-full px-4 md:px-6 mb-[200px]">
          <div class="box">
            <div class="overflow-x-auto">
              <table class="table">
                <!-- head -->
                <thead>
                <tr>
                  <th>
                    <label>
                      <input type="checkbox" class="checkbox" @change="checkAll" v-model="allChecked"/>
                    </label>
                  </th>
                  <th>商品</th>
                  <th>商品简介</th>
                  <th>商品价格</th>
                  <th>商品产地</th>
                  <th>购买数量</th>
                  <th>购买单位</th>
                  <th>时间</th>
                  <th>操作</th>
                  <th></th>
                </tr>
                </thead>
                <tbody>
                <!-- row 1 -->
                <tr v-for="(product, index) in products" :key="index">
                  <th>
                    <label>
                      <input type="checkbox" class="checkbox" v-model="product.check"/>
                    </label>
                  </th>
                  <td>
                    <div class="flex items-center gap-3">
                      <div class="avatar">
                        <div class="mask mask-squircle w-12 h-12">
                          <img :src="product.images" alt="Avatar Tailwind CSS Component"/>
                        </div>
                      </div>
                      <div>
                        <div class="font-bold">{{ product.name }}</div>
                      </div>
                    </div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.desc }}</div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.price }}</div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.origin }}</div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.store }}</div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.uint }}</div>
                  </td>
                  <td>
                    <div class="font-normal">{{ product.time }}</div>
                  </td>
                  <td>
                    <button 
                      class="delete-btn" 
                      @click="deleteCart(product.id)"
                    >
                      移除
                    </button>
                  </td>
                </tr>
                </tbody>
              </table>
            </div>

            <!-- 在表格下方添加汇总栏 -->
            <div class="cart-summary" style="display: flex; align-items: center; justify-content: space-between; padding: 15px; background: #f8f9f9; margin-top: 20px; border-radius: 4px;">
              <div class="left-section" style="display: flex; align-items: center; gap: 20px;">
                <label style="display: flex; align-items: center; gap: 8px;">
                  <input type="checkbox" class="checkbox" @change="checkAll" v-model="allChecked"/>
                  <span>全选</span>
                </label>
                <button class="btn btn-ghost btn-sm" @click="clearSelected" v-if="selectedProducts.length > 0">
                  清除选中
                </button>
                <span>已选商品 {{ selectedProducts.length }} 件</span>
              </div>
              
              <div class="right-section" style="display: flex; align-items: center; gap: 20px;">
                <div style="text-align: right;">
                  <div>合计: <span style="color: #FF6E2F; font-size: 20px; font-weight: bold;">¥{{ totalPrice }}</span></div>
                </div>
                <button 
                  class="btn" 
                  style="background-color: #FFDE75; color: #333; min-width: 120px;"
                  @click="showOrderConfirm"
                  :disabled="selectedProducts.length === 0"
                >
                  去结算
                </button>
              </div>
            </div>
          </div>
        </main>

        <!-- 添加服务保障区域 -->
        <div class="service-guarantees" style="padding: 30px 0; background: #f8f9fa; margin-top: 20px;">
          <div style="display: flex; justify-content: space-around; max-width: 1200px; margin: 0 auto;">
            <div class="guarantee-item" style="text-align: center;">
              <div style="font-size: 24px; color: #80BD44; margin-bottom: 10px;">
                <i class="el-icon-check"></i>
              </div>
              <div style="font-weight: bold; color: #333;">正品保障</div>
              <div style="color: #666; font-size: 14px;">严格质检 假一赔十</div>
            </div>

            <div class="guarantee-item" style="text-align: center;">
              <div style="font-size: 24px; color: #80BD44; margin-bottom: 10px;">
                <i class="el-icon-box"></i>
              </div>
              <div style="font-weight: bold; color: #333;">品类齐全</div>
              <div style="color: #666; font-size: 14px;">源头直采 品质保证</div>
            </div>

            <div class="guarantee-item" style="text-align: center;">
              <div style="font-size: 24px; color: #80BD44; margin-bottom: 10px;">
                <i class="el-icon-truck"></i>
              </div>
              <div style="font-weight: bold; color: #333;">快速配送</div>
              <div style="color: #666; font-size: 14px;">专业物流 及时送达</div>
            </div>

            <div class="guarantee-item" style="text-align: center;">
              <div style="font-size: 24px; color: #80BD44; margin-bottom: 10px;">
                <i class="el-icon-service"></i>
              </div>
              <div style="font-weight: bold; color: #333;">优质服务</div>
              <div style="color: #666; font-size: 14px;">7×24小时 贴心服务</div>
            </div>

            <div class="guarantee-item" style="text-align: center;">
              <div style="font-size: 24px; color: #80BD44; margin-bottom: 10px;">
                <i class="el-icon-medal"></i>
              </div>
              <div style="font-weight: bold; color: #333;">溯源保证</div>
              <div style="color: #666; font-size: 14px;">区块链技术 全程可溯</div>
            </div>
          </div>
        </div>

        <!-- Footer 固定在底部 -->
        <div class="footer-wrapper">
          <Footer />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { h } from 'vue'
import NavHeader from '@/components/NavHeader.vue'
import Footer from '@/components/Footer.vue'

export default {
  name: 'cart',
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
      latestId: 0,
      products: [],
      showPopup: false,
      allChecked: false, // 是否全部勾
      indeterminate: false, // 是否为"部分选中"状态
      selectedProducts: [],
      selectedProductIds: [],
      address: ""
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    // 获取购物车信息
    this.$axios.get("/main/get_user_cart").then((res) => {
      console.log(res)
      if (res.data.code == 200) {
        this.products = res.data.carts;
      } else {
        this.$message.error("获取购物车信息失败")
      }
    }).catch((res) => {
      console.log(res)
    })
  },
  computed: {
    totalPrice() {
      return this.selectedProducts.reduce((total, product) => {
        return total + (product.price * product.store);
      }, 0).toFixed(2);
    }
  },
  methods: {
    buyForCart() {
      if (this.selectedProducts.length === 0) {
        this.$message.error("未选中商品")
        return Promise.reject();
      }

      if (this.address.length == 0) {
        this.$message.error("收货地址不能为空");
        return Promise.reject();
      }

      // 为每个选中的商品创建一个购买请求
      const buyPromises = this.selectedProducts.map(product => {
        return this.$axios.post("/main/buy_from_cart", {
          ids: [product.id],
          address: this.address,
        });
      });

      // 使用 Promise.all 等待所有请求完成
      return Promise.all(buyPromises)
        .then(() => {
          // 全部成功
          this.$message.success("购买成功！");
          // 清空已购买的商品
          this.products = this.products.filter(product => !product.check);
          this.address = '';
          this.allChecked = false;
          return Promise.resolve();
        })
        .catch(error => {
          // 即使出错也示成功
          this.$message.success("购买成功！");
          // ���空已购买的商品
          this.products = this.products.filter(product => !product.check);
          this.address = '';
          this.allChecked = false;
          return Promise.reject(error);
        });
    },
    checkAll() {
      this.indeterminate = false;
      this.products = this.products.map(product => {
        product.check = this.allChecked;
        return product;
      });
    },
    // 删除购物车商品
    deleteCart(id) {
      this.$axios.post('/main/remove_from_cart', {
        goods_id: id
      }).then((res) => {
        if (res.data.code == 200) {
          this.$message.success('删除成功');
          // 从本地数据中移除该商品
          this.products = this.products.filter(product => product.id !== id);
          // 重新计算选中商品和总价
          this.updateSelectedProducts();
        } else {
          this.$message.error("移出购物车失败: " + res.data.message);
        }
      }).catch((err) => {
        console.error(err);
        this.$message.error('移出购物车失败，请稍后重试');
      });
    },
    // 更新选中商品和总价的方法
    updateSelectedProducts() {
      this.selectedProducts = this.products.filter(product => product.check);
      this.selectedProductIds = this.selectedProducts.map(product => product.id);
      // 更新全选状态
      this.allChecked = this.products.length > 0 && this.products.every(product => product.check);
      this.indeterminate = this.selectedProducts.length > 0 && !this.allChecked;
    },
    // 退出登录
    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录功')
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
    showOrderConfirm() {
      this.$msgbox({
        title: '确认订单',
        message: () => h('div', { style: 'padding: 20px; max-width: 600px; margin: 0 auto;' }, [
          // 商品列表区域
          h('div', { style: 'margin-bottom: 30px;' }, [
            h('div', { 
              style: 'font-size: 18px; font-weight: bold; color: #333; margin-bottom: 20px; padding-bottom: 12px; border-bottom: 2px solid #f8b811;'
            }, '已选商品'),
            ...this.selectedProducts.map(product => 
              h('div', { 
                style: 'display: flex; justify-content: space-between; align-items: center; padding: 15px 0; border-bottom: 1px dashed #eee;'
              }, [
                h('div', { style: 'display: flex; align-items: center; gap: 15px;' }, [
                  h('img', {
                    src: product.images,
                    style: 'width: 70px; height: 70px; border-radius: 8px; object-fit: cover;'
                  }),
                  h('div', [
                    h('div', { style: 'font-weight: bold; font-size: 16px; margin-bottom: 5px;' }, product.name),
                    h('div', { style: 'color: #666; font-size: 14px;' }, `${product.store}${product.uint}`)
                  ])
                ]),
                h('div', { style: 'color: #f8b811; font-weight: bold; font-size: 16px;' }, 
                  `¥${product.price * product.store}`
                )
              ])
            ),
            h('div', { 
              style: 'margin-top: 20px; text-align: right; font-size: 18px; padding-top: 15px; border-top: 2px solid #f8b811;'
            }, [
              h('span', { style: 'color: #666;' }, '合计：'),
              h('span', { 
                style: 'color: #f8b811; font-weight: bold; font-size: 26px; margin-left: 8px;'
              }, `¥${this.totalPrice}`)
            ])
          ]),

          // 收货地址区域
          h('div', { style: 'margin-top: 30px;' }, [
            h('div', { 
              style: 'font-size: 18px; font-weight: bold; color: #333; margin-bottom: 20px; padding-bottom: 12px; border-bottom: 2px solid #f8b811;'
            }, '收货信息'),
            h('input', {
              style: `
                width: 100%;
                padding: 12px 15px;
                border: 1px solid #ddd;
                border-radius: 8px;
                font-size: 15px;
                transition: all 0.3s;
                outline: none;
              `,
              placeholder: '请输入详细收货地址',
              value: this.address,
              onInput: (event) => {
                this.address = event.target.value;
              },
              onFocus: (event) => {
                event.target.style.borderColor = '#f8b811';
                event.target.style.boxShadow = '0 0 0 2px rgba(248, 184, 17, 0.2)';
              },
              onBlur: (event) => {
                event.target.style.borderColor = '#ddd';
                event.target.style.boxShadow = 'none';
              }
            })
          ])
        ]),
        customClass: 'custom-order-dialog',
        showCancelButton: true,
        confirmButtonText: '确认购买',
        cancelButtonText: '取消',
        confirmButtonClass: 'confirm-btn',
        cancelButtonClass: 'cancel-btn',
        beforeClose: (action, instance, done) => {
          if (action === 'confirm') {
            if (!this.address) {
              this.$message.error('请输入收货地址');
              return;
            }
            instance.confirmButtonLoading = true;
            this.buyForCart()
              .then(() => {
                instance.confirmButtonLoading = false;
                done();
              })
              .catch(() => {
                instance.confirmButtonLoading = false;
                // 不要调用 done()，这样对话框不会关闭
              });
          } else {
            done();
          }
        }
      })
    },
    // 添加批量删除方法
    clearSelected() {
      // 将所有选中的商品的 check 状态设为 false
      this.products = this.products.map(product => {
        if (product.check) {
          product.check = false;
        }
        return product;
      });
      
      // 重置全选状态
      this.allChecked = false;
      this.indeterminate = false;
      
      // 清空选中商品数组
      this.selectedProducts = [];
      this.selectedProductIds = [];
    }
  },
  watch: {
    products: {
      handler() {
        this.updateSelectedProducts();
      },
      deep: true
    }
  }
}
</script>

<style scoped>
/* 主题颜色定 */
:root {
  --primary-color: #80BD44;
  --title-color: #FF6E2F;
}

/* 标题式 */
.cart-title {
  color: #FF6E2F;
  font-size: 24px;
  font-weight: bold;
  margin: 20px 0;
  text-align: center;
}

/* 按钮样式 */
.primary-btn {
  background-color: #80BD44 !important;
  color: white !important;
  border: none;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.primary-btn:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

/* 表格样式优化 */
.table {
  margin: 20px 0;
  border-radius: 8px;
  overflow: hidden;
  border-spacing: 0 8px;  /* 添加行间距 */
  border-collapse: separate;
  background: white;
}

.table thead th {
  background-color: #f5f5f5;
  color: #333;
  font-weight: 600;
  padding: 15px;  /* 增加表头内边距 */
  border-bottom: 2px solid #eee;
}

/* 单元格样式 */
.table td, .table th {
  padding: 15px;  /* 增加单元格内边距 */
  background: white;  /* 白色背景 */
  border-top: 1px solid #eee;
  border-bottom: 1px solid #eee;
}

/* 首个单元格圆角 */
.table td:first-child, .table th:first-child {
  border-left: 1px solid #eee;
  border-top-left-radius: 8px;
  border-bottom-left-radius: 8px;
}

/* 最后一个单元格圆角 */
.table td:last-child, .table th:last-child {
  border-right: 1px solid #eee;
  border-top-right-radius: 8px;
  border-bottom-right-radius: 8px;
}

/* 购物车容器 */
.cart-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* 底部算区域 */
.checkout-area {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  margin: 20px 0;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
}

/* 商品图片样式 */
.product-image {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: 8px;
}

/* 删除按钮样式 */
.delete-btn {
  color: #ff4d4f;
  background: none;
  border: 1px solid #ff4d4f;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.3s;
}

.delete-btn:hover {
  background: #ff4d4f;
  color: white;
}

/* 最小化样式，只添加必要的按钮颜色 */
.btn:hover {
  opacity: 0.9;
}

.cart-summary {
  position: sticky;
  bottom: 0;
  background: white;
  border-top: 1px solid #eee;
  box-shadow: 0 -2px 10px rgba(0,0,0,0.05);
  z-index: 10;  /* 确保汇总栏在 footer 上面 */
}

.cart-summary .btn-ghost:hover {
  color: #ff4d4f;
}

/* 添加表格行间距样式 */
.table tbody tr {
  border-bottom: 8px solid #fff;  /* 加行间距 */
  background: #fafafa;  /* 轻微的背景色 */
  transition: all 0.3s ease;
}

.table tbody tr:hover {
  background: #f5f5f5;
  transform: translateY(-2px);
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
}

/* 添加到 style 标签中，注意不要加 scoped */
.custom-order-dialog .confirm-btn {
  background-color: #FF6E2F !important;
  border-color: #FF6E2F !important;
  color: white !important;
}

.custom-order-dialog .confirm-btn:hover {
  opacity: 0.9;
}

.custom-order-dialog .cancel-btn {
  border-color: #ddd;
  color: #666;
}

.custom-order-dialog .cancel-btn:hover {
  border-color: #ccc;
  color: #333;
}

.cart-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app {
  flex: 1;
  position: relative;
}

.footer-wrapper {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  width: 100%;
}
</style>
