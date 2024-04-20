import { useState, type FC } from 'react'
import type { Product } from '../../api/products'
import Button, { ButtonType } from './Button'
import Spinner from './Spinner'
import { useForm, type SubmitHandler } from 'react-hook-form'

interface ProductFormProps {
  product?: Product
  title: string
  isOpen: boolean
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
}

const ProductForm: FC<ProductFormProps> = ({
  product,
  title,
  isOpen,
  setIsOpen
}) => {
  return (
    <div
      className={`absolute right-0 left-0 bottom-10 top-16 mx-auto w-5/6 rounded-xl bg-gray-600/70 backdrop-blur ${isOpen ? 'opacity-100 z-10' : 'opacity-0 -z-10'} duration-300 ease-in-out overflow-y-scroll`}
    >
      <div className="m-auto">
        <h1 className="m-6 mt-28 text-4xl text-center text-zinc-100">
          {title}
        </h1>
        <hr className="mx-auto w-1/2 border-1" />
        <FormInput setIsOpen={setIsOpen} product={product} />
      </div>
    </div>
  )
}

interface FormInputProps {
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
  product?: Product
}

// for submitting form
enum StockStatus {
  InStock = '有現貨',
  OutOfStock = '無現貨',
  Discontinued = '斷貨'
}

const stockStatusToValue = (status: StockStatus) => {
  switch (status) {
    case StockStatus.InStock:
      return 'in-stock'
    case StockStatus.OutOfStock:
      return 'out-of-stock'
    case StockStatus.Discontinued:
      return 'discontinued'
  }
}

interface IFormInput {
  jpName: string
  chnName: string
  jpCategory: string
  chnCategory: string
  price: number
  stockStatus: StockStatus | string
  quantity: number
  jpDescription: string
  chnDescription: string
}

const FormInput: FC<FormInputProps> = ({ setIsOpen, product }) => {
  const [spinner, setSpinner] = useState<boolean>(false)
  const { handleSubmit, register } = useForm<IFormInput>()
  const onSubmit: SubmitHandler<IFormInput> = data => {
    setSpinner(true)
    data.stockStatus = stockStatusToValue(data.stockStatus as StockStatus)
  }

  return (
    <div className="pt-10 mx-auto w-4/5 lg:w-2/3 max-w-[1000px] h-[850px] md:h-[700px]">
      {spinner && (
        <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
          <Spinner />
        </div>
      )}
      <form
        className={`${spinner ? 'hidden' : ''}  p-6 w-3/4 rounded-xl mx-auto`}
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="flex flex-wrap mb-6">
          <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="chn-name"
              type="text"
              placeholder="品名（中文）"
              value={product?.name}
              {...register('chnName', { required: true })}
              required
            />
          </div>

          <div className="px-3 w-full md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="chn-category"
              type="text"
              placeholder="分類（中文）"
              {...register('chnCategory', { required: true })}
              required
            />
          </div>
        </div>
        <div className="flex flex-wrap mb-6">
          <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="jp-name"
              type="text"
              placeholder="品名（日文）"
              {...register('jPName', { required: true })}
              required
            />
          </div>
          <div className="px-3 w-full md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="jp-category"
              type="text"
              placeholder="種類（日文）"
              {...register('jpCategory', { required: true })}
              required
            />
          </div>
        </div>

        <div className="inline-flex mb-6 w-full">
          <div className="px-3 w-full md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="title"
              type="number"
              min="0"
              placeholder="售價（NTD）"
              {...register('price', { required: true })}
              required
            />
          </div>
          <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
            <select
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="service-type"
              {...register('stockStatus', { required: true })}
              required
            >
              <option> 有現貨 </option>
              <option> 無現貨 </option>
              <option> 斷貨 </option>
            </select>
          </div>
          <div className="px-3 w-full md:w-1/2">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="title"
              type="number"
              min="0"
              placeholder="庫存數量"
              {...register('quantity', { required: true })}
              required
            />
          </div>
        </div>

        <div className="inline-flex mb-6 w-full">
          <div className="px-3 mb-6 w-full md:mb-0">
            <textarea
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="chn-body"
              rows={10}
              placeholder="商品介紹（中文）"
              {...register('chnDescription', { required: true })}
              required
            />
          </div>
          <div className="px-3 mb-6 w-full md:mb-0">
            <textarea
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="jp-body"
              rows={10}
              placeholder="商品介紹（日文）"
              {...register('jpDescription', { required: true })}
              required
            />
          </div>
        </div>

        <div className="flex flex-wrap -mx-3 mb-6">
          <div className="flex gap-6 justify-center items-end w-full h-auto text-white">
            <Button type={ButtonType.Success} content="新增" />
            <Button
              type={ButtonType.Cancel}
              content="關閉"
              onClick={() => setIsOpen(false)}
            />
          </div>
        </div>
      </form>
    </div>
  )
}

export default ProductForm
