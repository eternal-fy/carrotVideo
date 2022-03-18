<template>
  <div>
    <table>
      <tr v-for="(line,index) in arr" :key="index">
        <td v-for="(num,index) in line" :key="index">
          {{ num }}
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: "Table",
  props:{
      arr: Number,
    }
}
</script>

<style scoped>

</style>