<template>
    <form @submit.prevent="handleSubmit">
        <div>
            <label for="username">Username:</label>
            <input id="username" v-model="dto.username" type="text" required />
        </div>
        <div>
            <label for="email">Email:</label>
            <input id="email" v-model="dto.email" type="email" required />
        </div>
        <div>
            <label for="password">Password:</label>
            <input id="password" v-model="dto.password" type="password" required />
        </div>
        <button type="submit">Sign Up</button>
    </form>
    denisdrawler@gmail.com
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import authUsecase from "@usecase/auth/usecase";
import {ISignupDTO} from "@repository/auth/repository";

const dto = reactive<ISignupDTO>({
    username: "",
    email: "",
    password: "",
})

const handleSubmit = async () => {
    await authUsecase.signup(dto)
}
</script>