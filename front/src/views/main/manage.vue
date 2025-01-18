<template>
  <div class="app">
    <div class="flex flex-col h-screen">
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
      <div class="container">
        <div class="header">
          <h1>我的商品</h1>
          <button class="button" @click="openBuy" style="float: right; margin-bottom: 20px;">添加商品</button>
        </div>
        <table class="table">
          <thead>
            <tr>
              <th class="w-10">商品</th>
              <th class="w-15">商品简介</th>
              <th class="w-8">商品库存</th>
              <th class="w-8">商品单价(rmb)</th>
              <th class="w-8">商品单位</th>
              <th class="w-10">商品产地</th>
              <th class="w-8">商品销量</th>
              <th class="w-8">商品种类</th>
              <th class="w-8">状态</th>
              <th class="w-17">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in items" :key="index">
              <td>{{ item.Name }}</td>
              <td>{{ item.Description }}</td>
              <td>{{ item.Stock }}</td>
              <td>{{ item.Price }}</td>
              <td>{{ item.UintGoods }}</td>
              <td>{{ item.Origin }}</td>
              <td>{{ item.Sales }}</td>
              <td>{{ item.Kind }}</td>
              <td>{{ getStatusText(item.Status) }}</td>
              <td>
                <button class="button mb-2" @click="openUpdate(item.Id)" style="margin-right: 10px;">修改商品</button>
                <button v-if="getStatusText(item.Status) === '上架'" class="button mb-2" @click="removeGoods(item.Id, 4)"
                  style="margin-right: 10px;">下架商品</button>
                <button v-if="getStatusText(item.Status) === '下架'" class="button mb-2" @click="removeGoods(item.Id, 2)"
                  style="margin-right: 10px;">上架商品</button>
                <button class="button mb-2" @click="openProcess(item.Id)" style="margin-right: 10px;">查看步骤</button>
                <button class="button mb-2" @click="addGoodsProcess(item.Id)">添加步骤</button>
              </td>
            </tr>
          </tbody>
        </table>
        <div class="pagination">
          <button @click="prevPage">«</button>
          <button>Page {{ page }}</button>
          <button @click="nextPage">»</button>
        </div>
      </div>
    </div>
  </div>

  <!-- 修改商品模态窗口 -->
  <div class="modal" role="dialog" id="update">
    <div class="modal-box">
      <div class="modal-header">
        <h3 class="font-bold text-lg">修改商品</h3>
        <a href="#" class="close-btn">&times;</a>
      </div>
      <form @submit.prevent="validateAndUpdate" class="grid gap-2">
        <div class="form-group">
          <label for="name">商品名</label>
          <input v-model="modalData.Name" placeholder="请输入商品名" :class="{ 'error': errors.name }"
            @input="validateField('name')" required />
          <span class="error-message color-red" v-if="errors.name">{{ errors.name }}</span>
        </div>

        <div class="form-group">
          <label for="description">商品简介</label>
          <input v-model="modalData.Description" placeholder="请输入商品简介" :class="{ 'error': errors.description }"
            @input="validateField('description')" required />
          <span class="error-message color-red" v-if="errors.description">{{ errors.description }}</span>
        </div>

        <div class="form-group">
          <label for="stock">库存</label>
          <input v-model="modalData.Stock" type="number" placeholder="请输入商品库存" :class="{ 'error': errors.stock }"
            @input="validateField('stock')" required min="0" />
          <span class="error-message color-red" v-if="errors.stock">{{ errors.stock }}</span>
        </div>

        <div class="form-group">
          <label for="price">单价(rmb)</label>
          <input v-model="modalData.Price" type="number" placeholder="请输入商品单价" :class="{ 'error': errors.price }"
            @input="validateField('price')" required min="0" step="0.01" />
          <span class="error-message color-red" v-if="errors.price">{{ errors.price }}</span>
        </div>

        <div class="form-group">
          <label for="unit">商品单位</label>
          <input v-model="modalData.UintGoods" placeholder="请输入商品单位" :class="{ 'error': errors.unit }"
            @input="validateField('unit')" required />
          <span class="error-message color-red" v-if="errors.unit">{{ errors.unit }}</span>
        </div>

        <div class="form-group">
          <label for="origin">商品产地</label>
          <input v-model="modalData.Origin" placeholder="请输入商品产地" :class="{ 'error': errors.origin }"
            @input="validateField('origin')" required />
          <span class="error-message color-red" v-if="errors.origin">{{ errors.origin }}</span>
        </div>

        <div class="modal-action">
          <button type="submit" class="btn btn-primary">确定</button>
          <a href="#" class="btn">取消</a>
        </div>
      </form>
    </div>
  </div>

  <!-- 添加商品模态窗口 -->
  <div class="modal" role="dialog" id="add">
    <div class="modal-box">
      <div class="modal-header">
        <h3 class="font-bold text-lg">添加商品</h3>
        <a href="#" class="close-btn">&times;</a>
      </div>
      <div class="grid gap-2">
        <label for="name">商品名</label>
        <input v-model="name" placeholder="请输入商品名" />
        <label for="desc">商品简介</label>
        <input v-model="desc" placeholder="请输入商品简介" />
        <label for="store">库存</label>
        <input v-model="store" type="number" placeholder="请输入商品库存" />
        <label for="price">单价(rmb)</label>
        <input v-model="price" type="number" placeholder="请输入商品单价" />
        <label for="unit">商品单位</label>
        <input v-model="uints" placeholder="请输入商品单位" />
        <label for="origin">商品产地</label>
        <input v-model="origin" placeholder="请输入商品产地" />
        <label for="kind">商品分类</label>
        <select v-model="sKind">
          <option disabled value="请选择商品分类">请选择商品分类</option>
          <option v-for="category in kind" :value="category.kind">{{ category.kind }}</option>
        </select>
        <input type="file" class="file-input" @change="handleFileUpload" />
      </div>
      <div class="modal-action">
        <a href="#" class="btn" @click="addGoods">确定</a>
        <a href="#" class="btn">取消</a>
      </div>
    </div>
  </div>

  <!-- 查看步骤模态窗口 -->
  <div class="modal" role="dialog" id="lookProcess">
    <div class="modal-box">
      <div class="modal-header">
        <h3 class="font-bold text-lg">商品步骤</h3>
        <a href="#" class="close-btn">&times;</a>
      </div>
      <div class="steps-container">
        <div class="step-item">
          <div class="step-circle completed">1</div>
          <div class="step-content">
            <h4>商品上链</h4>
            <p class="step-date">初始步骤</p>
          </div>
        </div>

        <template v-for="(step, index) in modalData.Process" :key="index">
          <div class="step-connector"></div>
          <div class="step-item">
            <div class="step-circle completed">{{ index + 2 }}</div>
            <div class="step-content" @click="openPopup(index)">
              <h4>{{ step.Operation }}</h4>
              <p class="step-date">{{ step.Date }}</p>
            </div>
          </div>
        </template>
      </div>

      <div class="modal-action">
        <a href="#" class="btn">关闭</a>
      </div>
    </div>
  </div>

  <!-- 在其他模态窗口后添加 -->
  <div class="modal" role="dialog" id="addProcess">
    <div class="modal-box">
      <div class="modal-header">
        <h3 class="font-bold text-lg">添加商品步骤</h3>
        <a href="#" class="close-btn">&times;</a>
      </div>
      
      <!-- 步骤展示 -->
      <div class="steps-container">
        <div class="step-item">
          <div class="step-circle completed">1</div>
          <div class="step-content">
            <h4>商品上链</h4>
            <p class="step-date">初始步骤</p>
          </div>
        </div>

        <template v-for="(step, index) in modalData.Process" :key="index">
          <div class="step-connector"></div>
          <div class="step-item">
            <div class="step-circle completed">{{ index + 2 }}</div>
            <div class="step-content">
              <h4>{{ step.Operation }}</h4>
              <p class="step-date">{{ step.Date }}</p>
            </div>
          </div>
        </template>
      </div>

      <!-- 添加新步骤表单 -->
      <div class="form-section">
        <div class="operater">
          <label class="form-label">操作</label>
          <select class="select-primary" v-model="updateProcessData.operater">
            <option disabled value="选择步骤">选择步骤</option>
            <option value="播种">播种</option>
            <option value="浇水">浇水</option>
            <option value="施肥">施肥</option>
            <option value="采摘">采摘</option>
            <option value="病虫害防治">病虫害防治</option>
            <option value="质量检验">质量检验</option>
          </select>
        </div>

        <div class="status-toggle">
          <label class="form-label">是否上架</label>
          <input type="checkbox" class="toggle" v-model="updateProcessData.ifSelect"/>
        </div>

        <div class="image-upload">
          <label class="form-label">图片</label>
          <input type="file" 
            class="file-input file-input-bordered file-input-secondary" 
            @change="handleFileUpload"
          />
        </div>
      </div>

      <div class="modal-action">
        <a href="#" class="btn">取消</a>
        <a href="#" class="btn" @click="updateProcess">确定</a>
      </div>
    </div>
  </div>

  <Footer />
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
      fileId: '',
      name: '',
      desc: '',
      uints: '', // 商品单位
      origin: '', // 商品产地
      sKind: "请选择商品分类", // 或者将它初始化为第一个选项的值
      store: 0,
      price: 0,
      page: 1,
      items: [],
      modalData: {},
      showPopup: false,
      updateProcessData: {
        operater: '选择步骤',
        ifSelect: false
      },
      PopupData: {},
      kind: [],
      showUpdateModal: false, // 控制修改商品模态窗口显示
      showAddModal: false, // 控制添加商品模态窗口显示
      showProcessModal: false, // 控制查看步骤模态窗口显示
      errors: {
        name: '',
        description: '',
        stock: '',
        price: '',
        unit: '',
        origin: ''
      }
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.getAllKind()
    this.$axios.get("/main/get_manage_list", {
      params: {
        page: this.page,
      }
    }).then((res) => {
      if (res.data.code == 200) {
        this.items = res.data.products
      } else {
        this.$message.error("获取商品列表失败: " + res.data.message)
      }
      console.log(res);
    }).catch((res) => {
      console.log(res);
    })
  },
  methods: {
    // 触发
    openPopup(index) {
      this.PopupData = this.modalData.Process[index]
      this.showPopup = true;
    },
    closePopup() {
      this.showPopup = false;
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
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.$axios.get("/main/get_manage_list", {
          params: {
            page: this.page,
          }
        }).then((res) => {
          if (res.data.code == 200) {
            this.items = res.data.products
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

      if (this.items.length == 6) {
        this.page += 1;
      } else {
        this.$message.error('已经是最后一页了')
        return
      }
      this.$axios.get("/main/get_manage_list", {
        params: {
          page: this.page,
        }
      }).then((res) => {
        if (res.data.code == 200) {
          this.items = res.data.products
        } else {
          this.$message.error("获取商品列表失败: " + res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
    },
    openModel() {
      window.location.href = '#my_modal_8';
    },
    openBuy() {
      window.location.href = '#add'; // 使用锚点链接
    },
    openUpdate(id) {
      const item = this.items.find(item => item.Id === id);
      if (item) {
        this.modalData = { ...item };
      }
      window.location.href = '#update'; // 使用锚点链接
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
    handleFileUpload(event) {
      const file = event.target.files[0]; // 获取用户选择的文件
      if (file) {
        const reader = new FileReader();

        reader.onload = (e) => {
          // 调用文件上传务
          this.$axios.post("/main/upload", {
            fileData: e.target.result
          }).then((rep) => {
            if (rep.data.code == 200) {
              this.fileId = rep.data.id
            }
          }).catch((res) => {
            console.log(res);
          })
          console.log(e.target.result); // 这里是文件的Base64编码
          // 你可以在这里处理转换后的Base64字符串，比如将其发送到服务器
        };

        reader.readAsDataURL(file); // 将文件读取为Data URL
      }
    },
    // 添加商品接口
    addGoods() {
      console.log(this.sKind);
      if (this.name.length == 0 || this.store == 0 || this.price < 0) {
        this.$message.error('请填写合法信息!')
        return
      }

      this.$axios.post('/main/add_goods', {
        name: this.name,
        desc: this.desc,
        store: this.store,
        uint: this.uints,
        origin: this.origin,
        price: this.price,
        images: this.fileId,
        kind: this.sKind,
      }).then((res) => {
        if (res.data.code == 200) {
          window.location.href = '#';
          this.$message.success("添加商品成功: " + res.data.transactionHash)
          this.resetForm();
          this.getProductList();
        } else {
          this.$message.error("商品添加失败")
        }
      })
    },
    resetForm() {
      this.name = '';
      this.desc = '';
      this.store = 0;
      this.price = 0;
      this.uints = '';
      this.origin = '';
      this.sKind = "请选择商品分类";
      this.fileId = '';
    },
    getAllKind() {
      this.$axios.get("/main/get_all_fruit_kind").then((res) => {
        if (res.data.code === 200) {
          this.kind = res.data.fruits;
        } else {
          this.$message.error("获取商品种类失败: " + res.data.message)
        }
      })
    },
    // 修改商品信息
    updateGoods() {
      this.$axios.post("/main/update_goods_info", {
        id: this.modalData.Id,
        name: this.modalData.Name,
        desc: this.modalData.Description,
        store: this.modalData.Stock,
        price: this.modalData.Price,
        kind: this.modalData.Kind,
        origin: this.modalData.Origin,
        uint: this.modalData.UintGoods,
        images: this.fileId,
      }).then((res) => {
        window.location.href = '#';
        if (res.data.code == 200) {
          this.$message.success("修改商品成功: " + res.data.transactionHash)
        } else {
          this.$message.error("修改商品失败")
        }
      }).catch((res) => {
        console.log(res);
      })
    },
    removeGoods(id, status) {
      this.$axios.post("/main/remove_goods", {
        id: id,
        status: status,
      }).then((res) => {
        if (res.data.code == 200) {
          this.$message.success("商品更新状态成功: " + res.data.transactionHash)
        } else {
          this.$message.error("商品更新状态失败: " + res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
    },
    getStatusText(status) {
      var kind = ["进展中", "待审核", "上架", "拒绝", "下架"];
      var index = parseInt(status, 10); // 转换为整数，基数为10

      // 确保index在数组的边界内
      if (index < 0 || index >= kind.length) {
        console.error("Status index is out of bounds.");
        return "未知"; // 或其他适当的默认值
      }

      return kind[index];
    },
    // 添加商品进程
    addGoodsProcess(id) {
      const item = this.items.find(item => item.Id === id);
      if (item) {
        this.modalData = {...item};
      }
      window.location.href = '#addProcess';
    },
    openProcess(id) {
      const item = this.items.find(item => item.Id === id);
      if (item) {
        this.modalData = { ...item };
      }
      window.location.href = '#lookProcess';
    },
    // 更新进展
    updateProcess() {
      let status = 0;
      // 判断商品是否上架
      if (this.updateProcessData.ifSelect) {
        status = 1;
      }
      if (this.updateProcessData.operater.length === 0) {
        this.$message.error("操作不能为空!!!")
        return
      }

      this.$axios.post("/main/update_goods_process", {
        id: this.modalData.Id,
        imageId: this.fileId,
        operator: this.updateProcessData.operater,
        status: status
      }).then((res => {
        if (res.data.code === 200) {
          this.$message.success("添加商品成功: " + res.data.transactionHash);
          window.location.href = '#';
          return
        }
        console.log(res)
      })).catch((res) => {
        console.log(res)
      })
    },
    validateField(field) {
      switch (field) {
        case 'name':
          this.errors.name = !this.modalData.Name ? '商品名不能为空' :
            this.modalData.Name.length > 50 ? '商品名不能超过50个字符' : '';
          break;
        case 'description':
          this.errors.description = !this.modalData.Description ? '商品简介不能为空' :
            this.modalData.Description.length > 200 ? '商品简介不能超过200个字符' : '';
          break;
        case 'stock':
          this.errors.stock = !this.modalData.Stock ? '库存不能为空' :
            this.modalData.Stock < 0 ? '库存不能为负数' : '';
          break;
        case 'price':
          this.errors.price = !this.modalData.Price ? '单价不能为空' :
            this.modalData.Price < 0 ? '单价不能为负数' : '';
          break;
        case 'unit':
          this.errors.unit = !this.modalData.UintGoods ? '商品单位不能为空' : '';
          break;
        case 'origin':
          this.errors.origin = !this.modalData.Origin ? '商品产地不能为空' : '';
          break;
      }
    },

    // 验证所有字段
    validateForm() {
      this.validateField('name');
      this.validateField('description');
      this.validateField('stock');
      this.validateField('price');
      this.validateField('unit');
      this.validateField('origin');

      // 检查是否有错误
      return !Object.values(this.errors).some(error => error !== '');
    },

    // 验证并更新
    validateAndUpdate() {
      if (this.validateForm()) {
        this.updateGoods();
      } else {
        this.$message.error('请检查表单填写是否正确');
      }
    }
  }
}

</script>

<style scoped>
.container {
  max-width: 90%;
  margin: 0 auto;
  /* 居中 */
  padding: 20px;
  /* 内边距 */
  background-color: #f9f9f9;
  /* 背景颜色 */
  border-radius: 8px;
  /* 圆角 */
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  /* 阴影效果 */
}

.header {
  text-align: center;
  /* 标题居中 */
  margin-bottom: 20px;
  /* 下边距 */
}

.header h1 {
  font-size: 24px;
  /* 标题字体大小 */
  color: #333;
  /* 标题颜色 */
}

.table {
  width: 100%;
  /* 表格宽度 */
  border-collapse: collapse;
  /* 合并边框 */
  margin-bottom: 20px;
  /* 下边距 */
}

.table th,
.table td {
  padding: 12px;
  text-align: center;
  /* 设置所有单元格文字居中 */
  border-bottom: 1px solid #ddd;
  vertical-align: middle;
  /* 垂直居中 */
}

.table th {
  background-color: #f8b811;
  color: #fff;
  font-weight: 500;
}

.table tr:hover {
  background-color: rgba(248, 184, 17, 0.1);
}

/* 如果表格内有按钮，确保按钮也居中显示 */
.table td .button {
  margin: 0 5px;
  /* 按钮之间的间距 */
  display: inline-block;
  /* 使按钮保持在同一行 */
}

/* 确保表格内的其他元素也居中 */
.table td * {
  vertical-align: middle;
}

.button {
  background-color: #f8b811;
  /* 按钮背景颜色 */
  color: #fff;
  /* 按钮字体颜色 */
  padding: 10px 20px;
  /* 内边距 */
  border: none;
  /* 去掉边框 */
  border-radius: 5px;
  /* 圆角 */
  cursor: pointer;
  /* 鼠标悬浮时显示手型 */
  transition: background-color 0.3s;
  /* 添加过渡效果 */
}

.button:hover {
  background-color: #6fae3b;
  /* 鼠标悬浮时按钮颜色 */
}

.pagination {
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: center;
  /* 居中 */
  margin-top: 20px;
  /* 上边距 */
}

.pagination button {
  background-color: #f8b811;
  /* 分页按钮背景颜色 */
  color: #fff;
  /* 分页按钮字体颜色 */
  border: none;
  /* 去掉边框 */
  padding: 10px 15px;
  /* 内边距 */
  margin: 0 5px;
  /* 左右间距 */
  border-radius: 5px;
  /* 圆角 */
  cursor: pointer;
  /* 鼠标悬浮时显示手型 */
  transition: background-color 0.3s;
  /* 添加过渡效果 */
}

.pagination button:hover {
  background-color: #6fae3b;
}

.modal {
}

.modal-box {
  width: 500px;
  margin-right: 10px;
}

.scrollable-container {
  display: flex;
  justify-content: center;
  width: 100%;
  overflow-x: auto;
  white-space: nowrap;
}

.modal {
  vertical-align: top;
}

.modal-box {
  width: 500px;
  margin-right: 10px;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.modal-box h3 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  color: #333;
}

.grid.gap-2 {
  display: grid;
  gap: 15px;
  margin-top: 20px;
}

.grid.gap-2 label {
  font-weight: 500;
  color: #4a5568;
  margin-bottom: 5px;
}

.grid.gap-2 input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s;
}

.grid.gap-2 input:focus {
  outline: none;
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(128, 189, 68, 0.1);
}

.modal-action {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

.btn {
  padding: 8px 16px;
  border-radius: 6px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn:hover {
  opacity: 0.9;
}

.modal-action .btn:first-child {
  background-color: #f8b811;
  color: white;
}

.modal-action .btn:last-child {
  background-color: #e2e8f0;
  color: #4a5568;
}

ul {
  list-style: none;
  padding: 0;
}

ul li {
  padding: 10px;
  border-bottom: 1px solid #e2e8f0;
  color: #4a5568;
}

ul li:last-child {
  border-bottom: none;
}

.file-input {
  width: 100%;
  padding: 8px;
  margin-top: 15px;
  border: 1px dashed #f8b811;
  border-radius: 6px;
  cursor: pointer;
}

select {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  background-color: white;
  font-size: 14px;
  color: #4a5568;
}

.modal::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: -1;
}

@media (max-width: 640px) {
  .modal-box {
    width: 90%;
    margin: 20px;
  }
}

.error-message {
  color: red;
}

.steps-container {
  padding: 20px;
  margin: 20px 0;
}

.step-item {
  display: flex;
  align-items: flex-start;
  position: relative;
  padding: 10px 0;
}

.step-connector {
  width: 2px;
  height: 40px;
  background-color: #f8b811;
  margin-left: 19px;
}

.step-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f8b811;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-right: 15px;
  flex-shrink: 0;
  border: 2px solid #f8b811;
  transition: all 0.3s ease;
}

.step-content {
  flex-grow: 1;
  padding: 5px 15px;
  background-color: #f8f9fa;
  border-radius: 8px;
  margin-top: 5px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.step-content:hover {
  background-color: #e9ecef;
  transform: translateX(5px);
}

.step-content h4 {
  margin: 0;
  color: #2d3748;
  font-size: 1.1em;
  font-weight: 600;
}

.step-date {
  margin: 5px 0 0;
  color: #718096;
  font-size: 0.9em;
}

.completed {
  background-color: #f8b811;
  box-shadow: 0 0 0 4px rgba(128, 189, 68, 0.2);
}

.step-item:hover .step-circle {
  background-color: #e3ab1dab;
  box-shadow: 0 0 0 4px rgba(128, 189, 68, 0.2);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  position: relative;
}

.close-btn {
  position: absolute;
  right: -10px;
  top: -10px;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background-color: #fff;
  border: 2px solid #f8b811;
  color: #f8b811;
  font-size: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  text-decoration: none;
  transition: all 0.3s ease;
  line-height: 1;
}

.close-btn:hover {
  background-color: #f8b811;
  color: #fff;
  transform: rotate(90deg);
}

.modal-box h3 {
  margin-right: 30px;
}

.responsive-table {
  width: 100%;
  table-layout: fixed;
  border-collapse: collapse;
}

.w-8 {
  width: 8%;
}

.w-10 {
  width: 10%;
}

.w-15 {
  width: 15%;
}

.w-17 {
  width: 17%;
}

.table th,
.table td {
  padding: 12px 8px;
  text-align: center;
  border-bottom: 1px solid #ddd;
  vertical-align: middle;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

@media screen and (max-width: 1024px) {
  .responsive-table {
    display: block;
    overflow-x: auto;
    white-space: nowrap;
  }
}

.table td .button {
  padding: 8px 12px;
  margin: 2px;
  white-space: nowrap;
}

.table td {
  min-width: 60px;
}

.table td:nth-child(2) {
  white-space: normal;
  word-break: break-word;
}

@media screen and (min-width: 1920px) {
  .container {
    max-width: 1800px;
  }
}

@media screen and (max-width: 768px) {
  .container {
    max-width: 95%;
    padding: 10px;
  }
}

/* 添加步骤模态窗口样式 */
.form-section {
  margin-top: 20px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.form-label {
  display: block;
  font-weight: bold;
  margin-bottom: 8px;
  color: #2F3529;
}

.select-primary {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 16px;
  background: white;
}

.status-toggle {
  margin: 16px 0;
}

.toggle {
  margin-top: 8px;
}

.image-upload {
  margin-top: 16px;
}

.file-input {
  width: 100%;
  padding: 8px;
  border: 1px dashed #f8b811;
  border-radius: 6px;
  background: white;
}

/* 步骤展示样式 */
.steps-container {
  padding: 20px;
  margin: 20px 0;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.step-item {
  display: flex;
  align-items: flex-start;
  position: relative;
  padding: 10px 0;
}

.step-connector {
  width: 2px;
  height: 40px;
  background-color: #f8b811;
  margin-left: 19px;
}

.step-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f8b811;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-right: 15px;
  flex-shrink: 0;
}

.step-content {
  flex-grow: 1;
  padding: 10px 15px;
  background-color: #f8f9fa;
  border-radius: 8px;
  margin-top: 5px;
}

.step-content h4 {
  margin: 0;
  color: #2F3529;
  font-weight: 600;
}

.step-date {
  margin: 5px 0 0;
  color: #718096;
  font-size: 0.9em;
}
</style>
