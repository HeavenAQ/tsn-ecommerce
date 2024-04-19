import type { Language } from '../../types/global'
import axios from 'axios'

export enum ProductStatus {
  IN_STOCK = 'in-stock',
  OUT_OF_STOCK = 'out-of-stock',
  DISCONTINUED = 'discontinued'
}

export interface Product {
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
      return response.data as Product[]
    })
    .catch(function (error) {
      console.log(error)
      return []
    })
}

interface DeleteProductResponse {
  message: string
}

export const deleteProduct = async (id: string) => {
  return axios
    .delete(`/api/v1/products/${id}`, {
      baseURL: 'http://localhost:8080',
      responseType: 'json'
    })
    .then(function (response) {
      return response.data as DeleteProductResponse
    })
    .catch(function (error) {
      console.log(error)
      return null
    })
}
