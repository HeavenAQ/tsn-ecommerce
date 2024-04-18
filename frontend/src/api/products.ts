import type { Language } from '../../types/global'
import axios from 'axios'

enum ProductStatus {
  IN_STOCK = 'in-stock',
  OUT_OF_STOCK = 'out-of-stock',
  DISCONTINUED = 'discontinued'
}

interface Product {
  id: string
  name: string
  price: number
  quantity: number
  status: ProductStatus
  category: string
  description: string
  created_at: string
  updated_at: string
  imageURLs: string[]
}

export const getProducts = async (language: Language) => {
  return axios
    .get('/api/v1/products', {
      baseURL: 'http://localhost:8080',
      params: {
        limit: 10,
        offset: 0,
        language: language
      },
      responseType: 'json'
    })
    .then(function (response) {
      console.log(response)
      return response.data as Product[]
    })
    .catch(function (error) {
      console.log(error)
      return []
    })
}
