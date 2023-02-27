<template>
	<section class="app-container">
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters" @submit.native.prevent>
				<el-form-item label="用户:">
					<el-input v-model="filters.user_id" placeholder="请输入用户名"></el-input>
				</el-form-item>
        <el-form-item label="商店:">
          <el-input v-model="filters.shop_id" placeholder="请输入名称"></el-input>
        </el-form-item>
        <el-form-item label="标题:">
          <el-input v-model="filters.title" placeholder="请输入联系方式"></el-input>
        </el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getBlogs">查询</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="blogs" highlight-current-row @selection-change="selsChange" style="width: 100%;">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column type="index" width="60">
			</el-table-column>
			<el-table-column prop="user_name" label="用户名" width="120">
			</el-table-column>
      <el-table-column prop="shop_name" label="商店" width="120">
      </el-table-column>
			<el-table-column prop="title" label="标题" width="160">
			</el-table-column>
			<el-table-column prop="liked" label="点赞数" width="120">
			</el-table-column>
			<el-table-column prop="comments" label="评论数" min-width="120">
			</el-table-column>
      <el-table-column prop="created_time" label="创建时间" min-width="160">
      </el-table-column>
			<el-table-column label="操作" width="150">
				<template slot-scope="scope">
					<el-button size="small" @click="blogDetail(scope.$index, scope.row)">详情</el-button>
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

	</section>
</template>

<script>
import util from '@/utils/table.js'
import {
  getBlogList,
  remove,
  batchRemove
} from '@/api/blog'

export default {
  data() {
    return {
      dialogStatus: '',
      filters: {
        title: '',
        shop_id: '',
        user_id:'',
      },
      blogs: [],
      total: 0,
      page: 1,
      page_size:10,
      sels: [], // 列表选中列
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.page = val
      this.getBlogs()

    },
    // 获取用户列表
    getBlogs() {
      const para = {
        page: this.page,
        page_size:this.page_size,
        shop_id: this.filters.shop_id,
        title: this.filters.title,
        user_id: this.filters.user_id,
      }
      getBlogList(para).then(res => {
        console.log("查看记录", res)
        this.total = res.data.total
        this.blogs = res.data.list
      })
    },
    // 删除
    handleDel(index, row) {
      this.$confirm('确认删除该记录吗?', '提示', {
        type: 'warning'
      })
        .then(() => {
          const para = { id: row.id }
          remove(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getBlogs()
          })
        })
        .catch(() => {})
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
          batchRemove(para).then(res => {
            this.$message({
              message: '删除成功',
              type: 'success'
            })
            this.getBlogs()
          })
        })
        .catch(() => {})
    },
    blogDetail(index, row) {
      this.$router.push('/blog/blogDetail/' + row.id)
    },
  },
  mounted() {
    this.getBlogs()
  }
}
</script>

<style scoped>

</style>
