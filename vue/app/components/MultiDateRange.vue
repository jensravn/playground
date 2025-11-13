<script setup lang="ts">
interface DateRangeItem {
  id: string;
  range: [Date, Date];
}

const props = withDefaults(
  defineProps<{
    modelValue?: DateRangeItem[];
    minDate?: Date;
    maxDate?: Date;
  }>(),
  {
    modelValue: () => [],
    minDate: () => new Date(Date.now() - 365 * 24 * 60 * 60 * 1000),
    maxDate: () => new Date(),
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: DateRangeItem[]];
}>();

const ranges = computed({
  get: () => props.modelValue,
  set: (v) => emit("update:modelValue", v),
});

const updateRange = (id: string, newRange: [Date, Date]) => {
  ranges.value = ranges.value.map((item) =>
    item.id === id ? { ...item, range: newRange } : item
  );
};
</script>

<template>
  <div v-for="item in ranges" :key="item.id" class="space-y-2">
    <div class="flex justify-between items-center"></div>
    <DateRange
      :model-value="item.range"
      @update:model-value="updateRange(item.id, $event)"
      :min-date="minDate"
      :max-date="maxDate"
    />
  </div>
</template>
