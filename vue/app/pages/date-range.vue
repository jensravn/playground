<script setup lang="ts">
const customMinDate = new Date("2024-01-01");
const customMaxDate = new Date("2024-12-31");

const customDateRange = ref<[Date, Date]>([
  new Date("2024-06-01"),
  new Date("2024-09-01"),
]);

const startDateOnly = ref<[Date?, Date?]>([new Date("2024-06-01"), undefined]);

const endDateOnly = ref<[Date?, Date?]>([undefined, new Date("2024-09-01")]);

const noDateSet = ref<[Date?, Date?]>([undefined, undefined]);
</script>

<template>
  <div class="container mx-auto p-8 max-w-4xl">
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-2">Date Range</h1>
    </div>

    <div class="space-y-8">
      <div>
        <h2 class="text-xl font-semibold mb-4">Start and End Dates</h2>
        <DateRange
          v-model="customDateRange"
          :min-date="customMinDate"
          :max-date="customMaxDate"
        />
      </div>

      <div>
        <h2 class="text-xl font-semibold mb-4">Only Start Date</h2>
        <DateRange
          v-model="startDateOnly"
          :min-date="customMinDate"
          :max-date="customMaxDate"
        />
      </div>

      <div>
        <h2 class="text-xl font-semibold mb-4">Only End Date</h2>
        <DateRange
          v-model="endDateOnly"
          :min-date="customMinDate"
          :max-date="customMaxDate"
        />
      </div>

      <div>
        <h2 class="text-xl font-semibold mb-4">No Dates Set</h2>
        <DateRange
          v-model="noDateSet"
          :min-date="customMinDate"
          :max-date="customMaxDate"
        />
      </div>
    </div>
  </div>
</template>
