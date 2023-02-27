<template>
	<section class="app-container">
		<!--工具条-->
		<el-col :span="24" class="toolbar" style="padding-bottom: 0px;">
			<el-form :inline="true" :model="filters" @submit.native.prevent>
				<el-form-item label="用户名:">
					<el-input v-model="filters.user_id" placeholder="请输入用户"></el-input>
				</el-form-item>
        <el-form-item label="订单状态:">
          <el-select v-model = "filters.status" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="未支付"/>
            <el-option key:2 value="2" label="已支付"/>
            <el-option key:3 value="3" label="已核销"/>
            <el-option key:4 value="4" label="已取消"/>
            <el-option key:5 value="5" label="退款中"/>
            <el-option key:6 value="6" label="已退款"/>
          </el-select>
        </el-form-item>
        <el-form-item label="支付方式:">
          <el-select v-model = "filters.pay_type" placeholder="请选择" clearable>
            <el-option key:1 value="1" label="余额"/>
            <el-option key:2 value="2" label="微信"/>
            <el-option key:3 value="3" label="支付宝"/>
          </el-select>
        </el-form-item>
				<el-form-item>
					<el-button type="primary" v-on:click="getOrders">查询</el-button>
				</el-form-item>
			</el-form>
		</el-col>

		<!--列表-->
		<el-table :data="orders" highlight-current-row style="width: 100%;">
			<el-table-column type="selection" width="55">
			</el-table-column>
			<el-table-column type="index" width="60">
			</el-table-column>
			<el-table-column prop="nick_name" label="用户名" width="120">
			</el-table-column>
      <el-table-column prop="voucher_title" label="优惠券" width="120">
      </el-table-column>
      <el-table-column prop="pay_type" label="支付方式" :formatter="payTypeFormat" width="120" >
      </el-table-column>
			<el-table-column prop="status" label="订单状态" :formatter="statusFormat" width="120" >
			</el-table-column>
			<el-table-column prop="created_time" label="创建时间" width="120">
			</el-table-column>
			<el-table-column prop="pay_time" label="支付时间" min-width="160">
			</el-table-column>
      <el-table-column prop="use_time" label="核销时间" min-width="160">
      </el-table-column>
      <el-table-column prop="refund_time" label="退款时间" min-width="160">
      </el-table-column>
			<el-table-column label="操作" width="150">
<!--				<template slot-scope="scope">-->
<!--					<el-button size="small" @click="handleEdit(scope.$index, scope.row)">详情</el-button>-->
<!--				</template>-->
			</el-table-column>
		</el-table>

	</section>
</template>

<script>
import util from '@/utils/table.js'
import {
  getOrderList,
  getOrderInfo
} from '@/api/order'

export default {
  data() {
    return {
      dialogStatus: '',
      filters: {
        user_id: '',
        pay_type:'',
        status:''
      },
      orders: [],
      total: 0,
      page: 1,
      page_size:10,
      sels: [], // 列表选中列
      addFormRules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }]
      },
      payType:[
        "余额",
        "支付宝",
        "微信"
      ],
      statusType:[
        "未支付",
        "已支付",
        "已核销",
        "已取消",
        "退款中",
        "已退款",
      ]
    }
  },
  methods: {
    handleCurrentChange(val) {
      this.page = val
      this.getOrders()

    },
    // 获取用户列表
    getOrders() {
      const para = {
        page: this.page,
        page_size:this.page_size,
        pay_type: this.filters.pay_type,
        status: this.filters.status,
        user_id: this.filters.user_id,
      }
      getOrderList(para).then(res => {
        console.log("查看记录", res)
        this.total = res.data.total
        this.orders = res.data.list
      })
    },
    payTypeFormat(val){
      return this.payType[val.pay_type-1]
    },
    statusFormat(val){
      return this.statusType[val.status-1]
    }
  },
  mounted() {
    this.getOrders()
  }
}
</script>

<style scoped>

</style>
