<template>
  <div class="app-container">
    <router-link to="/user/create">
      <el-button class="filter-item" type="primary" icon="el-icon-edit">添加新用户</el-button>
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
      <el-table-column label="Username">
        <template slot-scope="scope">
          {{ scope.row.username }}
        </template>
      </el-table-column>
      <el-table-column label="Email">
        <template slot-scope="scope">
          {{ scope.row.email }}
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="Actived" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.active | activeLabelFilter">{{ scope.row.active | activeFilter}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="Created At" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.created_on }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="modified_at" label="Modified At" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.modified_on }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="200" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange"
                     :current-page="listQuery.page" :page-sizes="[5,10,20,30,50]" :page-size="listQuery.limit"
                     layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
  import {fetchUsers, deleteUser} from '@/api/user'

  export default {
    filters: {
      activeLabelFilter(active) {
        if (active) return "success"
        return "info"
      },
      activeFilter(active) {
        if (active) return "Yes"
        return "No"
      }
    },
    data() {
      return {
        list: null,
        total: null,
        listLoading: true,
        listQuery: {
          page: 1,
          limit: 10,
        }
      }
    },
    created() {
      this.getList()
    },
    methods: {
      getList() {
        this.listLoading = true
        fetchUsers(this.listQuery).then(response => {
          this.list = response.data.lists
          this.total = response.data.total
          this.listLoading = false
        })
      },
      handleUpdate(row) {
        this.$router.push({
          path: '/user/edit',
          query: {'id': row.id}
        })
      },
      handleDelete(row) {
        const index = this.list.indexOf(row)
        const confirm = this.$confirm(`确定删除用户` + row.username + '?')
        confirm.then(() => {
          deleteUser(row.id).then(() => {
            this.$notify({
              title: '成功',
              message: '删除成功',
              type: 'success',
              duration: 2000
            })
            this.list.splice(index, 1)
          })
        }).catch(() => {
        })
      },
      handleSizeChange(val) {
        this.listQuery.limit = val
        this.getList()
      },
      handleCurrentChange(val) {
        this.listQuery.page = val
        this.getList()
      },
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .pagination-container {
    float: right;
    margin-top: 20px;
  }
</style>
