<template>
    <main class="page">
        <form @submit.prevent="handleSubmit">
            <h1>Log In</h1>

            <div>
                <label for="username">Email Or Username</label>
                <Input
                    id="username"
                    v-model="dto.usernameOrEmail"
                    type="text"
                    maxlength="50"
                    placeholder="Input your email or username"
                    required
                />
            </div>

            <div>
                <label for="password">Password</label>
                <Input
                    id="password"
                    v-model="dto.password"
                    type="password"
                    min="8"
                    maxlength="50"
                    placeholder="Input your password"
                    required
                />
            </div>
            <Button type="submit" :disabled="!canSendReq">Log In</Button>
            <router-link :to="{ name: RouteName.SIGNUP }">Don't have an account? <span class="primary">Register</span></router-link>
        </form>
    </main>
</template>

<script lang="ts" setup>
import {reactive, computed} from "vue";
import authUsecase from "@usecase/auth/usecase";
import {ILoginDTO} from "@repository/auth/repository";
import router, {RouteName} from "../../app/routes";
import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

import "./style.css"

const dto = reactive<ILoginDTO>({
    usernameOrEmail: "",
    password: "",
})

const canSendReq = computed(() => dto.usernameOrEmail.length >= 8 && dto.password.length >= 8)

const handleSubmit = async () => {
    if (await authUsecase.login(dto)) {
        await router.push({
            name: RouteName.DASHBOARD
        })
    }
}
</script>