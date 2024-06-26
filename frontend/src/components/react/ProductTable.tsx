import { client } from '../store.ts'
import { useMutation, useQuery } from '@tanstack/react-query'
import { Language } from '../../../types/global.ts'
import { getProducts, deleteProduct, type Product } from '../../api/products.ts'
import Spinner from './Spinner.tsx'
import { RxCrossCircled } from 'react-icons/rx'
import toast, { Toaster } from 'react-hot-toast'
import Button, { ButtonType } from './Button.tsx'
import { useState } from 'react'
import ProductForm from '../react/ProductForm.tsx'

const ProductTableHeader = () => {
  const [formOpen, setFormOpen] = useState(false)
  return (
    <div className="inline-flex justify-center items-center mb-10 w-[1300px]">
      <h1 className="mr-auto">Products</h1>
      <Button
        type={ButtonType.Primary}
        content="新增"
        onClick={() => setFormOpen(true)}
      />
      <ProductForm title="新增商品" isOpen={formOpen} setIsOpen={setFormOpen} />
    </div>
  )
}

const ProductTable = () => {
  const { data: products, isLoading, error } = useProductQuery(Language.JP)

  // if error, show something went wrong
  if (error) {
    toast.error('發生錯誤請聯絡管理員')
  }

  // if loading, show spinner
  if (isLoading) {
    return (
      <div className="flex justify-center items-center w-full h-full">
        <Spinner />
      </div>
    )
  }

  return (
    <div className="mx-auto rounded-2xl w-[1300px]">
      <Toaster
        position="top-right"
        gutter={12}
        containerStyle={{ margin: '8px' }}
        toastOptions={{
          success: {
            duration: 3000
          },
          error: {
            duration: 5000
          },
          style: {
            fontSize: '16px',
            maxWidth: '500px',
            padding: '16px 24px',
            backgroundColor: '#333',
            color: '#fff'
          }
        }}
      />
      <ProductTableHeader />
      <table className="overflow-hidden w-full text-left bg-white shadow-md">
        <thead className="h-14 bg-zinc-100">
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
            <ProductRow key={product.id} product={product} />
          ))}
        </tbody>
      </table>
    </div>
  )
}

interface Prop {
  product: Product
}

const ProductRow: React.FC<Prop> = ({ product }) => {
  // for deleting product
  const { isPending, mutate } = useMutation(
    {
      mutationFn: deleteProduct,
      onSuccess: () => {
        toast.success('商品已成功刪除')
        client.invalidateQueries({
          queryKey: ['products', Language.JP]
        })
      },
      onError: err => {
        toast.error(err.message)
      }
    },
    client
  )

  // if deleting, show spinner
  if (isPending) {
    return (
      <tr className="h-20">
        <td colSpan={9} className="flex justify-center items-center">
          <Spinner />
        </td>
      </tr>
    )
  }

  // return product rows
  return (
    <tr
      className="mb-4 h-20 transition-colors duration-200 ease-in-out cursor-pointer hover:bg-zinc-200"
      onClick={() => console.log(product.id)}
    >
      <td
        className="z-10 rounded-tl-2xl rounded-bl-2xl duration-200 ease-in-out hover:bg-red-300 group/delete"
        onClick={() => mutate(product.id)}
      >
        <RxCrossCircled className="mx-auto text-lg text-red-500 group-hover/delete:text-white" />
      </td>
      <td>{product.name}</td>
      <td>{product.category}</td>
      <td>{product.price}</td>
      <td>{product.quantity}</td>
      <td>{product.description}</td>
      <td>{product.status}</td>
      <td>{product.updated_at.slice(0, 10)}</td>
      <td className="rounded-tr-2xl rounded-br-2xl">
        {product.created_at.slice(0, 10)}
      </td>
    </tr>
  )
}

const useProductQuery = (language: Language) => {
  return useQuery(
    {
      queryKey: ['products', language],
      queryFn: async () => getProducts(language)
    },
    client
  )
}

export default ProductTable
