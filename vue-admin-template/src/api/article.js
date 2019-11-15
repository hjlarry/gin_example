import request from '@/utils/request'

export function fetchList(params) {
  return request({
    url: '/articles',
    method: 'get',
    params
  })
}

export function createArticle(data) {
  return request({
    url: '/articles',
    method: 'post',
    data
  })
}
