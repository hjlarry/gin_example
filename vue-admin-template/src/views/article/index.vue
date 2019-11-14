<template>
  <div class="app-container">
    <router-link to="/article/create">
      <el-button class="filter-item" type="primary" icon="el-icon-edit">添加新文章</el-button>
      <br><br>
    </router-link>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="Author" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Tags" width="400" align="center">
        <template slot-scope="scope" >
          <el-tag v-for="tag in scope.row.tags" v-bind:key="tag.id" style="margin-right: 3px;">{{tag.name}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusLabelFilter">{{ scope.row.status | statusFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="Created At" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.created_on }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="200" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
          <el-button size="mini" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  import {getList} from '@/api/article'

  export default {
    filters: {
      statusFilter(status) {
        const statusMap = {
          1: 'published',
          0: 'draft',
          2: 'deleted'
        }
        return statusMap[status]
      },
      statusLabelFilter(status) {
        const statusMap = {
          1: 'success',
          0: 'warning',
          2: 'danger'
        }
        return statusMap[status]
      }
    },
    data() {
      return {
        list: null,
        listLoading: true
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        this.listLoading = true
        getList().then(response => {
          this.list = response.data.lists
          this.listLoading = false
        })
      },
      handleUpdate(row) {
        this.$router.push({
          path: '/article/edit',
          query: {'id': row.id}
        })
      }
    }
  }
</script>
