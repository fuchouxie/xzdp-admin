import request from '@/utils/request'

export function getShopList(params) {
  return request({
    url: '/shop/ShopList',
    method: 'get',
    params: params
  })
}
export function removeShop(params) {
  return request({
    url: '/shop/Remove',
    method: 'post',
    params: params
  })
}
export function batchRemoveShop(params) {
  return request({
    url: '/shop/BatchRemove',
    method: 'post',
    params: params
  })
}
export function updateShop(params) {
  return request({
    url: '/shop/Update',
    method: 'post',
    params: params
  })
}
export function addShop(params) {
  return request({
    url: '/shop/Create',
    method: 'post',
    params: params
  })
}
