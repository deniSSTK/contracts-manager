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
            <Input
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
            <Input
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
            <Input
                id="email"
                type="email"
                v-model="dto.email"
                placeholder="Optional"
            />
        </div>

        <div class="form-group">
            <label for="phone">Phone</label>
            <Input
                id="phone"
                type="tel"
                v-model="dto.phone"
                placeholder="Optional, E.164 format"
            />
        </div>

        <Button type="submit">Submit</Button>
    </form>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useRoute } from "vue-router";
import {CreatePersonDTO} from "@repository/person/repository";
import personUsecase from "@usecase/person/usecase";
import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

const route = useRoute();
const name = route.params.entity;

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
    width: 100%;
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
</style>
