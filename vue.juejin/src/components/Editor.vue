<template>
<div>
    <el-button type="primary" plain class="mg8" >提交</el-button>
    <div class="markdown">
        <MarkdownPro
        :height="500"
        />
    </div>
    <div class="markdown">
        <MarkdownPro
        :height="500"
        :isPreview="true"
        />
    </div>
    </div>
</template>
  
  <script>
import Vue from 'vue'
import { MarkdownPro } from 'vue-meditor'
export default Vue.extend({
    name:'Editor',
    components: {MarkdownPro},
    data() {
        return {
        }
    },
    methods: {
        
    },
    mounted() {
      
    },
    beforeDestroy() {
     
    }
})
</script>
