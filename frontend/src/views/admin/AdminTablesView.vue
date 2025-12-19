<template>
    <router-link :to="{ query: {} }" v-if="isNew">
        <Button>Cancel</Button>
    </router-link>

    <router-link :to="{ query: { new: 'true' } }" v-else>
        <Button>Create new {{name}}</Button>
    </router-link>

    <AdminCreateNewEntity v-if="isNew"/>
    <AdminPanelTable v-else/>
</template>

<script lang="ts" setup>
import {ref, watch} from "vue"
import {useRoute} from "vue-router";
import AdminPanelTable from "@component/admin/panelTable/AdminPanelTable.vue";
import AdminCreateNewEntity from "@view/admin/AdminCreateNewEntity.vue";
import Button from "@component/ui/button/Button.vue";

const route = useRoute()

const name = route.params.entity
const isNew = ref<boolean>(!!route.query.new)

watch(
    () => route.query.new,
    (newVal) => {
        isNew.value = !!newVal
    }
)
</script>

<style>
.button {
    width: max-content;
}
</style>