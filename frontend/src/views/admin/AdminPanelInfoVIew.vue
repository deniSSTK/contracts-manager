<template>
    <h1>Info</h1>

    <div v-if="loading">Loading...</div>

    <div v-else class="info">
        <div
            v-for="row in entityConfig.rows"
            :key="row.key"
            class="info-row"
        >
            <label class="label">
                {{ row.label }}<span v-if="!row.optional && row.canChange" class="red">*</span>
            </label>

            <select
                v-if="row.values"
                v-model="form[row.key]"
                :disabled="!row.canChange"
            >
                <option
                    v-for="v in row.values"
                    :key="v"
                    :value="v"
                >
                    {{ v }}
                </option>
            </select>

            <Input
                v-else
                v-model="form[row.key]"
                :disabled="!row.canChange"
            />
        </div>

        <Button>Update</Button>
    </div>
</template>

<script lang="ts" setup>
import {useRoute} from "vue-router";
import {onMounted, reactive, ref, computed} from "vue";
import infoEntities from "@entity/info";
import Input from "@component/ui/input/Input.vue";
import Button from "@component/ui/button/Button.vue";

const route = useRoute()

const id = route.params.entityId as string
const entity = route.params.entity as string

const entityConfig = computed(() => infoEntities[entity])

const loading = ref(true)
const form = reactive<Record<string, any>>({})

async function fetchData() {
    loading.value = true

    const data = await entityConfig.value.usecase.get(id)

    entityConfig.value.rows.forEach(row => {
        form[row.key] = (data as any)[row.key] ?? null
    })

    loading.value = false
}

onMounted(fetchData)
</script>

<style scoped>
.info {
    display: flex;
    flex-direction: column;
    gap: var(--gap-sm);

    .info-row {
        display: grid;
        grid-template-columns: 150px 1fr;
        gap: var(--gap-sm);
        align-items: center;
    }

    select, .input {
        max-width: 30%;
    }

    .button {
        width: max-content;
    }

    .label {
        opacity: .7;
    }
}
</style>
