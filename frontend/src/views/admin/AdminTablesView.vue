<template>
    <router-link :to="{ name: RouteName.ADMIN_PANEL_NEW }" v-if="infoEntities[route.params.entity as string].canCreate">
        <Button >Create new {{route.params.entity}}</Button>
    </router-link>

    <AdminPanelTable />
</template>

<script lang="ts" setup>
import {useRoute} from "vue-router";
import {RouteName} from "@app/router/types";

import Button from "@component/ui/button/Button.vue";
import AdminPanelTable from "@component/admin/panelTable/AdminPanelTable.vue";
import infoEntities from "@entity/info";

const route = useRoute()
</script>

<style>
.button {
    width: max-content;
}
</style>