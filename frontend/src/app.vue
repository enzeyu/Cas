<template>
<!-- vue组件必须有一个独立的外部节点包围。也就是说在template不能有两个并列的节点 -->
<div id="app">
  <div id="cover"></div>
  <Header></Header>  <!--显示的是标题 JSTodo-->
  <Footer></Footer>  <!--显示的是作者做法 Written by AUTHER-->
  <Tod></Tod>    <!--显示的是接下来要做什么这个组件 -->


</div>
</template>

<script>
import Header from './todo/header.vue'
import Footer from './todo/footer.jsx'
import Tod from './todo/todo.vue'


export default {
  components:{
    Header,
    Footer,
    Tod,
  },
  
}
</script>
<style lang="stylus" scoped>
//加了scoped使得样式只应用在这层组件，使得不会跟别的组件中相同类属性的样式冲突
#app{  /*分别代表app和背景的相对偏移位置*/
  position absolute
  left 0
  right 0
  top 0             
  bottom 0
}
#cover{
  position absolute
  left 0
  top 0
  right 0
  bottom 0
  background-color rgba(0,0,0,0.3)  /*调整背景色*/
  opacity .1  /*透明度*/
  z-index -1  /*设置元素堆叠顺序*/
}
</style>