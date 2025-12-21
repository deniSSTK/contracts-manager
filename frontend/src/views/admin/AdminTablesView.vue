<template>
    <nav>
        <router-link :to="{ name: RouteName.ADMIN_PANEL_NEW }" v-if="infoEntities[route.params.entity as string].canCreate">
            <Button >Create new {{route.params.entity}}</Button>
        </router-link>
        <template v-if="entity">
            Export
            <Button @click="entity.export('csv')">
                .csv
            </Button>
            <Button @click="entity.export('json')">
                .json
            </Button>
            Import
            <Button @click="triggerImport">Upload .csv | .json</Button>
            <input
                ref="fileInput"
                type="file"
                accept=".csv,.json"
                @change="onFile"
                style="display: none"
            />
        </template>
    </nav>

    <AdminPanelTable />
</template>

<script lang="ts" setup>
import {useRoute} from "vue-router";
import {RouteName} from "@app/router/types";

import Button from "@component/ui/button/Button.vue";
import AdminPanelTable from "@component/admin/panelTable/AdminPanelTable.vue";
import infoEntities from "@entity/info";
import {computed, ref} from "vue";
import fileEntities from "@entity/file";

const fileInput = ref<HTMLInputElement | null>(null);

const route = useRoute()

const entity = computed(() => fileEntities[route.params.entity as string])

async function onFile(e: Event) {
    const file = (e.target as HTMLInputElement).files?.[0];
    if (!file) return;

    await entity.value.import(file);

    window.location.reload();
}

function triggerImport() {
    fileInput.value?.click();
}

</script>

<style>
.button {
    width: max-content;
}
</style>