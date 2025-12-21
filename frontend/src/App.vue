<template>
    <Header />
    <router-view />

    <ContractPreviewPopup />
</template>
<script setup lang="ts">
import ContractPreviewPopup from "@component/popup/contractPreview/ContractPreviewPopup.vue";
import Header from "@component/header/Header.vue";
</script>