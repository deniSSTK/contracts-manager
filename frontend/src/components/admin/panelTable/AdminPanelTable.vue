<template>
    <div class="table-container">

        <div class="filters">
            <Input
                v-for="f in entity.filters"
                :key="f.key"
                v-model="filters[f.key]"
                :placeholder="f.placeholder"
            />
            <Button @click="applyFilters">Apply</Button>
            <Button @click="resetFilters">Reset</Button>
        </div>

        <table>
            <thead>
            <tr>
                <th v-for="c in entity.columns" :key="c.key">
                    {{c.actions ? 'Action' : c.label  }}
                </th>
            </tr>
            </thead>

            <tbody>
            <tr v-for="row in data" :key="row.id">
                <td v-for="c in entity.columns" :key="c.key">
                    <div v-if="c.actions" class="actions">
                        <Button
                            v-for="(a, i) in c.actions"
                            :key="i"
                            size="sm"
                            @click="a.callback(row)"
                        >
                            {{ a.label }}
                        </Button>
                    </div>

                    <span v-else>
                        {{ c.optional && !row[c.key] ? "-" : row[c.key]  }}
                    </span>
                </td>
            </tr>
            </tbody>
        </table>

        <div class="pagination">
            <Button :disabled="page === 1" @click="goToPage(page - 1)">
                Prev
            </Button>

            <Button
                v-for="p in totalPages"
                :key="p"
                :class="{ active: p === page }"
                @click="goToPage(p)"
            >
                {{ p }}
            </Button>

            <Button
                :disabled="page === totalPages || totalPages === 0"
                @click="goToPage(page + 1)"
            >
                Next
            </Button>
        </div>

    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, onMounted, watch } from "vue";
import { useRoute } from "vue-router";

import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

import { entityRegistry } from "@entity/tables";

import "./admin-panel-table.css";

const route = useRoute();

const entity = computed(() => {
    const key = route.params.entity;
    const e = entityRegistry[key as keyof typeof entityRegistry];

    if (!e) {
        throw new Error(`Unknown entity: ${key}`);
    }

    return e;
});

const data = ref<any[]>([]);
const page = ref(1);
const limit = 10;
const totalPages = ref(1);

const filters = reactive<Record<string, string>>({});

entity.value.filters.forEach(f => {
    filters[f.key] = "";
});

async function fetchData() {
    const params = new URLSearchParams({
        page: page.value.toString(),
        limit: limit.toString(),
        ...Object.fromEntries(
            Object.entries(filters).filter(([_, v]) => v)
        )
    });

    //@ts-ignore
    const result = await entity.value.usecase.list(params);

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

function resetFilters() {
    Object.keys(filters).forEach(key => {
        filters[key] = "";
    });

    fetchData();
}

watch(
    () => route.name,
    () => {
        page.value = 1;
        fetchData();
    }
);

onMounted(async () => {
    await fetchData()
});
</script>
