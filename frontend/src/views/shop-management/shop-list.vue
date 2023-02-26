<template>
	<section class="app-container">
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters" @submit.native.prevent>
				<el-form-item label="商店名称:">
					<el-input v-model="filters.name" placeholder="请输入商店名称"></el-input>
				</el-form-item>
        <el-form-item label="商圈名称:">
          <el-input v-model="filters.area" placeholder="请输入商圈名称"></el-input>
        </el-form-item>
        <el-form-item label="详细地址:">
          <el-input v-model="filters.address" placeholder="请输入地址"></el-input>
        </el-form-item>
        <el-form-item label="商店类型:">
          <el-select v-model = "filters.type_id" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="美食"/>
            <el-option key:2 value="2" label="KTV"/>
            <el-option key:3 value="3" label="丽人·美发"/>
            <el-option key:4 value="4" label="健身运动"/>
            <el-option key:5 value="5" label="按摩·足疗"/>
            <el-option key:6 value="6" label="美容SPA"/>
            <el-option key:7 value="7" label="亲子游乐"/>
            <el-option key:8 value="8" label="酒吧"/>
          </el-select>
        </el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getShops">查询</el-button>
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="handleAdd">新增</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="shops" highlight-current-row @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column type="index" width="60">
      </el-table-column>
      <el-table-column prop="name" label="名称" width="120">
      </el-table-column>
			<el-table-column prop="type_id" label="商铺类型" width="120">
			</el-table-column>
			<el-table-column prop="area" label="商圈" width="120">
			</el-table-column>
			<el-table-column prop="address" label="详细地址" min-width="160">
			</el-table-column>
      <el-table-column prop="open_hours" label="营业时间" min-width="160">
      </el-table-column>
        <el-table-column prop="sold" label="销量" min-width="160">
        </el-table-column>
        <el-table-column prop="avg_price" label="人均价" min-width="160">
        </el-table-column>
        <el-table-column prop="score" label="评分" min-width="160">
        </el-table-column>
        <el-table-column prop="comments" label="评论数" min-width="160">
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
        <el-form-item label="商店名称" prop="name">
          <el-input v-model="editForm.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item v-if="dialogStatus=='create'" label="位置" prop="password">
          <el-input v-model="editForm.location" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="商店图" prop="images">
          <el-upload
            action="https://jsonplaceholder.typicode.com/posts/"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :on-remove="handleRemove"
            :on-change="handlePictureChange">
            <i class="el-icon-plus"></i>
          </el-upload>
          <el-dialog :visible.sync="dialogVisible">
            <img width="100%" :src="dialogImageUrl" alt="">
          </el-dialog>
        </el-form-item>

        <el-form-item label="商圈名称" prop="area">
          <el-input v-model="editForm.area" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="详细地址" prop="address">
          <el-input v-model="editForm.address" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="营业时间" prop="open_hours">
          <el-input v-model="editForm.open_hours" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="均价" prop="avg_price">
          <el-input v-model="editForm.avg_price" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="商店类型">
          <el-select v-model = "editForm.type_id" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="美食"/>
            <el-option key:2 value="2" label="KTV"/>
            <el-option key:3 value="3" label="丽人·美发"/>
            <el-option key:4 value="4" label="健身运动"/>
            <el-option key:5 value="5" label="按摩·足疗"/>
            <el-option key:6 value="6" label="美容SPA"/>
            <el-option key:7 value="7" label="亲子游乐"/>
            <el-option key:8 value="8" label="酒吧"/>
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
import util from '@/utils/table.js'
import {
  getShopList,
  removeShop,
  batchRemoveShop,
  updateShop,
  addShop,
} from '@/api/shop'

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
        type_id:'',
        area:'',
        address: '',
      },
      shops: [],
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
        name: '',
        type_id:1,
        images:'',
        area:'',
        address: '',
        sold: 1,
        avg_price: 1,
        comments : 1,
        score:1.5,
        open_hours:''
      },

      addFormVisible: false, // 新增界面是否显示
      addFormRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
      },
      dialogImageUrl: '',
      dialogVisible: false
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.page = val
      this.getShops()

    },
    // 获取用户列表
    getShops() {
      const para = {
        name: this.filters.name,
        type_id:this.filters.type_id,
        area:this.filters.area,
        address: this.filters.address,
        page: this.page,
        page_size:this.page_size,

      }
      getShopList(para).then(res => {
        console.log("查看记录", res)
        this.total = res.data.total
        this.shops = res.data.list
      })
    },
    // 删除
    handleDel(index, row) {
      this.$confirm('确认删除该记录吗?', '提示', {
        type: 'warning'
      })
        .then(() => {
          const para = { id: row.id }
          removeShop(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getShops()
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
              updateShop(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getShops()
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
              addShop(para).then(res => {
                this.$message({
                  message: '提交成功',
                  type: 'success'
                })
                this.$refs['editForm'].resetFields()
                this.dialogFormVisible = false
                this.getShops()
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
          batchRemoveShop(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getShops()
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
    }
  },
  mounted() {
    this.getShops()
  }
}
</script>

<style>

</style>
