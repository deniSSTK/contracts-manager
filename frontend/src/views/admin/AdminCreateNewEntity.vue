<template>
    <form @submit.prevent="handleSubmit" class="form-container">
        <h2>Create {{ name }}</h2>

        <div class="form-group">
            <label for="type">Type</label>
            <select id="type" v-model="dto.type" required>
                <option value="">Select type</option>
                <option value="individual">Individual</option>
                <option value="entity">Entity</option>
            </select>
        </div>

        <div class="form-group">
            <label for="name">Name</label>
            <input
                id="name"
                type="text"
                v-model="dto.name"
                required
                minlength="2"
                maxlength="255"
            />
        </div>

        <div class="form-group">
            <label for="code">Code</label>
            <input
                id="code"
                type="text"
                v-model="dto.code"
                required
                minlength="3"
                maxlength="50"
            />
        </div>

        <div class="form-group">
            <label for="email">Email</label>
            <input
                id="email"
                type="email"
                v-model="dto.email"
                placeholder="Optional"
            />
        </div>

        <div class="form-group">
            <label for="phone">Phone</label>
            <input
                id="phone"
                type="tel"
                v-model="dto.phone"
                placeholder="Optional, E.164 format"
            />
        </div>

        <button type="submit">Submit</button>
    </form>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { importName } from "@utils/utils";
import { useRoute } from "vue-router";
import {CreatePersonDTO} from "@repository/person/repository";
import personUsecase from "@usecase/person/usecase";

const route = useRoute();
const name = importName(route);

const dto = reactive<CreatePersonDTO>({
    type: "",
    name: "",
    code: "",
    email: null,
    phone: null,
});

async function handleSubmit() {
    console.log(await personUsecase.create(dto))
}
</script>

<style scoped>
.form-container {
    max-width: 400px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.form-group {
    display: flex;
    flex-direction: column;
}

label {
    font-weight: 500;
    margin-bottom: 4px;
}

input, select, button {
    padding: 8px;
    font-size: 14px;
}

button {
    cursor: pointer;
    background-color: #3b82f6;
    color: white;
    border: none;
    border-radius: 4px;
    margin-top: 12px;
}
</style>
