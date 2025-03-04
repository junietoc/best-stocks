<script setup lang="ts">
import { defineProps } from "vue";

interface Stock {
  id: string;
  ticker: string;
  target_to: string;
}

const props = defineProps({
  stocks: {
    type: Array as () => Stock[],
    default: () => []
  }
});
</script>

<template>
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
</template>
