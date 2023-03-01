<template>
  <section class="app-container">
    <!--工具条-->
    <el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
      <el-form :inline="true" :model="filters" @submit.native.prevent>
        <el-form-item label="商店:">
          <el-input v-model="filters.shop_name" placeholder="请输入商店名称"></el-input>
        </el-form-item>
        <el-form-item label="标题:">
          <el-input v-model="filters.title" placeholder="请输入标题"></el-input>
        </el-form-item>
        <el-form-item label="上架状态:">
          <el-select v-model = "filters.status" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="上架"/>
            <el-option key:2 value="2" label="下架"/>
            <el-option key:3 value="3" label="过期"/>
          </el-select>
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model = "filters.type" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="普通券"/>
            <el-option key:2 value="2" label="秒杀券"/>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" v-on:click="getVouchers">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleAdd">新增</el-button>
        </el-form-item>
      </el-form>
    </el-col>

    <el-table :data="vouchers" highlight-current-row @selection-change="selsChange" style="width: 100%;">
      <el-table-column type="selection" width="55">
      </el-table-column>
      <el-table-column type="index" width="60">
      </el-table-column>
      <el-table-column prop="shop_name" label="商店" width="120">
      </el-table-column>
      <el-table-column prop="title" label="标题" width="120">
      </el-table-column>
      <el-table-column prop="sub_title" label="副标题" width="120">
      </el-table-column>
      <el-table-column prop="pay_value" label="支付金额" min-width="160">
      </el-table-column>
      <el-table-column prop="actual_value" label="抵扣金额" min-width="160">
      </el-table-column>
      <el-table-column prop="type" label="类型" :formatter="typeFormat" min-width="160">
      </el-table-column>
      <el-table-column prop="status" label="状态" :formatter="statusFormat" min-width="160" >
      </el-table-column>
      <el-table-column prop="created_time" label="创建时间" min-width="160">
      </el-table-column>
      <el-table-column prop="rules" label="使用规则" min-width="160">
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template slot-scope="scope">
          <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDel(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--工具条-->
    <el-col :span="24" class="toolbar">
      <el-button type="danger" @click="batchRemove" :disabled="this.sels.length===0">批量删除</el-button>
      <el-pagination layout="prev, pager, next" @current-change="handleCurrentChange" :page-size=this.page_size :total="total" style="float:right;">
      </el-pagination>
    </el-col>

    <!--编辑界面-->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
      <el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
        <el-form-item label="商店名称:" prop="shop_id">
          <el-input v-model="editForm.shop_id" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="主标题:" prop="title">
          <el-input v-model="editForm.title" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="副标题:" prop="sub_title">
          <el-input v-model="editForm.sub_title" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="支付金额:" prop="pay_value">
          <el-input v-model="editForm.pay_value" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="抵扣金额:" prop="actual_value">
          <el-input v-model="editForm.actual_value" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="使用规则:" prop="rules">
          <el-input v-model="editForm.rules" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model = "editForm.status"  placeholder="请选择" clearable>
            <el-option key:1 value="1" label="上架"/>
            <el-option key:2 value="2" label="下架"/>
            <el-option key:3 value="3" label="停用"/>
          </el-select>
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model = "editForm.type" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="普通"/>
            <el-option key:2 value="2" label="秒杀"/>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click.native="dialogFormVisible=false">取消</el-button>
        <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">添加</el-button>
        <el-button v-else type="primary" @click="updateData">修改</el-button>
      </div>
    </el-dialog>
  </section>
</template>

<script>
import {addVoucher, batchRemoveVoucher, getVoucherList, removeVoucher, updateVoucher} from "@/api/voucher";

export default {
  name: "voucher-list",
  data() {
    return {
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '新建'
      },
      dialogFormVisible: false,
      filters: {
        shop_name: '',
        status:'',
        title:'',
        type:''
      },
      vouchers: [],
      total: 0,
      page: 1,
      page_size:10,
      sels: [], // 列表选中列
      editFormRules: {
        name: [{ required: true, message: '请输入商店名称', trigger: 'blur' }],
      },
      // 编辑界面数据
      editForm: {
        id: '0',
        shop_id: '',
        title:1,
        sub_title:'',
        rules:'',
        pay_value: '',
        actual_value: 1,
        type: 1,
        status : 1,
      },
      addFormVisible: false, // 新增界面是否显示
      addFormRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
      },
      statusMap:[
        "上架",
        "下架",
        "停用"
      ],
      typeMap:[
        "普通券",
        "秒杀券"
      ]
    }
  },
  methods:{
    handleCurrentChange(val) {
      this.page = val
      this.getVouchers()

    },
    getVouchers() {
      const para = {
        shop_name: this.filters.shop_name,
        title:this.filters.title,
        status:this.filters.status,
        type : this.filters.type,
        page: this.page,
        page_size:this.page_size,
      }
      getVoucherList(para).then(res => {
        console.log("查看记录", res)
        this.total = res.data.total
        this.vouchers = res.data.list
      })
    },
    // 删除
    handleDel(index, row) {
      this.$confirm('确认删除该记录吗?', '提示', {
        type: 'warning'
      })
        .then(() => {
          const para = { id: row.id }
          removeVoucher(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getVouchers()
          })
        })
        .catch(() => {})
    },
    // 显示编辑界面
    handleEdit(index, row) {
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.editForm = Object.assign({}, row)
    },
    // 显示新增界面
    handleAdd() {
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.editForm = {
        id: '0',

      }
    },
    // 编辑
    updateData() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          this.$confirm('确认提交吗？', '提示', {})
            .then(() => {
              const para = Object.assign({}, this.editForm)
              console.log(para)
              updateVoucher(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getVouchers()
              })
            })
            .catch(e => {
              // 打印一下错误
              console.log(e)
            })
        }
      })
    },
    // 新增
    createData: function() {
      this.$refs.editForm.validate(valid => {
        if (valid) {
          this.$confirm('确认提交吗？', '提示', {})
            .then(() => {
              const para = Object.assign({}, this.editForm)
              console.log(para)
              addVoucher(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getVouchers()
              })
            })
            .catch(e => {
              // 打印一下错误
              console.log(e)
            })
        }
      })
    },
    // 全选单选多选
    selsChange(sels) {
      this.sels = sels
    },
    // 批量删除
    batchRemove() {
      var ids = this.sels.map(item => item.id).toString()
      this.$confirm('确认删除选中记录吗？', '提示', {
        type: 'warning'
      })
        .then(() => {
          const para = { ids: ids }
          batchRemoveVoucher(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getVouchers()
          })
        })
        .catch(() => {})
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    handlePictureChange(file){
      console.log("icon发生改变");
      this.editForm.images = file.url;
    },
    statusFormat (val) {
      return this.statusMap[val.status-1]
    },
    typeFormat (val) {
      return this.typeMap[val.type-1]
    }
  },
  mounted() {
    this.getVouchers()
  }
}
</script>

<style scoped>

</style>
