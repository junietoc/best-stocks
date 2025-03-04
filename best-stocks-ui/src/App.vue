<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <div class="bg-gray-50 min-h-screen flex flex-col">

    <header class="bg-blue-600 text-white">
      <div class="max-w-screen-xl mx-auto p-4 flex flex-col items-center md:flex-row md:justify-between">

        <div class="flex items-center mb-2 md:mb-0">
          <img
            alt="Vue logo"
            src="@/assets/logo.svg"
            width="60"
            height="60"
            class="mr-4"
          />
          <span class="text-xl font-bold">My App</span>
        </div>
      </div>
    </header>

    <main class="flex-grow py-8 px-4">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
</style>
