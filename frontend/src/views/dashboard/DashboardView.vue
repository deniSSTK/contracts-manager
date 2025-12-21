<template>
    <main class="page">
        <DashboardHeader />
        <span v-if="!loading && contracts.length === 0">You don't have any contract yet.</span>
        <div class="contract__container" v-else-if="!loading && contracts.length > 0">
            <div
                v-for="contract in contracts"
                class="contract"
                @click="openContractPreviewPopup(contract.id)"
            >
                <span>Title: {{contract.title}}</span>
                <span>Code: {{contract.code}}</span>
            </div>
        </div>
    </main>
</template>
<script setup lang="ts">
import DashboardHeader from "@component/dashboard/DashboardHeader.vue";
import {onMounted, ref} from "vue";
import Contract from "@model/contract/entity";
import authUsecase from "@usecase/auth/usecase";
import router from "@app/router/routes";
import {useRoute} from "vue-router";

import "./dashboard.scss"

const route = useRoute()

const contracts = ref<Contract[]>([])
const loading = ref<boolean>(true)

const openContractPreviewPopup = async (contractId: string) => {
    await router.push({
        name: route.name,
        params: route.params,
        query: {
            ...route.query,
            contractPreview: contractId
        }
    });
}

onMounted(async() => {
    contracts.value = await authUsecase.getUserContracts()
    loading.value = false
})
</script>