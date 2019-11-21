<template>
  <div class="app-container">
    <el-form ref="article" :model="article" :rules="rules">
      <el-form-item prop="title">
        <el-col :span="11">
          <MDinput v-model="article.title" name="name" required :maxlength="100">
            标题
          </MDinput>
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <MDinput v-model="article.slug" name="name" :maxlength="100">
            Slug
          </MDinput>
        </el-col>
      </el-form-item>

      <el-form-item prop="summary">
        <MDinput v-model="article.summary" name="name" :maxlength="255">
          Summary
        </MDinput>
      </el-form-item>

      <el-form-item prop="content" class="editor-container">
        <MarkdownEditor ref="editor" v-model="article.content" />
      </el-form-item>

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
            placeholder="请选择文章标签"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <label>Created At</label>&nbsp;&nbsp;
          <el-date-picker
            v-model="article.createdAt"
            type="datetime"
            placeholder="选择日期时间"
          />
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-col :span="11">
          <label>Can Comment</label>&nbsp;&nbsp;<el-switch v-model="article.canComment" />
        </el-col>
        <el-col :span="2">&nbsp;</el-col>
        <el-col :span="11">
          <label>Status</label>&nbsp;&nbsp;
          &nbsp;<el-switch v-model="article.status" />
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
import MarkdownEditor from '@/components/MarkdownEditor'
import { createArticle, editArticle, fetchArticle } from '@/api/article'

const defaultForm = {
  title: '',
  slug: '',
  summary: '',
  content: '',
  status: false,
  canComment: true,
  tags: [],
  createdAt: ''
}
export default {
  name: 'ArticleDetail',
  components: { MDinput, MarkdownEditor },
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
    return {
      article: Object.assign({}, defaultForm),
      options: [{
        value: '选项1',
        label: '选项1'
      }],
      rules: {
        title: [{ validator: validateRequire }],
        content: [{ validator: validateRequire }]
      }
    }
  },
  created() {
    if (this.isEdit) {
      this.fetchData()
    }
  },
  methods: {
    submitForm() {
      this.$refs.article.validate(valid => {
        if (valid) {
          this.onSubmit()
        } else {
          return false
        }
      })
    },
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
        article.createdAt = article.created_on
        article.status = article.status.toString()
        if (article.tags == null) {
          article.tags = []
        } else {
          var tags = []
          var i
          for (i in article.tags) {
            tags.push(article.tags[i].name)
          }
          article.tags = tags
        }
        Object.assign(this.article, article)
      })
    }
  }
}
</script>

