<template>
  <div class="app-container">
    <el-form ref="form" :model="article">
      <el-form-item>
        <el-col :span="11">
          <MDinput name="name" v-model="article.title" required :maxlength="100">
            标题
          </MDinput>
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <MDinput name="name" v-model="article.slug" :maxlength="100">
            Slug
          </MDinput>
        </el-col>
      </el-form-item>

      <el-form-item style="margin-bottom: 40px;" prop="summary">
        <MDinput name="name" v-model="article.summary" :maxlength="255">
          Summary
        </MDinput>
      </el-form-item>

      <div class="editor-container">
        <MarkdownEditor ref="editor" v-model="article.content"></MarkdownEditor>
      </div>

      <el-form-item>
        <el-col :span="11">
          <label>Tags</label>&nbsp;&nbsp;
          <el-select
            v-model="article.tags"
            multiple
            filterable
            allow-create
            default-first-option
            style="width: 90%"
            placeholder="请选择文章标签">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <label>Created At</label>&nbsp;&nbsp;
          <el-date-picker
            v-model="article.createdAt"
            type="datetime"
            placeholder="选择日期时间">
          </el-date-picker>
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-col :span="11">
          <label>Can Comment</label>&nbsp;&nbsp;<el-switch v-model="article.canComment"/>
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <label>Status</label>&nbsp;&nbsp;
          <el-radio-group v-model="article.status">
            <el-radio label="1">published</el-radio>
            <el-radio label="0">draft</el-radio>
            <el-radio label="2">deleted</el-radio>
          </el-radio-group>
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onClear">清空</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import MDinput from '@/components/MDinput'
  import MarkdownEditor from '@/components/MarkdownEditor'
  import {createArticle} from '@/api/article'

  const defaultForm = {
    title: '',
    slug: '',
    summary: '',
    content: '',
    status: '0',
    canComment: true,
    tags:[],
    createdAt:'',
  }
  export default {
    name: 'articleDetail',
    components: {MDinput, MarkdownEditor},
    props: {
      isEdit: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        article: Object.assign({}, defaultForm),
        options: [{
          value: '选项1',
          label: '黄金糕'
        }, {
          value: '选项2',
          label: '双皮奶',
          disabled: true
        }, {
          value: '选项3',
          label: '蚵仔煎'
        }, {
          value: '选项4',
          label: '龙须面'
        }, {
          value: '选项5',
          label: '北京烤鸭'
        }],
      }
    },
    created() {
      if (this.isEdit) {
        this.fetchData()
      }
    },
    methods: {
      onSubmit() {
        if (this.isEdit) {
          editArticle(this.article).then(response => {
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
          createArticle(this.article).then(response => {
            if (response.code === 20000) {
              this.$notify({
                title: '成功',
                message: '已创建',
                type: 'success',
                duration: 2000
              })
              Object.assign(this.article, defaultForm)
            }
          })
        }
      },
      onClear() {
        Object.assign(this.article, defaultForm)
      },
      fetchData() {
        fetchArticle(this.$route.query.id).then(response => {
          const article = response.data
          Object.assign(this.article, article)
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
