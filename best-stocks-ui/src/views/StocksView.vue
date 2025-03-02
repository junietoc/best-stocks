<script setup lang="ts">
import { ref, onMounted } from "vue";
import { fetchStocks } from "@/api/stocks";

const stocks = ref([]);

onMounted(async () => {
    stocks.value = await fetchStocks();
});
</script>

<template>
  <div class="p-4">
    <h1 class="text-xl font-bold">Stock List</h1>
    <ul>
      <li v-for="stock in stocks" :key="stock.id">
        {{ stock.name }} - ${{ stock.price }}
      </li>
    </ul>
  </div>
</template>