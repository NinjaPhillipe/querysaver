<script lang="ts">
  import { onMount } from "svelte";
  import { searchResultsTags } from "../lib/store";

  let searchInput = "";

  async function searchTag(name: string) {
    console.log("search");

    if (name === "") {
      const response = await fetch("http://localhost:8080/tags/query/all");

      searchResultsTags.set(await response.json());
      return;
    } else {
      const response = await fetch(
        `http://localhost:8080/tags/query/search/${name}`
      );

      searchResultsTags.set(await response.json());
    }
  }

  onMount(async () => {
    const response = await fetch("http://localhost:8080/tags/query/all");

    searchResultsTags.set(await response.json());
  });
</script>

<input
  type="search"
  bind:value={searchInput}
  class="input"
  on:change={() => searchTag(searchInput)}
/>
