import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/articles',
    method: 'get',
    params
  })
}

export function createArticle(data) {
  console.log(data)
  return request({
    url: '/articles',
    method: 'post',
    data
  })
}
