<template>
    <router-view />
<!--    <Footer />-->

    <ContractPreviewPopup />
</template>
<script setup lang="ts">
import ContractPreviewPopup from "@component/popup/contractPreview/ContractPreviewPopup.vue";
import Footer from "@component/footer/Footer.vue";
</script>