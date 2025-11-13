<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    modelValue?: [Date, Date];
    minDate?: Date;
    maxDate?: Date;
  }>(),
  {
    minDate: () => new Date(Date.now() - 365 * 24 * 60 * 60 * 1000),
    maxDate: () => new Date(),
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: [Date, Date]];
}>();

const value = computed({
  get: () => props.modelValue?.map((d) => d.getTime()) as [number, number],
  set: (v) =>
    emit("update:modelValue", v.map((t) => new Date(t)) as [Date, Date]),
});

const format = (timestamp: number) =>
  new Date(timestamp).toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });

const toDateInputFormat = (timestamp: number) => {
  const date = new Date(timestamp);
  return date.toISOString().split("T")[0];
};

const startDate = computed({
  get: () => toDateInputFormat(value.value[0]),
  set: (v: string) => {
    const timestamp = new Date(v).getTime();
    value.value = [timestamp, value.value[1]];
  },
});

const endDate = computed({
  get: () => toDateInputFormat(value.value[1]),
  set: (v: string) => {
    const timestamp = new Date(v).getTime();
    value.value = [value.value[0], timestamp];
  },
});
</script>

<template>
  <div class="flex gap-4 items-center">
    <UInput type="date" v-model="startDate" />
    <USlider
      v-model="value"
      :min="minDate.getTime()"
      :max="maxDate.getTime()"
      :step="86400000"
      tooltip
      class="flex-1"
    />
    <UInput type="date" v-model="endDate" />
  </div>
</template>
