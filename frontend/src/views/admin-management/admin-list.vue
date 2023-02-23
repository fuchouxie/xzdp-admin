<template>
	<section class="app-container">
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters" @submit.native.prevent>
				<el-form-item label="用户名:">
					<el-input v-model="filters.username" placeholder="请输入用户名"></el-input>
				</el-form-item>
        <el-form-item label="名称:">
          <el-input v-model="filters.name" placeholder="请输入名称"></el-input>
        </el-form-item>
        <el-form-item label="联系方式:">
          <el-input v-model="filters.phone" placeholder="请输入联系方式"></el-input>
        </el-form-item>
        <el-form-item label="角色:">
          <el-select v-model = "filters.role_id" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="超级管理员"/>
            <el-option key:2 value="2" label="运营管理员"/>
          </el-select>
        </el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getAdminUsers">查询</el-button>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="handleAdd">新增</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="adminUsers" highlight-current-row @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column type="index" width="60">
			</el-table-column>
			<el-table-column prop="username" label="用户名" width="120">
			</el-table-column>
      <el-table-column prop="name" label="名称" width="120">
      </el-table-column>
			<el-table-column prop="phone" label="联系方式" width="120">
			</el-table-column>
			<el-table-column prop="role_id" label="角色" width="120">
			</el-table-column>
			<el-table-column prop="remark" label="备注" min-width="160">
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
			<el-pagination layout="prev, pager, next" @current-change="handleCurrentChange" :page-size="20" :total="total" style="float:right;">
			</el-pagination>
		</el-col>

		<!--编辑界面-->
		<el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
			<el-form :model="editForm" label-width="80px" :rules="editFormRules" ref="editForm">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="editForm.username" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item v-if="dialogStatus=='create'" label="密码" prop="password">
          <el-input v-model="editForm.password" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="editForm.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="editForm.remark" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="联系方式" prop="phone">
          <el-input v-model="editForm.phone" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="角色">
          <el-radio-group v-model="editForm.role_id">
            <el-radio class="radio" :label=1 value="1">超级管理员</el-radio>
            <el-radio class="radio" :label=2 value="2">运营管理员</el-radio>
          </el-radio-group>
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
import util from '@/utils/table.js'
import {
  getAdminList,
  removeAdminUser,
  batchRemoveAdminUser,
  updateAdminUser,
  addAdminUser,
} from '@/api/admin'

export default {
  data() {
    return {
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '新建'
      },
      dialogFormVisible: false,
      filters: {
        name: '',
        username: '',
        role_id:'',
        phone:''
      },
      adminUsers: [],
      total: 0,
      page: 1,
      sels: [], // 列表选中列
      editFormRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
        username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
        password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
      },
      // 编辑界面数据
      editForm: {
        id: '0',
        name: '',
        username:'',
        password:'',
        phone: 1,
        remark: '',
        role_id: 1
      },

      addFormVisible: false, // 新增界面是否显示
      addFormRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
      }
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.page = val
      this.getAdminUsers()

    },
    // 获取用户列表
    getAdminUsers() {
      const para = {
        page: this.page,
        username: this.filters.username,
        phone: this.filters.phone,
        role_id: this.filters.role_id,
        name: this.filters.name,
      }
      getAdminList(para).then(res => {
        console.log("查看记录", res)
        this.total = res.data.total
        this.adminUsers = res.data
      })
    },
    // 删除
    handleDel(index, row) {
      this.$confirm('确认删除该记录吗?', '提示', {
        type: 'warning'
      })
        .then(() => {
          const para = { id: row.id }
          removeAdminUser(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getAdminUsers()
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
        name: '',
        sex: 1,
        age: 0,
        birth: '',
        addr: ''
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
              updateAdminUser(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getAdminUsers()
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
              addAdminUser(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getAdminUsers()
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
          batchRemoveAdminUser(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getAdminUsers()
          })
        })
        .catch(() => {})
    }
  },
  mounted() {
    this.getAdminUsers()
  }
}
</script>

<style scoped>

</style>
