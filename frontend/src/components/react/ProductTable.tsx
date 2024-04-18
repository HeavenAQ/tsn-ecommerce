import { client } from '../store.ts'
import { useQuery } from '@tanstack/react-query'
import { Language } from '../../../types/global.ts'
import { getProducts } from '../../api/products.ts'
import Spinner from './Spinner.tsx'
import { RxCrossCircled } from 'react-icons/rx'

export default function ProductTable() {
  const { data: products, isLoading, error } = useProductQuery(Language.JP)
  if (isLoading) {
    return (
      <div className="flex justify-center items-center w-full h-full">
        <Spinner />
      </div>
    )
  }

  return (
    <table className="mx-auto text-left shadow-md w-[1200px]">
      <thead className="h-14">
        <tr>
          <th className="w-20"></th>
          <th>名稱</th>
          <th>種類</th>
          <th>價格</th>
          <th>數量</th>
          <th>介紹</th>
          <th>庫存</th>
          <th>更新日期</th>
          <th>新增日期</th>
        </tr>
      </thead>
      <tbody className="bg-white">
        {products?.map(product => (
          <tr
            key={product.id}
            className="h-20 transition-colors duration-200 ease-in-out cursor-pointer hover:bg-zinc-200"
          >
            <td className="z-10">
              <RxCrossCircled className="mx-auto text-lg text-red-500" />
            </td>
            <td>{product.name}</td>
            <td>{product.category}</td>
            <td>{product.price}</td>
            <td>{product.quantity}</td>
            <td>{product.description}</td>
            <td>{product.status}</td>
            <td>{product.updated_at.slice(0, 10)}</td>
            <td>{product.created_at.slice(0, 10)}</td>
          </tr>
        ))}
      </tbody>
    </table>
  )
}

const useProductQuery = (language: Language) => {
  return useQuery(
    {
      queryKey: ['products'],
      queryFn: async () => getProducts(language)
    },
    client
  )
}
