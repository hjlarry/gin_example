import request from '@/utils/request'

export function getList(params) {
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
