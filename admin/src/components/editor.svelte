<script lang="ts">
	import { onMount } from "svelte";
	import { EditorState, basicSetup } from "@codemirror/basic-setup";
	import { EditorView, keymap } from "@codemirror/view";
	import { indentWithTab } from "@codemirror/commands";
	import { oneDark } from "@codemirror/theme-one-dark";
	import { json } from "@codemirror/lang-json";

	export let height = 300;
	export let editable = "true";
	let elem: HTMLDivElement;
	let view: EditorView;

	onMount(() => {
		if (!elem) {
			console.log("no element to attach editor to");
			return;
		}
		let extensions = [
			basicSetup,
			keymap.of([indentWithTab]),
			json(),
			oneDark,
			EditorView.theme({
				"&": {
					color: "#e7eaf0",
					height: `${height}px`
				},
				".cm-content": {
					color: "#e7eaf0"
				}
			})
		];
		if (editable === "false") {
			extensions.push(EditorView.editable.of(false));
		}
		view = new EditorView({
			state: EditorState.create({
				extensions: [...extensions]
			}),
			parent: elem
		});
		if (editable === "false") {
			console.log("not editable");
			let content = view.contentDOM;
			if (content) {
				console.log("setting editable to false");
				content.contentEditable = "false";
				console.log("editable: ", content.contentEditable);
			}
		}
	});

	function input() {
		console.log(input);
	}
</script>

<div bind:this={elem} />
