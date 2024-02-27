<script>
  import { onMount } from "svelte";
  import ace from "ace-builds/src-noconflict/ace";
  import "ace-builds/src-noconflict/theme-one_dark";
  import "ace-builds/src-noconflict/mode-javascript";

  let editor;

  onMount(() => {
    editor = ace.edit("ace-editor");
    editor.setTheme("ace/theme/one_dark");
    editor.getSession().setMode("ace/mode/javascript");
    editor.setOptions({
      showPrintMargin: false,
    });
  });
</script>

<div id="ace-editor" class="flex justify-center items-center flex-grow">
</div>
