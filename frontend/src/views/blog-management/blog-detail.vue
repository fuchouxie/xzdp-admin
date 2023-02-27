<template>
  <div class="container">
    <section class="search-area">
      <p class="title">笔记详情</p>
    </section>
    <section class="table-area">
      <el-descriptions title="笔记信息" border direction="vertical" :column="6">
        <el-descriptions-item label="昵称">{{ blog.user_name }}</el-descriptions-item>
        <el-descriptions-item label="商店名称">{{ blog.shop_name }}</el-descriptions-item>

        <el-descriptions-item label="点赞数">{{ blog.liked }}</el-descriptions-item>
        <el-descriptions-item label="评论数">{{ blog.comments }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ blog.created_time }}</el-descriptions-item>

      </el-descriptions>
      <el-descriptions border direction="vertical" :column="6">
      <el-descriptions-item white-space="pre-wrap" label="内容" v-html="blog.content" >{{ blog.content}}</el-descriptions-item>
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

import {getBlogInfo} from '@/api/blog'

export default {
  name: 'blog-detail',
  data() {
    return {
      fullLoading: null,
      blog: [],
      id: this.$route.params.id,
      genderMap: [
        "男","女"
      ],
    }
  },
  mounted() {
    this.getBlogDetail()
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
    getBlogDetail() {
      this.openFullscreenLoading()
      getBlogInfo({id: this.id}).then((res) => {
        this.blog = res.data || {}
        this.closeFullscreenLoading()
        console.log(this.blog)
      }).catch(() => this.closeFullscreenLoading())
    },
    goBack() {
      this.$router.push('/blog/blogList')
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
  white-space : pre-line;
}

.table-area-div {
  margin-top: 40px;
}
</style>
