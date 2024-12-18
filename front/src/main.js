import './assets/main.css';
import {createApp} from 'vue';
import {createPinia} from 'pinia';
import App from './App.vue';
import router from './router';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import axios from "axios";
import highlightDirective from './directives/highlight';

axios.defaults.baseURL = "http://localhost:8080";

const app = createApp(App);
app.use(createPinia());
app.use(router);
app.use(ElementPlus);
app.directive('highlight', highlightDirective);

app.mount('#app');
