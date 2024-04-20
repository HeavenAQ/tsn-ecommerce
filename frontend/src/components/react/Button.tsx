import React, { type FC } from 'react'

export enum ButtonType {
  Primary = 'primary',
  Success = 'success',
  Cancel = 'cancel'
}

interface Props {
  content: string
  type: ButtonType
  onClick?: () => void
}

const Button: FC<Props> = ({ content, type, onClick }) => {
  if (type === ButtonType.Primary) {
    return (
      <button
        className="z-10 py-2 px-4 font-semibold bg-transparent rounded border transition-all duration-300 hover:text-white hover:border-transparent text-zinc-700 border-zinc-500 hover:bg-zinc-500"
        onClick={onClick}
      >
        {content}
      </button>
    )
  }

  if (type === ButtonType.Cancel) {
    return (
      <button
        type="reset"
        className="py-2 px-4 font-semibold text-white bg-transparent rounded border border-red-500 transition-all duration-300 hover:bg-red-500 hover:border-transparent"
        onClick={onClick}
      >
        {content}
      </button>
    )
  }

  if (type === ButtonType.Success) {
    return (
      <button
        className="py-2 px-4 font-semibold text-white bg-transparent rounded border border-green-500 transition-all duration-300 hover:bg-green-400 hover:border-transparent"
        onClick={onClick}
      >
        {content}
      </button>
    )
  }

  return (
    <button
      className="py-2 px-4 font-semibold bg-transparent rounded border hover:text-white hover:border-transparent text-zinc-700 border-zinc-500 hover:bg-zinc-500"
      onClick={onClick}
    >
      {content}
    </button>
  )
}

export default Button
