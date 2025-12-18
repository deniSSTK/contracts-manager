<template>
    <form @submit.prevent="handleSubmit">
        <div>
            <label for="username">Email Or Username:</label>
            <input id="username" v-model="dto.usernameOrEmail" type="text" required />
        </div>
        <div>
            <label for="password">Password:</label>
            <input id="password" v-model="dto.password" type="password" required />
        </div>
        <button type="submit">Log In</button>
    </form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import authUsecase from "@usecase/auth/usecase";
import {ILoginDTO} from "@repository/auth/repository";
import router, {RouteName} from "../../app/routes";

const dto = reactive<ILoginDTO>({
    usernameOrEmail: "",
    password: "",
})

const handleSubmit = async () => {
    if (await authUsecase.login(dto)) {
        await router.push({
            name: RouteName.DASHBOARD
        })
    }
}
</script>