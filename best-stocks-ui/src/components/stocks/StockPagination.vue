<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  currentPage: {
    type: Number,
    default: 1
  },
  totalPages: {
    type: Number,
    default: 1
  }
});

const emits = defineEmits(["prev", "next"]);

function onPrevClick() {
  emits("prev");
}

function onNextClick() {
  emits("next");
}
</script>

<template>
  <div class="flex justify-between items-center mt-4">
    <button
      class="px-4 py-2 bg-gray-300 text-gray-700 rounded disabled:opacity-50"
      :disabled="currentPage === 1"
      @click="onPrevClick"
    >
      Previous
    </button>

    <div>
      Page {{ currentPage }} of {{ totalPages }}
    </div>

    <button
      class="px-4 py-2 bg-gray-300 text-gray-700 rounded disabled:opacity-50"
      :disabled="currentPage === totalPages"
      @click="onNextClick"
    >
      Next
    </button>
  </div>
</template>
