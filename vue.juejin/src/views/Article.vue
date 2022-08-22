<template>
    <div>
        <Navbar/>
        <atc/>
    </div>
</template>
<script>
import Navbar from "@/components/Navbar.vue";
import atc from "@/components/atc.vue";

export default {
  name: 'Article',
  data() {
    return {}
  },
  components: { Navbar, atc },
}
</script>
<style>
</style>