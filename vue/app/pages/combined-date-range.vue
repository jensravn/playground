<script setup lang="ts">
interface DateRangeItem {
  id: string;
  label: string;
  range: [Date, Date];
}

const mainRange = ref<[Date, Date]>([
  new Date("2025-01-01"),
  new Date("2025-03-31"),
]);

const additionalRanges = ref<DateRangeItem[]>([
  {
    id: "1",
    label: "Q2 2025",
    range: [new Date("2025-04-01"), new Date("2025-06-30")],
  },
  {
    id: "2",
    label: "Q3 2025",
    range: [new Date("2025-07-01"), new Date("2025-09-30")],
  },
]);

const minDate = ref(new Date("2024-01-01"));
const maxDate = ref(new Date("2025-12-31"));
</script>

<template>
  <div class="container mx-auto p-8 max-w-4xl">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-2">Combined Date Range</h1>
    </div>

    <UCard>
      <CombinedDateRange
        v-model="mainRange"
        v-model:additional-ranges="additionalRanges"
        :min-date="minDate"
        :max-date="maxDate"
      />
    </UCard>

    <div class="mt-8">
      <h2 class="text-xl font-semibold mb-4">Current Values</h2>
      <div class="space-y-4">
        <UCard>
          <h3 class="text-sm font-semibold mb-2">Main Range</h3>
          <pre class="text-sm overflow-auto">{{
            JSON.stringify(mainRange, null, 2)
          }}</pre>
        </UCard>
        <UCard>
          <h3 class="text-sm font-semibold mb-2">Additional Ranges</h3>
          <pre class="text-sm overflow-auto">{{
            JSON.stringify(additionalRanges, null, 2)
          }}</pre>
        </UCard>
      </div>
    </div>
  </div>
</template>
