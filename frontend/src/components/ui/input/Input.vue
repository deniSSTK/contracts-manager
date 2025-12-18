<template>
    <input
        v-bind="$attrs"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        class="input"
    />
</template>

<script setup lang="ts">
import "./style.css";

defineProps<{
    modelValue: string
}>();

defineEmits<{
    (e: 'update:modelValue', value: string): void
}>();

defineOptions({
    inheritAttrs: false
})
</script>
