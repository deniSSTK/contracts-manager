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
                <label for="username">Password</label>
                <div class="input-container">
                    <Input
                        id="password"
                        v-model="dto.password"
                        :type="canSeePassword ? 'text' : 'password'"
                        minlength="8"
                        maxlength="50"
                        placeholder="Input password"
                        required
                    />
                    <Button type="button" @click="canSeePassword = !canSeePassword">
                        {{ canSeePassword ? 'close' : 'see'}}
                    </Button>
                </div>
            </div>

            <Button type="submit" :disabled="!canSendReq">Log In</Button>
            <router-link :to="{ name: RouteName.SIGNUP }">Don't have an account? <span class="primary">Sign up</span></router-link>
        </form>
    </main>
</template>

<script lang="ts" setup>
import {reactive, computed, ref} from "vue";
import authUsecase from "@usecase/auth/usecase";
import {ILoginDTO} from "@repository/auth/repository";
import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

import "./style.css"
import {RouteName} from "@app/router/types";

const canSeePassword = ref<boolean>(false)

const dto = reactive<ILoginDTO>({
    usernameOrEmail: "",
    password: "",
})

const canSendReq = computed(() => dto.usernameOrEmail.length >= 8 && dto.password.length >= 8)

const handleSubmit = async () => {
    if (await authUsecase.login(dto)) {
        window.location.reload();
    }
}
</script>