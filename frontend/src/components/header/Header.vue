<template>
    <header>
        <span>Contract Manager</span>
        <div class="links-wrapper">
            <a href="https://github.com/deniSSTK/contracts-manager" target="_blank">
                <Icon icon="simple-icons:github"></Icon>
                GitHub
            </a>
            <a href="https://www.linkedin.com/in/denis-tkachenko-b7a801327/" target="_blank">
                <Icon icon="mdi:linkedin"></Icon>
                LinkedIn
            </a>
        </div>
        <Button v-if="!AuthPages.includes($route.name as RouteName)" @click="authStore.logout()">Logout</Button>
    </header>
</template>

<script lang="ts" setup>
import {Icon} from "@iconify/vue";
import useAuthStore from "@store/auth/store";
import Button from "@component/ui/button/Button.vue";
import {AuthPages, RouteName} from "@app/router/types";

const authStore = useAuthStore()

</script>

<style scoped>
header {
    width: 100%;
    height: 50px;

    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 var(--gap-md);

    background: var(--black);

    span {
        color: var(--white);
    }

    &:has(a:hover) a:not(:hover) {
        opacity: .5;
    }

    .links-wrapper {
        display: flex;
        align-items: center;
        gap: var(--gap-md);
        position: absolute;
        left: 50%;
        transform: translateX(-50%);
    }

    a {
        width: max-content;

        display: flex;
        align-items: center;
        gap: var(--gap-sm);

        transition: var(--tr-duration);

        color: var(--white);

        &:hover {
            color: var(--color-primary);
        }
    }

    .button {
        width: max-content;
    }
}
</style>