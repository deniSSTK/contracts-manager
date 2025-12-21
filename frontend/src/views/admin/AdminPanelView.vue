<template>
    <main class="page">
        <header>
            <h1>Admin Panel: <span v-if="route.name != RouteName.ADMIN_PANEL">{{route.params.entity}}</span> </h1>
            <router-link :to="{ name: RouteName.DASHBOARD }">
                <Button>To Dashboard</Button>
            </router-link>
        </header>
        <nav>
            <label>Pages:</label>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL }"
                v-if="route.name !== RouteName.ADMIN_PANEL"
            >
                <Button>Reset</Button>
            </router-link>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL_TABLE, params: { entity: 'contract' } }"
                v-if="$route.params.entity !== 'contract'"
            >
                <Button>Contracts</Button>
            </router-link>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL_TABLE, params: { entity: 'user' } }"
                v-if="$route.params.entity !== 'user'"
            >
                <Button>Users</Button>
            </router-link>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL_TABLE, params: { entity: 'person' } }"
                v-if="$route.params.entity !== 'person'"
            >
                <Button>Persons</Button>
            </router-link>
        </nav>
        <Return v-if="route.name != RouteName.ADMIN_PANEL"/>
        <router-view :key="route.name as string + `${route.params.entity ?? route.params.entity}`" />
    </main>
</template>

<script lang="ts" setup>
import Button from "@component/ui/button/Button.vue";
import {RouteName} from "@app/router/types";
import {useRoute} from "vue-router";
import Return from "@component/ui/button/return/Return.vue";

const route = useRoute()
</script>

<style>
.page {
    gap: var(--gap-md);

    nav {
        display: flex;
        gap: var(--gap-sm);

        align-items: center;
    }
}
</style>