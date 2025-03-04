<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

const props = defineProps({
  filterText: {
    type: String,
    default: ""
  }
});

const emits = defineEmits(["update:filterText"]);

function onInput(e: Event) {
  const value = (e.target as HTMLInputElement).value;
  emits("update:filterText", value);
}
</script>

<template>
  <div class="mb-4">
    <input
      :value="filterText"
      @input="onInput"
      type="text"
      placeholder="Filter by ticker"
      class="border p-2 w-full"
    />
  </div>
</template>
