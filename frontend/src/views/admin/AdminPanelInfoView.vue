<template>
    <h1 v-if="isNew">Create new</h1>
    <h1 v-else>Info</h1>

    <div v-if="loading">Loading...</div>

    <form v-else class="info" @submit.prevent="handleSubmit" >
        <template v-for="row in entityConfig.rows" :key="row.key">
            <div
                v-if="!isNew || (isNew && row.canChange)"
                class="info-row"
            >
                <label class="label">
                    {{ row.label }}<span v-if="!row.optional && row.canChange" class="red">*</span>
                </label>
                <select
                    v-if="row.values"
                    v-model="form[row.key]"
                    :disabled="!row.canChange"
                    :required="!row.optional"
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
                    :required="!row.optional"
                    :minlength="row.min"
                    :maxlength="row.max"
                />
            </div>
        </template>

        <span v-if="created">
            <span v-if="createdError" class="red">An error occurred</span>
            <span v-else class="primary">Success</span>
        </span>
        <Button v-else type="submit">{{ isNew ? 'Create' : 'Update' }}</Button>
    </form>
</template>

<script lang="ts" setup>
import {useRoute} from "vue-router";
import {onMounted, reactive, ref, computed} from "vue";
import infoEntities from "@entity/info";
import Input from "@component/ui/input/Input.vue";
import Button from "@component/ui/button/Button.vue";
import {RouteName} from "@app/router/types";

const route = useRoute()

const id = route.params.entityId as string
const entity = route.params.entity as string
const isNew = route.name === RouteName.ADMIN_PANEL_NEW

const entityConfig = computed(() => infoEntities[entity])

const loading = ref<boolean>(!isNew)
const created = ref<boolean>(false)
const createdError = ref<boolean>(false)
const form = reactive<Record<string, any>>({})

async function fetchData() {
    if (!isNew) {
        loading.value = true

        const data = await entityConfig.value.usecase.get(id)

        entityConfig.value.rows.forEach(row => {
            form[row.key] = (data as any)[row.key] ?? null
        })

        loading.value = false
    }
}

const handleSubmit = async () => {
    if (isNew && entityConfig.value.canCreate) {
        if (!await entityConfig.value.usecase.create(form)) {
            createdError.value = true
        }
        created.value = true
    }
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
}
</style>
