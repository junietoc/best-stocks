<script setup lang="ts">
import { ref, onMounted } from "vue";
import { fetchStocks } from "@/api/stocks";

const stocks = ref([]);

onMounted(async () => {
  try {
    stocks.value = await fetchStocks();
  } catch (error) {
    console.error("Failed to fetch stocks:", error);
  }
});
</script>

<template>
  <div class="p-4">
    <h1 class="text-xl font-bold">Stock List</h1>
    <ul>
      <li v-for="stock in stocks" :key="stock.id">
        {{ stock.ticker }} - ${{ stock.target_to }}
      </li>
    </ul>
  </div>
</template>
