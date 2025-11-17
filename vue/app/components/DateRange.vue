<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    modelValue?: [Date?, Date?];
    minDate?: Date;
    maxDate?: Date;
  }>(),
  {
    minDate: () => new Date(Date.now() - 1 * 24 * 60 * 60 * 1000),
    maxDate: () => new Date(Date.now() + 1 * 24 * 60 * 60 * 1000),
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: [Date?, Date?]];
}>();

const value = computed({
  get: () => props.modelValue?.map((d) => d?.getTime()) as [number?, number?],
  set: (v) =>
    emit(
      "update:modelValue",
      v.map((t) => (t ? new Date(t) : undefined)) as [Date?, Date?]
    ),
});

const toDateInputFormat = (timestamp?: number) => {
  if (timestamp === undefined) return "";
  const date = new Date(timestamp);
  return date.toISOString().split("T")[0];
};

const startDate = computed({
  get: () => toDateInputFormat(value.value[0]),
  set: (v: string) => {
    const timestamp = v ? new Date(v).getTime() : undefined;
    value.value = [timestamp, value.value[1]];
  },
});

const endDate = computed({
  get: () => toDateInputFormat(value.value[1] ? value.value[1] : undefined),
  set: (v: string) => {
    const timestamp = v ? new Date(v).getTime() : undefined;
    value.value = [value.value[0], timestamp];
  },
});
</script>

<template>
  <div class="flex gap-4 items-center">
    <UInput type="date" v-model="startDate" />
    <USlider
      v-if="value[0] === undefined && value[1] !== undefined"
      :model-value="value[1]"
      @update:model-value="(v) => value = [undefined, v as number]"
      :min="minDate.getTime()"
      :max="maxDate.getTime()"
      :step="86400000"
      tooltip
      class="flex-1"
    />
    <USlider
      v-else-if="value[0] !== undefined && value[1] === undefined"
      :model-value="value[0]"
      @update:model-value="(v) => value = [v as number, undefined]"
      :min="minDate.getTime()"
      :max="maxDate.getTime()"
      :step="86400000"
      inverted
      tooltip
      class="flex-1"
    />
    <USlider
      v-else-if="value[0] !== undefined && value[1] !== undefined"
      :model-value="value as [number, number]"
      @update:model-value="(v) => value = v as [number, number]"
      :min="minDate.getTime()"
      :max="maxDate.getTime()"
      :step="86400000"
      tooltip
      class="flex-1"
    />
    <div v-else class="flex-1">
      <USlider></USlider>
    </div>
    <UInput type="date" v-model="endDate" />
  </div>
</template>
