<template>
<div>
  <router-view/>
</div>
</template>

<script>
import Navbar from "./components/Navbar.vue";

export default {
  name: 'app',
  data() {
    return {}
  },
  components: { Navbar,},

}
</script>

<style>
</style>
