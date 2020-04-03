//首先要进行挂载
import Vue from 'vue'//把vue这个类库引用进来
import App from './app.vue'
import VueResource from 'vue-resource';
// import './assets/styles/footer.styl'
import './assets/styles/global.styl'

const root = document.createElement('div');
document.body.appendChild(root);

new Vue({
  render:(h) => h(App)
}).$mount(root)
Vue.use(VueResource)
