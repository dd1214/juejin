<template>
    <div>
        <Navbar/>
        <LowerHalf/>
    </div>
</template>
<script>
import Navbar from "@/components/Navbar.vue";
import Adv from "@/components/adv.vue";
import LowerHalf from "@/components/LowerHalf.vue";

export default {
  name: 'Home',
  data() {
    return {}
  },
  components: { Navbar, Adv, LowerHalf },
}
</script>

