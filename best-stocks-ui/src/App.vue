<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <!-- Wrap the entire app in a background and min-h-screen if you like -->
  <div class="bg-gray-50 min-h-screen">
    <header class="bg-blue-600 text-white p-4 flex flex-col items-center md:flex-row md:justify-between">
      <img
        alt="Vue logo"
        class="logo"
        src="@/assets/logo.svg"
        width="125"
        height="125"
      />

      <nav class="space-x-4 mt-2 md:mt-0">
        <RouterLink
          to="/"
          class="hover:text-blue-200 transition-colors"
        >
          Home
        </RouterLink>
        <RouterLink
          to="/about"
          class="hover:text-blue-200 transition-colors"
        >
          About
        </RouterLink>
      </nav>
    </header>

    <!-- Where the route-based component gets injected -->
    <main class="py-8">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.logo {
  /* If you want to keep your custom style for the logo, do it here */
  display: block;
  margin: 0 auto;
}
</style>
