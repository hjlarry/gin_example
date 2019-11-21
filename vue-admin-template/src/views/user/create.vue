<template>
  <user-detail :is-edit="false" />
</template>

<script>
import UserDetail from './UserDetail'
export default {
  name: 'CreateForm',
  components: { UserDetail }
}
</script>
