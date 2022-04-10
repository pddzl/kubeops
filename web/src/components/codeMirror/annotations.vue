<template>
  <div>
    <div>
      <textarea ref="codeEditor" />
    </div>
  </div>
</template>

<script>
import { onBeforeUnmount, onMounted, ref, toRefs } from 'vue'
import yaml from 'js-yaml'
import 'codemirror/mode/yaml/yaml.js'
// codemirror基础资源引入
import _CodeMirror from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/lib/codemirror.js'
import 'codemirror/mode/javascript/javascript.js'

// 代码错误检查
// import 'script-loader!jsonlint'
// import 'codemirror/addon/lint/lint.css'
// import 'codemirror/addon/lint/lint'
// import 'codemirror/addon/lint/json-lint'

import 'codemirror/addon/display/autorefresh'
import 'codemirror/addon/edit/matchbrackets'

// 折叠资源引入:开始
import 'codemirror/addon/fold/foldgutter.css'
import 'codemirror/addon/fold/foldcode.js'
import 'codemirror/addon/fold/brace-fold.js'
import 'codemirror/addon/fold/comment-fold.js'
import 'codemirror/addon/fold/indent-fold.js'
import 'codemirror/addon/fold/foldgutter.js'
// 折叠资源引入:结束

const CodeMirror = window.CodeMirror || _CodeMirror

export default {
  name: 'VueCodeMirrorAnnotations',
  props: {
    modelValue: {
      type: String,
      default: () => ''
    },
    defaultValue: {
      type: String,
      default: ''
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  setup(props, context) {
    const mimeList = ref(['YAML', 'JSON'])
    const mimeType = ref('YAML')
    const { modelValue, defaultValue, readOnly } = toRefs(props)
    const codeEditor = ref()
    let editor
    // watch(readOnly, () => {
    //   if (editor != null) {
    //     editor.setOption('readOnly', readOnly.value)
    //   }
    // })
    onMounted(() => {
      editor = CodeMirror.fromTextArea(codeEditor.value, {
        value: modelValue.value,
        mode: 'yaml',
        mime: 'text/x-yaml',
        indentWithTabs: false, // 在缩进时，是否需要把 n*tab宽度个空格替换成n个tab字符，默认为false
        smartIndent: true, // 自动缩进，设置是否根据上下文自动缩进（和上一行相同的缩进量）。默认为true
        lineNumbers: true, // 是否在编辑器左侧显示行号
        matchBrackets: true, // 括号匹配
        readOnly: readOnly.value,
        autoRefresh: true,
        // lint: true, // window.jsonlint not defined, CodeMirror JSON linting cannot run
        // 启用代码折叠相关功能:开始
        foldGutter: true,
        lineWrapping: true,
        gutters: ['CodeMirror-linenumbers', 'CodeMirror-foldgutter', 'CodeMirror-lint-markers'],
        // 启用代码折叠相关功能:结束
        // styleActiveLine: true // 光标行高亮
      })
      if (defaultValue.value) {
        editor.setValue(defaultValue.value)
      } else {
        const s1 = modelValue.value.replace(/\\"/g, '"').replace(/"\{/g, '{').replace(/\\n"/g, '')
        // editor.setValue(JSON.stringify(JSON.parse(s1), null, 2))
        editor.setValue(yaml.dump(JSON.parse(s1)))
      }
    })
    onBeforeUnmount(() => {
      if (editor !== null) {
        editor.toTextArea()
        editor = null
      }
    })
    return {
      codeEditor,
      mimeList,
      mimeType
    }
  }
}
</script>

