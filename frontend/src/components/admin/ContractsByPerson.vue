<template>
    <h1>Person Contracts</h1>

    <span v-if="!loading && contracts.length === 0">Person does not have an contract yet</span>

    <table v-else>
        <thead>
        <tr>
            <th>Code</th>
            <th>Title</th>
            <th>Description</th>
            <th>Start Date</th>
            <th>End Date</th>
            <th>Actions</th>
        </tr>
        </thead>

        <tbody>
        <tr v-for="contract in contracts" :key="contract.id">
            <td><span>{{contract.code}}</span></td>
            <td><span>{{contract.title}}</span></td>
            <td><span>{{contract.description ?? '-'}}</span></td>
            <td><span>{{contract.startDate ? toLocalTime(contract.startDate) : '-'}}</span></td>
            <td><span>{{contract.endDate ? toLocalTime(contract.endDate) : '-'}}</span></td>
            <td>
                <div class="actions">
                    <Button @click="openContractPreviewPopup(contract.id)">Preview</Button>
                    <router-link :to="{ name: RouteName.ADMIN_PANEL_INFO, params: { entity: 'contract', entityId: contract.id }}">
                        <Button>Info</Button>
                    </router-link>
                </div>
            </td>
        </tr>
        </tbody>
    </table>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import Contract from "@model/contract/entity";
import {useRoute} from "vue-router";
import contractUsecase from "@usecase/contract/usecase";
import router from "@app/router/routes";
import Button from "@component/ui/button/Button.vue";
import {RouteName} from "@app/router/types";
import {toLocalTime} from "@util/time";

const route = useRoute()
const id = route.params.entityId as string

const loading = ref<boolean>(true)
const contracts = ref<Contract[]>([])

async function loadContracts() {
    contracts.value = await contractUsecase.getContractsByPersonId(id)
    console.log(contracts.value)
    loading.value = false
}

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
    await loadContracts()
})
</script>