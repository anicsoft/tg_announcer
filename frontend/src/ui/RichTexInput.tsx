import React, { useState } from 'react'
import StarterKit from '@tiptap/starter-kit';
import { useEditor } from '@tiptap/react';
import { RichTextEditor } from '@mantine/tiptap';

export default function RichTexInput({ initcontent }: { initcontent: string, }) {
  // const content =
  //   '<h2 style="text-align: left;">Welcome to Mantine rich text editor</h2><p><code>RichTextEditor</code> component focuses on usability and is designed to be as simple as possible to bring a familiar editing experience to regular users. <code>RichTextEditor</code> is based on <a href="https://tiptap.dev/" rel="noopener noreferrer" target="_blank">Tiptap.dev</a> and supports all of its features:</p><ul><li>General text formatting: <strong>bold</strong>, <em>italic</em>, <u>underline</u>, <s>strike-through</s> </li><li>Headings (h1-h6)</li><li>Sub and super scripts (<sup>&lt;sup /&gt;</sup> and <sub>&lt;sub /&gt;</sub> tags)</li><li>Ordered and bullet lists</li><li>Text align&nbsp;</li><li>And all <a href="https://tiptap.dev/extensions" target="_blank" rel="noopener noreferrer">other extensions</a></li></ul>';

  const [content, onChange] = useState(initcontent)

  const editor = useEditor({
    extensions: [
      StarterKit
    ],
    content,
    onUpdate(props) {
      const val = props.editor.getHTML()
      console.log(val);
      onChange(val)
    },
  });
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

      <RichTextEditor.Content />
    </RichTextEditor>
  )
}
