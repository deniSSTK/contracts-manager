<template>
    <header>
        <h1>Your contracts</h1>
        <router-link :to="{ name: RouteName.ADMIN_PANEL }">
            <Button v-if="authStore.isAdmin">To Admin Panel</Button>
        </router-link>
    </header>
</template>

<script lang="ts" setup>
import Button from "@component/ui/button/Button.vue";
import useAuthStore from "@store/auth/store";
import {RouteName} from "@app/router/routes";

const authStore = useAuthStore()
</script>