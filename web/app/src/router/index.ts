import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import Home from '../views/Home.vue';
import BookList from '../views/BookList.vue';
import Checkouts from '../views/Checkouts.vue';
import Book from '../views/Book.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/checkouts',
    name: 'Checkouts',
    component: Checkouts,
  },
  {
    path: '/books',
    name: 'BookList',
    component: BookList,
  },
  {
    path: '/books/:id',
    name: 'Book',
    component: Book,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
