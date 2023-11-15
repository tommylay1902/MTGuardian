<script lang="ts">
  export let name: string;
  export let valueType: string = "text";
  export let mode: string = "create";

  import FormStore from "$lib/store/Form";
  import { convertStringISO8601ToDateHtmlInput } from "$lib/utils/date";

  $: formatValue =
    valueType === "date" && mode === "update"
      ? convertStringISO8601ToDateHtmlInput($FormStore.data[name] as string)
      : $FormStore.data[name];
</script>

<label class="block text-sm font-medium text-white" for={name}>
  {name.toUpperCase()}
</label>
<input
  type={valueType}
  id={name}
  {name}
  value={formatValue}
  class="w-full px-3 py-2 border rounded-md shadow-sm"
/>
