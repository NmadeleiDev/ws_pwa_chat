<template>
  <div class="alert-container" v-if="isActive">
    <h3>{{ message }}</h3>
  </div>
</template>

<script>
    import {alertNotifierChannel} from "../../main";

    export default {
      name: "Alert",
      data() {
        return {
          message: '',
          isActive: false
        }
      },
      created() {
        alertNotifierChannel.$on('show', (event) => {
          this.isActive = true;
          this.message = event.message;
        });
      }
    }
</script>

<style scoped>
.alert-container {
  display: block;
  position: absolute;
  top: 50px;
  right: 50px;
}
</style>
