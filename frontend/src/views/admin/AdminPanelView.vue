<template>
    <main class="page">
        <header>
            <h1>Admin Panel{{buildName()}}</h1>
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
                :to="{ name: RouteName.ADMIN_PANEL_CONTRACTS }"
                v-if="$route.name != RouteName.ADMIN_PANEL_CONTRACTS"
            >
                <Button>Contracts</Button>
            </router-link>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL_USERS }"
                v-if="$route.name != RouteName.ADMIN_PANEL_USERS"
            >
                <Button>Users</Button>
            </router-link>
            <router-link
                :to="{ name: RouteName.ADMIN_PANEL_PERSONS }"
                v-if="$route.name != RouteName.ADMIN_PANEL_PERSONS"
            >
                <Button>Persons</Button>
            </router-link>
        </nav>
        <router-view :key="route.name" />
    </main>
</template>

<script lang="ts" setup>
import Button from "@component/ui/button/Button.vue";
import {RouteName} from "../../app/routes";
import {useRoute} from "vue-router";

const route = useRoute()

function buildName(): string {
    if (route.name !== RouteName.ADMIN_PANEL) {
        const list = route.name.split("-")

        return ": " + list[list.length - 1]
    }
}
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