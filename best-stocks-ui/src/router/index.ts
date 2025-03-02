import { createRouter, createWebHistory } from "vue-router";
import StocksView from "@/views/StocksView.vue";

const routes = [
    { path: "/", component: StocksView },
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});