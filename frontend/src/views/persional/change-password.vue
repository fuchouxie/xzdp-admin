<template>
  <div class="container">
    <main>
      <div class="box">
        <el-form label-width="80px" :model="form" :rules="rules" ref="ruleForm">
          <el-form-item label="用户名" prop="username">
            <el-input disabled v-model="user.username"></el-input>
          </el-form-item>
          <el-form-item label="旧密码" prop="old_password">
            <el-input
              type="password"
              maxlength="16"
              placeholder="请输入旧密码"
              v-model="form.old_password"
            ></el-input>
          </el-form-item>
          <el-form-item label="新密码" prop="new_password">
            <el-input
              type="password"
              maxlength="16"
              placeholder="请输入新密码"
              v-model="form.new_password"
            ></el-input>
          </el-form-item>
          <el-form-item label="确认密码" prop="new_password2">
            <el-input
              type="password"
              maxlength="16"
              placeholder="请确认新密码"
              v-model="form.new_password2"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')"
              >提交</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </main>
  </div>
</template>
<script>
import { changePassword } from "@/api/admin";
export default {
  data() {
    var checkPassword = (rule, value, callback) => {
      if (value.length < 8 || value.length > 16) {
        callback(new Error("密码长度8-16位"));
        return;
      }

      if (
        !/(?!^[0-9]+$)(?!^[A-z]+$)(?!^[^A-z0-9]+$)^[^\s\u4e00-\u9fa5]{8,16}$/.test(
          value
        )
      ) {
        callback(new Error("密码必须字母、数字和标点至少包含两种"));
        return;
      }
      callback();
    };

    var checkPassword2 = (rule, value, callback) => {
      if (value.length < 8 || value.length > 16) {
        callback(new Error("密码长度8-16位"));
        return;
      }
      if (value !== this.form.new_password) {
        callback(new Error("密码不一致"));
        return;
      }

      if (
        !/(?!^[0-9]+$)(?!^[A-z]+$)(?!^[^A-z0-9]+$)^[^\s\u4e00-\u9fa5]{8,16}$/.test(
          value
        )
      ) {
        callback(new Error("密码必须字母、数字和标点至少包含两种"));
        return;
      }
      callback();
    };

    return {
      user:{
        username:'admin',
      },
      form: {
        username: "zhangsan",
        old_password: "",
        new_password: "",
        new_password2: "",
      },
      rules: {
        old_password: [
          {
            required: true,
            message: "请输入旧密码",
            trigger: "blur",
          },
        ],
        new_password: [{ validator: checkPassword, trigger: "change" }],
        new_password2: [{ validator: checkPassword2, trigger: "blur" }],
      },
    };
  },
  created() {},
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          changePassword(this.form).then((res) => {
            if (res.code === 1) {
              this.$message({
                message: "密码修改成功",
                type: "success",
              });
              this.$router.back();
            } else {
              this.$message.error(res.msg || "密码修改失败");
            }
          });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.container {
  padding: 15px 26px;
  background-color: #F4F6F9;
  position: relative;
}
.box {
  width: 400px;
}
main {
  background-color: #fff;
  width: 100%;
  min-height: 80vh;
  border-radius: 15px;
  box-shadow: 0px 2px 8px 0px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  .item {
    width: 300px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 24px;
    .label {
      flex-shrink: 0;
      margin-right: 16px;
      width: 80px;
      text-align: right;
    }
  }
}
</style>
