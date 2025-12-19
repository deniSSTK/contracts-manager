<template>
    <button
        v-bind="$attrs"
        class="button"
    >
        <slot />
    </button>
</template>

<script lang="ts" setup>
import "./style.css"

defineOptions({
    inheritAttrs: false
})
</script>