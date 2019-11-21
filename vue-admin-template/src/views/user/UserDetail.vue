<template>
  <div class="app-container">
    <el-form ref="user" :model="user" :rules="rules">


      <el-form-item prop="username">
        <MDinput name="name" v-model="user.username" :maxlength="20" :readonly="isEdit">
          Username
        </MDinput>
      </el-form-item>

      <el-form-item prop="password">
        <MDinput name="name" v-model="user.password" :maxlength="64" type="password">
          Password
        </MDinput>
      </el-form-item>

      <el-form-item prop="email">
        <MDinput name="name" v-model="user.email" :maxlength="20" type="email" required>
          Email
        </MDinput>
      </el-form-item>

      <el-form-item>
        <el-col :span="11">
          <label>Active</label>&nbsp;&nbsp;<el-switch v-model="user.active"/>
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="onClear">清空</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import MDinput from '@/components/MDinput'
  import {createUser, editUser, fetchUser} from '@/api/user'

  const defaultForm = {
    username: '',
    password: '',
    email: '',
    active: true,
  }
  export default {
    name: 'userDetail',
    components: {MDinput},
    props: {
      isEdit: {
        type: Boolean,
        default: false
      }
    },
    data() {
      const validateRequire = (rule, value, callback) => {
        if (value === '') {
          this.$message({
            message: rule.field + '为必传项',
            type: 'error'
          })
          callback(new Error(rule.field + '为必传项'))
        } else {
          callback()
        }
      }

      const validatePassword = (rule, value, callback) => {
        if ((!this.isEdit) && (value === '')) {
          this.$message({
            message: rule.field + '为必传项',
            type: 'error'
          })
          callback(new Error('密码为必传项'))
        } else if (this.isEdit && value !== '') {
          if (value.length <= 3) {
            this.$message({
              message: '密码需大于三位',
              type: 'error'
            })
            callback(new Error('密码需大于三位'))
          } else {
            callback()
          }
        } else {
          callback()
        }
      }
      return {
        user: Object.assign({}, defaultForm),
        rules: {
          username: [{validator: validateRequire}],
          email: [{validator: validateRequire}],
          password: [{validator: validatePassword}],
        },
      }
    },
    created() {
      if (this.isEdit) {
        this.fetchData()
      }
    },
    methods: {
      submitForm() {
        this.$refs.user.validate(valid => {
          if (valid) {
            this.onSubmit()
          } else {
            return false
          }
        })
      },
      onSubmit() {
        if (this.isEdit) {
          editUser(this.user).then(response => {
            if (response.code === 20000) {
              this.$notify({
                title: '成功',
                message: '已编辑',
                type: 'success',
                duration: 2000
              })
            }
          })
        } else {
          createUser(this.user).then(response => {
            if (response.code === 20000) {
              this.$notify({
                title: '成功',
                message: '已创建',
                type: 'success',
                duration: 2000
              })
              Object.assign(this.user, defaultForm)
            }
          })
        }
      },
      onClear() {
        Object.assign(this.user, defaultForm)
      },
      fetchData() {
        fetchUser(this.$route.query.id).then(response => {
          const user = response.data
          Object.assign(this.user, user)
        })
      }
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .editor-container {
    min-height: 500px;
    margin: 0 0 30px;

  .editor-upload-btn-container {
    text-align: right;
    margin-right: 10px;

  .editor-upload-btn {
    display: inline-block;
  }

  }
  }
</style>
