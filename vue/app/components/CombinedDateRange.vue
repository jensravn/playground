<script setup lang="ts">
interface DateRangeItem {
  id: string;
  label: string;
  range: [Date, Date];
}

const props = withDefaults(
  defineProps<{
    modelValue?: [Date, Date];
    additionalRanges?: DateRangeItem[];
    minDate?: Date;
    maxDate?: Date;
  }>(),
  {
    additionalRanges: () => [],
    minDate: () => new Date(Date.now() - 365 * 24 * 60 * 60 * 1000),
    maxDate: () => new Date(),
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: [Date, Date]];
  "update:additionalRanges": [value: DateRangeItem[]];
}>();

const mainRange = computed({
  get: () => props.modelValue,
  set: (v) => v && emit("update:modelValue", v),
});

const additionalRanges = computed({
  get: () => props.additionalRanges,
  set: (v) => emit("update:additionalRanges", v),
});

const isExpanded = ref(false);
</script>

<template>
  <div class="space-y-4">
    <DateRange v-model="mainRange" :min-date="minDate" :max-date="maxDate" />

    <div v-if="additionalRanges.length > 0" class="border-t pt-4">
      <button
        @click="isExpanded = !isExpanded"
        class="flex items-center gap-2 text-sm text-gray-600 hover:text-gray-900"
      >
        <UIcon
          :name="
            isExpanded
              ? 'i-heroicons-chevron-down'
              : 'i-heroicons-chevron-right'
          "
          class="w-4 h-4"
        />
        {{ additionalRanges.length }} additional range{{
          additionalRanges.length !== 1 ? "s" : ""
        }}
      </button>

      <div v-if="isExpanded" class="mt-4">
        <MultiDateRange
          v-model="additionalRanges"
          :min-date="minDate"
          :max-date="maxDate"
        />
      </div>
    </div>
  </div>
</template>
