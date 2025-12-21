<template>
    <teleport to="body" v-if="show">
        <div class="popup__wrapper" @click="close">
            <div class="popup" @click.stop>
                <Button @click="close">Close</Button>
                <template v-if="contract">
                    <h1>Contract</h1>
                    <span>Title: {{contract.title}}</span>
                    <span>Code: {{contract.code}}</span>
                    <span>Created: {{toLocalTime(contract.createdAt)}}</span>
                    <span v-if="contract.startDate">Start Date: {{toLocalTime(contract.startDate)}}</span>
                    <span v-if="contract.endDate">End Date: {{toLocalTime(contract.endDate)}}</span>
                    <template v-if="contract.description">
                        <label>Description:</label>
                        <span>{{contract.description}}</span>
                    </template>
                </template>
                <template v-if="contract && persons.length > 0">
                    <h1>Persons</h1>
                    <div v-for="(person, n) in persons" :key="person.id" class="person">
                        <span class="primary">{{n+1}}</span>
                        <span>Code: {{person.code}}</span>
                        <span>Name: {{person.name}}</span>
                        <span>Type: {{person.type}}</span>
                        <span v-if="person.email">Email: {{person.email}}</span>
                        <span v-if="person.phone">Phone: {{person.phone}}</span>
                    </div>
                </template>
            </div>
        </div>
    </teleport>
</template>

<script lang="ts" setup>
import {ref, watch} from "vue";
import {useRoute} from "vue-router";

import "./contract-preview-popup.scss"
import router from "@app/router/routes";
import Button from "@component/ui/button/Button.vue";
import contractUsecase from "@usecase/contract/usecase";
import Contract from "@model/contract/entity";
import {toLocalTime} from "@util/time";
import Person from "@model/person/model";

const route = useRoute()

const show = ref<boolean>(false)

const contract = ref<Contract | null>(null)
const persons = ref<Person[]>([])

async function close() {
    const newQuery = { ...route.query }
    delete newQuery.contractPreview
    await router.replace({ query: newQuery })
}

const showContract = async () => {
    const contractId = route.query.contractPreview as string

    document.body.style.overflow = 'hidden'
    contract.value = await contractUsecase.get(contractId)
    persons.value = await contractUsecase.getPersonsByContractId(contractId)
}

watch(
    () => route.query,
    query => {
         show.value = !!query.contractPreview;
    }
);

watch(show,async (value) => {
    if (value) {
        await showContract()
    } else {
        document.body.style.overflow = ''
    }
})
</script>