<script setup lang="ts">
const components = [
  {
    title: "Date Range",
    description: "Interactive date range selector using dual-thumb slider",
    path: "/date-range",
    icon: "📅",
  },
  {
    title: "Multi Date Range",
    description: "Manage multiple date ranges with individual sliders",
    path: "/multi-date-range",
    icon: "📊",
  },
  {
    title: "Combined Date Range",
    description: "Main date range with collapsible additional ranges",
    path: "/combined-date-range",
    icon: "🗂️",
  },
];
</script>

<template>
  <div class="min-h-screen p-8">
    <div class="max-w-4xl mx-auto">
      <h1 class="text-5xl font-bold mb-4">Vue Playground</h1>

      <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        <NuxtLink
          v-for="component in components"
          :key="component.path"
          :to="component.path"
          class="group block p-6 border border-gray-200 dark:border-gray-800 rounded-lg bg-white dark:bg-gray-900 hover:border-primary-500 dark:hover:border-primary-400 transition-colors"
        >
          <div class="text-4xl mb-4">{{ component.icon }}</div>
          <h2
            class="text-2xl font-semibold mb-2 group-hover:text-primary-500 dark:group-hover:text-primary-400 transition-colors"
          >
            {{ component.title }}
          </h2>
          <p class="text-gray-600 dark:text-gray-400">
            {{ component.description }}
          </p>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
