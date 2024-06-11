import React, { DependencyList, useEffect, useMemo, useState } from 'react'
import StarterKit from '@tiptap/starter-kit';
import { Editor, EditorOptions, useEditor } from '@tiptap/react';
import { RichTextEditor } from '@mantine/tiptap';

function useForceUpdate() {
  const [, setValue] = useState(0);
  return () => setValue((value) => value + 1);
}


export const useCustomEditor = (
  options: Partial<EditorOptions> = {},
  deps: DependencyList = []
) => {
  const [editor, setEditor] = useState<Editor>(() => new Editor(options));
  const forceUpdate = useForceUpdate();

  useEffect(() => {
    let instance: Editor;

    if (editor.isDestroyed) {
      instance = new Editor(options);
      setEditor(instance);
    } else {
      instance = editor;
    }

    instance.on("transaction", () => {
      requestAnimationFrame(() => {
        requestAnimationFrame(() => {
          forceUpdate();
        });
      });
    });

    return () => {
      instance.destroy();
    };
  }, deps);

  return editor;
};

export default function RichTexInput({ initcontent }: { initcontent: string, }) {
  // const content =
  //   '<h2 style="text-align: left;">Welcome to Mantine rich text editor</h2><p><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a href="https://tiptap.dev/" rel="noopener noreferrer" target="_blank">Tiptap.dev</a> and supports all of its features:</p><ul><li>General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s> </li><li>Headings (h1-h6)</li><li>Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</li><li>Ordered and bullet lists</li><li>Text align&nbsp;</li><li>And all <a href="https://tiptap.dev/extensions" target="_blank" rel="noopener noreferrer">other extensions</a></li></ul>';

  const [content, onChange] = useState(initcontent)
  const options = {
    extensions: [
      StarterKit
    ],
    content,
    onUpdate(props) {
      const val = props.editor.getHTML()
      console.log(val);
      onChange(val)
    },
  }
  // const [editor, setEditor] = useState<Editor>(() => new Editor(options))
  const editor = useEditor(options, []);

  // useEffect(() => {
  //   if (editor && !editor.isDestroyed) {
  //     editor.chain().focus().setContent(content).run();
  //   }
  // }, [content, editor]);

  // useEffect(() => {
  //   console.log('Rich EFFECT!');
  //   let instance: Editor

  //   // // setEditor(instance)
  //   if (editor.isDestroyed) {
  //     // instance = new Editor({ ...options, content: content })
  //     // eslint-disable-next-line react-hooks/rules-of-hooks
  //     setEditor(editor)
  //   } else {
  //     instance = editor
  //   }
  //   return () => {
  //     instance.destroy()
  //   }
  // }, [])

  // const contentView = useMemo(() => <RichTextEditor.Content />
  //   , [content])

  return (
    <RichTextEditor editor={editor}>
      <RichTextEditor.Toolbar stickyOffset={60}>
        <RichTextEditor.ControlsGroup >
          <RichTextEditor.Bold />
          <RichTextEditor.Italic />
          {/* <RichTextEditor.Underline /> */}
          {/* <RichTextEditor.Strikethrough /> */}
          {/* <RichTextEditor.ClearFormatting />
          <RichTextEditor.Highlight /> */}
          {/* <RichTextEditor.Code /> */}
        </RichTextEditor.ControlsGroup>

        <RichTextEditor.ControlsGroup>
          <RichTextEditor.H1 />
          <RichTextEditor.H2 />
          <RichTextEditor.H3 />
          <RichTextEditor.H4 />
        </RichTextEditor.ControlsGroup>

        <RichTextEditor.ControlsGroup>
          <RichTextEditor.Blockquote />
          <RichTextEditor.Hr />
          <RichTextEditor.BulletList />
          <RichTextEditor.OrderedList />
          {/* <RichTextEditor.Subscript />
          <RichTextEditor.Superscript /> */}
        </RichTextEditor.ControlsGroup>

        {/* <RichTextEditor.ControlsGroup>
          <RichTextEditor.Link />
          <RichTextEditor.Unlink />
        </RichTextEditor.ControlsGroup> */}

        {/* <RichTextEditor.ControlsGroup>
          <RichTextEditor.AlignLeft />
          <RichTextEditor.AlignCenter />
          <RichTextEditor.AlignJustify />
          <RichTextEditor.AlignRight />
        </RichTextEditor.ControlsGroup> */}

        {/* <RichTextEditor.ControlsGroup>
          <RichTextEditor.Undo />
          <RichTextEditor.Redo />
        </RichTextEditor.ControlsGroup> */}
      </RichTextEditor.Toolbar>

      <RichTextEditor.Content></RichTextEditor.Content>
    </RichTextEditor>
  )

  // , [content])
}