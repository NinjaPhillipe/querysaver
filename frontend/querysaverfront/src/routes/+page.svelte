<script lang="ts">
  import TagsList from "../components/tags-list.svelte";
</script>

<div class="flex min-h-screen">
  <div class="w-1/4">
    <div>SEARCH BAR</div>
    <div>SEARCH RESULTS TAGS</div>
    <div><TagsList /></div>
    <div>CORESPONDING FILE</div>
  </div>
  <div class="w-3/4">
    <div class="h-1/6">SEARCH STATE</div>
    <div class="h-5/6">PREVIEW FILE</div>
  </div>
</div>
