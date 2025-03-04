<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { fetchStocks } from "@/api/stocks";

import StockFilters from "@/components/stocks/StockFilters.vue";
import StockPagination from "@/components/stocks/StockPagination.vue";

const stocks = ref<any[]>([]);
const filterText = ref("");
const currentPage = ref(1);
const pageSize = 12;

onMounted(async () => {
  try {
    stocks.value = await fetchStocks();
  } catch (error) {
    console.error("Failed to fetch stocks:", error);
  }
});

const filteredStocks = computed(() => {
  if (!filterText.value) {
    return stocks.value;
  }
  return stocks.value.filter((stock) =>
    stock.ticker.toLowerCase().includes(filterText.value.toLowerCase())
  );
});

const totalPages = computed(() => {
  return Math.ceil(filteredStocks.value.length / pageSize);
});

const paginatedStocks = computed(() => {
  const startIndex = (currentPage.value - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  return filteredStocks.value.slice(startIndex, endIndex);
});

function goToPreviousPage() {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
}

function goToNextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
}
</script>

<template>
  
  <main class="w-full min-h-screen p-4 bg-gray-50">
    <h1 class="text-2xl font-bold mb-4 text-center">Stocks</h1>

    
    <StockFilters
      :filterText="filterText"
      @update:filterText="(val) => (filterText.value = val)"
    />

    <div class="overflow-x-auto mt-4">
      <table class="w-full border-collapse text-left bg-white shadow-sm">
        <thead class="bg-gray-100 border-b border-gray-200">
          <tr>
            <th class="p-3 font-semibold text-gray-700">Ticker</th>
            <th class="p-3 font-semibold text-gray-700">Target From</th>
            <th class="p-3 font-semibold text-gray-700">Target To</th>
            <th class="p-3 font-semibold text-gray-700">Company</th>
            <th class="p-3 font-semibold text-gray-700">Action</th>
            <th class="p-3 font-semibold text-gray-700">Brokerage</th>
            <th class="p-3 font-semibold text-gray-700">Rating From</th>
            <th class="p-3 font-semibold text-gray-700">Rating To</th>
            <th class="p-3 font-semibold text-gray-700">Time</th>
          </tr>
        </thead>

        <tbody>
          <tr
            v-for="stock in paginatedStocks"
            :key="stock.id"
            class="border-b border-gray-100 hover:bg-gray-50"
          >
            <td class="p-3">{{ stock.ticker }}</td>
            <td class="p-3">{{ stock.target_from }}</td>
            <td class="p-3">{{ stock.target_to }}</td>
            <td class="p-3">{{ stock.company }}</td>
            <td class="p-3">{{ stock.action }}</td>
            <td class="p-3">{{ stock.brokerage }}</td>
            <td class="p-3">{{ stock.rating_from }}</td>
            <td class="p-3">{{ stock.rating_to }}</td>
            <td class="p-3">{{ stock.time }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    
    <div class="mt-4">
      <StockPagination
        :current-page="currentPage"
        :total-pages="totalPages"
        @prev="goToPreviousPage"
        @next="goToNextPage"
      />
    </div>
  </main>
</template>
