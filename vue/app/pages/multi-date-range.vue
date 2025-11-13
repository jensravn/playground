<script setup lang="ts">
interface DateRangeItem {
  id: string;
  label: string;
  range: [Date, Date];
}

const dateRanges = ref<DateRangeItem[]>([
  {
    id: "1",
    label: "Q1 2025",
    range: [new Date("2025-01-01"), new Date("2025-03-31")],
  },
  {
    id: "2",
    label: "Q2 2025",
    range: [new Date("2025-04-01"), new Date("2025-06-30")],
  },
]);

const minDate = new Date("2024-01-01");
const maxDate = new Date("2025-12-31");
</script>

<template>
  <div class="container mx-auto p-8 max-w-4xl">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-2">Multi Date Range</h1>
    </div>

    <UCard>
      <MultiDateRange
        v-model="dateRanges"
        :min-date="minDate"
        :max-date="maxDate"
      />
    </UCard>

    <div class="mt-8">
      <h2 class="text-xl font-semibold mb-4">Current Values</h2>
      <UCard>
        <pre class="text-sm overflow-auto">{{
          JSON.stringify(dateRanges, null, 2)
        }}</pre>
      </UCard>
    </div>
  </div>
</template>
