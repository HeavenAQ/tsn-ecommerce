import { useStore } from '@nanostores/react'
import { client } from '../store.ts'
import { useQuery } from '@tanstack/react-query'

export default function ProductTable() {
  const products = useProductQuery()
  return <div>ProductTable</div>
}

const useProductQuery = () => {
  return useQuery(
    {
      queryKey: ['products'],
      queryFn: async () => {
        const response = await fetch('/api/products')
        return response.json()
      }
    },
    client
  )
}
