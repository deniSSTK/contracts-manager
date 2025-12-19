<template>
    <div class="table-container">
        <div class="filters">
            <Input v-model="filters.name" placeholder="Name" />
            <Input v-model="filters.type" placeholder="Type" />
            <Input v-model="filters.code" placeholder="Code" />
            <Button @click="applyFilters">Apply</Button>
        </div>

        <table>
            <thead>
            <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Code</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="person in data" :key="person.id">
                <td>{{ person.name }}</td>
                <td>{{ person.type }}</td>
                <td>{{ person.code }}</td>
            </tr>
            </tbody>
        </table>

        <div class="pagination">
            <Button :disabled="page === 1" @click="goToPage(page - 1)">Prev</Button>

            <Button
                v-for="p in totalPages"
                :key="p"
                :class="{ active: p === page }"
                @click="goToPage(p)"
            >
                {{ p }}
            </Button>

            <Button :disabled="page === totalPages || totalPages === 0" @click="goToPage(page + 1)">Next</Button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from "vue";
import personUsecase from "@usecase/person/usecase";
import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

import "./admin-panel-table.css"

interface Person {
    id: number;
    name: string;
    type: string;
    code: string;
}

const data = ref<Person[]>([]);
const page = ref(1);
const limit = 20;
const totalPages = ref(1);

const filters = reactive({
    name: "",
    type: "",
    code: ""
});

async function fetchData() {
    const params = new URLSearchParams({
        page: page.value.toString(),
        limit: limit.toString(),
        ...(filters.name && { name: filters.name }),
        ...(filters.type && { type: filters.type }),
        ...(filters.code && { code: filters.code })
    });

    const result = await personUsecase.list(params)

    data.value = result.data;
    totalPages.value = Math.ceil(result.total / result.limit);
}

function goToPage(p: number) {
    if (p < 1 || p > totalPages.value) return;
    page.value = p;
    fetchData();
}

function applyFilters() {
    page.value = 1;
    fetchData();
}

onMounted(() => {
    fetchData();
});
</script>