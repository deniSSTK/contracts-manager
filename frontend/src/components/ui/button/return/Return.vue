<template>
    <Button @click="router.back()" class="return">
        <Icon icon="icon-park-outline:return" />
    </Button>
</template>
<script setup lang="ts">
import router from "@app/router/routes";
import Button from "@component/ui/button/Button.vue";
import {Icon} from "@iconify/vue";
</script>

<style scoped>
.return {
    width: max-content;
    padding: var(--padding-md);
}
</style>