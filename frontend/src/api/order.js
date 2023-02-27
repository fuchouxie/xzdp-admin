import request from "@/utils/request";

export function getOrderList(params) {
  return request({
    url: '/order/OrderList',
    method: 'get',
    params: params
  })
}


export function getOrderInfo(params) {
  return request({
    url: '/order/GetOrderInfo',
    method: 'get',
    params: params
  })
}
