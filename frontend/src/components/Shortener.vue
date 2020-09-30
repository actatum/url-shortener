<template>
  <div>
    <h1>Url Shortener</h1>
    <form @submit.prevent="create">
      <label>Url</label>
      <br />
      <input v-model="url" />
      <br />
      <button>Shorten</button>
    </form>
    <p v-if="result.url">{{ result }}</p>
  </div>
</template>

<script>
import { ref } from "vue";
import axios from "axios";

export default {
  setup() {
    const url = ref("");
    const result = ref({});
    console.log(result.value);

    const create = async () => {
      console.log(url.value);
      const req = {
        url: url.value,
      };
      try {
        const resp = await axios.post("http://localhost:8080/url", req);
        console.log(resp);
        result.value = resp.data;
      } catch (e) {
        console.log(e.response.data.message);
      }
    };

    return {
      url,
      create,
      result,
    };
  },
};
</script>

<style scoped>
</style>