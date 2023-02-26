import request from '@/utils/request'

export function getVoucherList(params) {
  return request({
    url: '/voucher/VoucherList',
    method: 'get',
    params: params
  })
}
export function removeVoucher(params) {
  return request({
    url: '/voucher/Remove',
    method: 'post',
    params: params
  })
}
export function batchRemoveVoucher(params) {
  return request({
    url: '/voucher/BatchRemove',
    method: 'post',
    params: params
  })
}
export function updateVoucher(params) {
  return request({
    url: '/voucher/Update',
    method: 'post',
    params: params
  })
}
export function addVoucher(params) {
  return request({
    url: '/voucher/Create',
    method: 'post',
    params: params
  })
}
