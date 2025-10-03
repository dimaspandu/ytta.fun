import { createRouter, createWebHistory } from 'vue-router';
import ProductList from '../pages/ProductList.vue';
import ProductDetail from '../pages/ProductDetail.vue';
import SessionExpired from '../pages/SessionExpired.vue';

const routes = [
  { path: '/', name: 'Home', component: ProductList },
  { path: '/product/:id', name: 'ProductDetail', component: ProductDetail },
  { path: '/expired', name: 'SessionExpired', component: SessionExpired }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
