<template>
    <h1>Contract Persons</h1>

    <form class="person-add" @submit.prevent="addPerson">
        <span>add person</span>
        <Input v-model="addDTOForm.personId" placeholder="Input personId" required/>
        <select v-model="addDTOForm.role" class="role-select" required>
            <option :value="null" disabled hidden>Select role</option>
            <option
                v-for="val in Object.values(ContractRole)"
                :key="val"
                :value="val"
            >
                {{val}}
            </option>
        </select>
        <Button type="submit">Add</Button>
    </form>

    <table>
        <thead>
            <tr>
                <th>Code</th>
                <th>Name</th>
                <th>Type</th>
                <th>Actions</th>
            </tr>
        </thead>

        <tbody>
            <tr v-for="person in persons" :key="person.id">
                <td><span>{{person.code}}</span></td>
                <td><span>{{person.name}}</span></td>
                <td><span>{{person.type}}</span></td>
                <td>
                    <div class="actions">
                        <router-link :to="{ name: RouteName.ADMIN_PANEL_INFO, params: { entity: 'person', entityId: person.id }}">
                            <Button>Info</Button>
                        </router-link>
                        <Button @click="removePerson(person.id)">Remove</Button>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>

</template>

<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import contractUsecase from "@usecase/contract/usecase";
import Person from "@model/person/model";
import {useRoute} from "vue-router";
import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";
import {AddPersonDTO, ContractRole} from "@repository/contract/repository";
import {RouteName} from "@app/router/types";

const route = useRoute()

const id = route.params.entityId as string

const persons = ref<Person[]>([])
const addDTOForm = reactive<{
    personId: string
    role: ContractRole | null
}>({
    personId: "",
    role: null
})

const addPerson = async () => {
    if (!addDTOForm.role) return

    const dto: AddPersonDTO = {
        personId: addDTOForm.personId,
        role: addDTOForm.role,
        contractId: id,
    }

    const data = await contractUsecase.addPerson(dto)

    if (data.Person.id) {
        addDTOForm.personId = ""
        addDTOForm.role = null
        persons.value.push(data.Person)
    }
}

const removePerson = async (personId: string) => {
    if (await contractUsecase.removePerson(id, personId)) {
        persons.value = persons.value.filter(p => p.id !== personId)
    }
}

onMounted(async() => {
    persons.value = await contractUsecase.getPersonsByContractId(id)
})
</script>

<style scoped>
.person-add {
    display: flex;
    gap: var(--gap-sm);
    align-items: center;
    max-width: max-content;

    .button {
        width: max-content;
    }

    .input {
        width: auto;
    }
}
</style>