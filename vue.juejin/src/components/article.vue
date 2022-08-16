<template>
        <div class="article white shadow mgb1">{{artTitle}}</div>
</template>

<script>
export default {
    name:'article',
    props: ['artTitle'], 
    data(){
        return {artTitle:'文章'}
    }
}
</script>