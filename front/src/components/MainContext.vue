<template>
  <div id="mainContext">
    <router-view/>
<!--    <ul>
      <li v-for="(num,index) in 10" :key="index">{{ num }}</li>
    </ul>-->
    <div>
      <form @submit.prevent="add">
        <input v-model="text">
        <input type="submit" value="add">
      </form>
    </div>
    <div>
      <ul>
        <Cline v-for="(item, index) in items" :key="index" :title="item" @remove="items.splice(index,1)"></Cline>
      </ul>
    </div>
    <button @keyup.enter.exact="alert('hehe')">keyboardClick</button>
    <button @mousedown.right="alert('hehe')">MouseClick</button>
  </div>


</template>

<script>
import Cline from "@/components/Cline";

export default {
  name: "MainContext",
  data() {
    return {
      text: '',
      items: ["one", "two", "three"],
      index: 3
    }
  },
  components: {Cline},
  methods: {
    add() {
      this.items.push(this.text)
    },
    alert(mes){
      alert(mes)
    }
  },

}
</script>

<style scoped>
#mainContext {
  padding-left: 200px;
  padding-top: 100px;
}
</style>