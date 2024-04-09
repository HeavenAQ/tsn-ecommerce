import { type ReactNode } from 'react'

interface Props {
  title: string
  href: string
  children: ReactNode 
}

export default function SideBarItem({ title, href, children }: Props) {
  return (
    <a href={href} className="">
      <div className="inline-flex gap-2 justify-start items-center py-3 px-3 w-full h-14 text-xl font-medium rounded-lg duration-150 cursor-pointer text-zinc-600 visited:bg-zinc-100 hover:bg-zinc-100">
        {children}{title}
      </div>
    </a>
  )
}
