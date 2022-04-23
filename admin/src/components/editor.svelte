<script lang="ts">
	import { onMount } from "svelte";
	import { EditorState, basicSetup } from "@codemirror/basic-setup";
	import { EditorView, keymap } from "@codemirror/view";
	import { indentWithTab } from "@codemirror/commands";
	import { json } from "@codemirror/lang-json";

	let elem: HTMLDivElement;
	let view: EditorView;
	onMount(() => {
		if (!elem) {
			console.log("no element to attach editor to");
			return;
		}
		view = new EditorView({
			state: EditorState.create({
				extensions: [
					basicSetup,
					json(),
					EditorView.theme({
						"&": { height: "580px", maxHeight: "580px", backgroundColor: "#fff" },
						".cm-scroller": { overflow: "auto" }
					})
				]
			}),
			parent: elem
		});
	});

	function input() {
		console.log(input);
	}
</script>

<div bind:this={elem} on:input={input} />

<style>
</style>
