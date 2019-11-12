<template>
  <div class="app-container">
    <el-form ref="form" :model="article">
      <el-form-item style="margin-bottom: 40px;" prop="title">
        <MDinput name="name" v-model="article.title" required :maxlength="100">
          标题
        </MDinput>
      </el-form-item>

      <div class="editor-container">
        <MarkdownEditor ref="editor" v-model="article.content"></MarkdownEditor>
      </div>
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
    content: '',
    order: '',
    module_id: '',
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
        module: {},
        fileList: [],
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
