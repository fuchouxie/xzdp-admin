<template>
  <div class="container">
    <section class="search-area">
      <p class="title">用户详情</p>
    </section>
    <section class="table-area">
      <el-descriptions title="用户信息" border direction="vertical" :column="6">
        <el-descriptions-item label="昵称">{{ user.nick_name }}</el-descriptions-item>
        <el-descriptions-item label="手机号">{{ user.phone }}</el-descriptions-item>
        <el-descriptions-item label="个人介绍">{{ user.introduce }}</el-descriptions-item>
        <el-descriptions-item label="城市">{{ user.city }}</el-descriptions-item>
        <el-descriptions-item label="会员等级">{{ user.level }}</el-descriptions-item>

        <el-descriptions-item label="粉丝数">{{ user.fans }}</el-descriptions-item>
        <el-descriptions-item label="关注数">{{ user.followee }}</el-descriptions-item>
        <el-descriptions-item label="性别">{{ genderMap[user.gender] }}</el-descriptions-item>
        <el-descriptions-item label="生日">{{ user.birthday }}</el-descriptions-item>
        <el-descriptions-item label="积分">{{ user.credits }}</el-descriptions-item>
      </el-descriptions>

      <el-row type="flex" justify="center">
        <el-col :span="1">
          <el-button plain @click="goBack">返 回</el-button>
        </el-col>
      </el-row>
    </section>
  </div>
</template>
<script>

import {getUserInfo} from '@/api/user'

export default {
  name: 'user-detail',
  data() {
    return {
      fullLoading: null,
      user: [],
      id: this.$route.params.id,
      genderMap: [
        "男","女"
      ],
    }
  },
  mounted() {
    this.getUserDetail()
  },
  methods:{
    openFullscreenLoading(options = {}) {
      this.fullLoading = this.$loading(
        Object.assign({ lock: true, text: "加载中..." }, options)
      );
    },
    closeFullscreenLoading() {
      this.fullLoading.close();
    },
    getUserDetail() {
      this.openFullscreenLoading()
      getUserInfo({id: this.id}).then((res) => {
        this.user = res.data || {}
        this.closeFullscreenLoading()
        console.log(this.user)
      }).catch(() => this.closeFullscreenLoading())
    },
    goBack() {
      this.$router.push('/user/userList')
    }

  },

}
</script>

<style lang="scss" scoped>
.container {
  padding: 15px 26px;
  background-color: #F4F6F9;
  position: relative;
}

.search-area {
  width: 100%;
  background-color: #fff;
  padding: 20px 0;
// margin-bottom: 30px;
  border-radius: 10px 10px 0 0;
  box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
}

.title {
  font-size: 20px;
  font-weight: bold;
  margin: 0 0 15px 15px;
}

.table-area {
  width: 100%;
  background-color: #fff;
  padding: 20px;
  border-radius: 0 0 10px 10px;
  box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.1);
}

.table-area-div {
  margin-top: 40px;
}
</style>
