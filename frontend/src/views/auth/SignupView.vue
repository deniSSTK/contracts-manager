<template>
    <main class="page">
        <form @submit.prevent="handleSubmit">
            <h1>Sign Up</h1>

            <div>
                <label for="username">Username</label>
                <Input
                    id="username"
                    v-model="dto.username"
                    type="text"
                    maxlength="50"
                    placeholder="Input your username"
                    required
                />
            </div>

            <div>
                <label for="email">Email</label>
                <Input
                    id="email"
                    v-model="dto.email"
                    type="email"
                    maxlength="100"
                    placeholder="Input your email"
                    required
                />
            </div>

            <div>
                <label for="password">Password</label>
                <div class="input-container">
                    <Input
                        id="password"
                        v-model="dto.password"
                        :type="canSeePassword ? 'text' : 'password'"
                        minlength="8"
                        maxlength="50"
                        placeholder="Input your password"
                        required
                    />
                    <Button type="button" @click="canSeePassword = !canSeePassword">
                        {{ canSeePassword ? 'close' : 'see'}}
                    </Button>
                </div>
            </div>

            <div>
                <label for="password">Confirm password</label>
                <div class="input-container">
                    <Input
                        id="confirm-password"
                        v-model="confirmPassword"
                        :type="canSeeConfirmPassword ? 'text' : 'password'"
                        minlength="8"
                        maxlength="50"
                        placeholder="Confirm password"
                        required
                    />
                    <Button type="button" @click="canSeeConfirmPassword = !canSeeConfirmPassword">
                        {{ canSeeConfirmPassword ? 'close' : 'see'}}
                    </Button>
                </div>
            </div>

            <Button type="submit" :disabled="!canSendReq">
                Sign Up
            </Button>

            <router-link :to="{ name: RouteName.LOGIN }">Already have an account? <span class="primary">Login</span></router-link>
        </form>
    </main>
</template>

<script lang="ts" setup>
import {reactive, computed, ref} from "vue";
import authUsecase from "@usecase/auth/usecase";
import { ISignupDTO } from "@repository/auth/repository";

import Button from "@component/ui/button/Button.vue";
import Input from "@component/ui/input/Input.vue";

import "./style.css";
import {RouteName} from "@app/router/types";

const canSeePassword = ref(false);
const canSeeConfirmPassword = ref(false);

const confirmPassword = ref<string>("");

const dto = reactive<ISignupDTO>({
    username: "",
    email: "",
    password: "",
});

const canSendReq = computed(() =>
    dto.username.length >= 3 &&
    dto.email.length >= 5 &&
    dto.password.length >= 8 &&
    dto.password === confirmPassword.value
);

const handleSubmit = async () => {
    if (confirmPassword.value !== dto.password) return;

    if (await authUsecase.signup(dto)) {
        window.location.reload();
    }
};
</script>
