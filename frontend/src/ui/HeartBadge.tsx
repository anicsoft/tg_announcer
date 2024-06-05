import React, { StyleHTMLAttributes } from 'react'

export default function HeartBadge({ count, style }: { count: number, style: any }) {
  return (
    <div style={style}>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="red" stroke="red" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-heart">
        <g>
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572" />
          <text x="12" y="16" font-size="12" text-anchor="middle" fill="white" stroke="none">{count}</text>
        </g>
      </svg>
    </div>
  )
}
