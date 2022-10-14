<template>
  <div class="preview">
    <div v-html="htmlData"></div>
  </div>
  
</template>

<script>
export default {
    name:'Preview',
    props:{
      htmlData:{
        type: String,
        default:''
      }},
    mounted(){

    },
    
  
}
</script>

<style scoped>
.preview{
  border: 1px solid #ccc;
  width:50%;
  max-height:600px;
  overflow:auto;
}
</style>