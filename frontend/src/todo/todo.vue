<template>
  <section class="real-app">
    <input
      type="text"
      class="add-input"
      autofocus="autofocus"
      placeholder="输入待检测URL："  
      @keyup.enter="sendUrl"
    /> 
    <table id="table-3">
        <tr>
            <th>URL</th>
            <th>状态码</th>
            <th>响应时间(s)</th>
			<th>域名</th>
			<th>暗链检测</th>
        </tr>
        <tbody>
        <tr v-for="site in todos">
			<td><a :href="site.request_url" target="_blank">{{site.request_url}}</a></td>
			<td v-cloak>{{site.resp_code}}</td>
			<td v-cloak>{{site.resp_time/1000000000}}</td>
			<td v-cloak>{{site.domain}}</td>
			<td v-cloak><a href="#" class="btn-gradient" οnclick="checkDarklinks()">check me</a></td>
        </tr>
        </tbody>
    </table>


  </section>
</template>

<script>

var backend_api = 'http://127.0.0.1:8081/geturls'

export default {
  data() {
    return {
      todos: [],
    };
  },


  methods: {
    sendUrl(e) {
		this.todos.unshift({   //向数组的开头添加一个或更多元素
        content: e.target.value.trim(),  //event.target.value获取当前文本框的值
      })
      var send_url = JSON.stringify(e.target.value.trim());
	  e.target.value = '';   //加完将值清零
	  this.$http.post(backend_api,{url:send_url},{emulateJSON:true}).then(response => {
		  this.todos = response.body.message;
	  })
	  
    },
	checkDarklinks(e){
		console
	}
  }
};
</script>

<style lang="stylus" scoped>
.real-app {
  width: 1000px;
  margin: 0 auto;
  box-shadow: 0 0 5px #666;
}

.add-input {
  position: relative;
  margin: 0;
  width: 100%
  font-size: 24px;
  font-family: inherit;
  font-weight: inherit;
  line-height: 1.4em;
  border: none;
  outline: none;
  color: inherit;
  box-sizing: border-box;
  font-smoothing: antialiased;
  padding: 16px 16px 16px 36px;
  border: none;
  box-shadow: inset 0 -2px 1px rgba(0, 0, 0, 0.03);
}
#table-3 thead, #table-3 tr {
border-top-width: 1px;
border-top-style: solid;
border-top-color: rgb(235, 242, 224);
}
#table-3 {
width: 1000px;
border-bottom-width: 1px;
border-bottom-style: solid;
border-bottom-color: rgb(235, 242, 224);
}

/* Padding and font style */
#table-3 td, #table-3 th {
padding: 5px 10px;
width: 200px;
font-size: 20px;
font-family: Verdana;
color: rgb(149, 170, 109);
}

/* Alternating background colors */
#table-3 tr:nth-child(even) {
background: rgb(230, 238, 214)
}
#table-3 tr:nth-child(odd) {
background: #FFF
}


.btn-gradient {
  text-decoration: none;
  color: white;
  width: 100px;
  padding: 10px 30px;
  display: inline-block;
  position: relative;
  border: 1px solid rgba(0,0,0,0.21);
  border-bottom: 4px solid rgba(0,0,0,0.21);
  border-radius: 4px;
  text-shadow: 0 1px 0 rgba(0,0,0,0.15);
  background: rgba(250,90,90,1);
}

[v-cloak]{
  display:none;
}
</style>