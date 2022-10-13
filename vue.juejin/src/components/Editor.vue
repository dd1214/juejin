<template>
<div>
    <el-button type="primary" plain class="mg8">提交文章</el-button>
    <div class="flex">

    <div style="border: 1px solid #ccc; width:50%">
        <Toolbar
            style="border-bottom: 1px solid #ccc"
            :editor="editor"
            :mode="mode"
            :defaultConfig="toolbarConfig"
        />
        <Editor
            style="height: 400px; overflow-y: hidden;"
            v-model="html"
            :defaultConfig="editorConfig"
            :mode="mode"
            @onCreated="onCreated"
        />
    </div>
    <div style="border: 1px solid #ccc; width:50%">实时预览区域</div>
        </div>
    </div>
</template>
  
  <script>
import Vue from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { DomEditor } from '@wangeditor/editor'
export default Vue.extend({
    components: { Editor, Toolbar },
    data() {
        return {
            editor: null,
            html: '<p>欢迎来到编辑页</p>',
            toolbarConfig:{toolbarKeys:[
                    'headerSelect',//格式区
                    "blockquote",//引用
                    '|',
                    "bold",//黑体
                    "underline",//下划线
                     'italic', //删除线
                    "through",
                    
                    '|',//添加区
                    "codeBlock",//代码块 代码高亮使用Prism.js实现
                    "emotion",//表情
                    "bulletedList",//无序列表
                    "numberedList",//有序列表
                    "insertLink",//插入链接
                    "insertTable",//插入表格
                    {
                    key: 'group-add', // 要以 group 开头
                    title: '...',
                    menuKeys: ["todo","sub","sup","enter","divider",] // 下级菜单的key //分割线
                    }, 
                    '|',//文字编辑区
                    "fontFamily",//默认字体
                    "fontSize",  
                    "color",
                    "bgColor",
                    '|',//图像区
                    "uploadImage",
                    "editImage",
                    '|',//功能区
                    "undo",
                    "redo",
                    "fullScreen",
                ]},
                
            editorConfig: { placeholder: '请输入您的文章' },
            mode: 'default', // or 'simple'
            toolbar:DomEditor.getToolbar(this.editor),
        }
    },
    methods: {
        onCreated(editor) {
            this.editor = Object.seal(editor) // 一定要用 Object.seal() ，否则会报错
        },
        
    },
    mounted() {
    },
    beforeDestroy() {
        const editor = this.editor
        if (editor == null) return
        editor.destroy() // 组件销毁时，及时销毁编辑器
    }
})
</script>

<style src="@wangeditor/editor/dist/css/style.css"></style>
