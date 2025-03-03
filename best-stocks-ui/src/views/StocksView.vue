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
  <div class="max-w-3xl mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4 text-center">Stock List</h1>

    <ul class="bg-white shadow-md rounded-lg divide-y divide-gray-200">
      <li
        v-for="stock in stocks"
        :key="stock.id"
        class="p-4 hover:bg-gray-50 flex items-center justify-between"
      >
        <span class="font-medium text-gray-700">
          {{ stock.ticker }}
        </span>
        <span class="text-green-600 font-semibold">
          {{ stock.target_to }}
        </span>
      </li>
    </ul>
  </div>
</template>

