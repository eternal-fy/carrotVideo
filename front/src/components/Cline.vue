<template>
  <div id="Cline">
    <p>{{title}}</p>
    <button @click="$emit('remove')">remove</button>
  </div>
</template>

<script>
export default {
  name: "Cline",
  props:['title']
}
</script>

<style scoped>

</style>