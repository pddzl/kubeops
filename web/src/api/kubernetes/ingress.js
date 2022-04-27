import service from '@/utils/request'

// 获取ingress list
export const getIngressList = (data) => {
  return service({
    url: '/ingress/getIngressList',
    method: 'post',
    data
  })
}

// 获取ingress in raw
export const getIngressRaw = (data) => {
  return service({
    url: '/ingress/getIngressRaw',
    method: 'post',
    data
  })
}

// 获取ingress detail
export const getIngressDetail = (data) => {
  return service({
    url: '/ingress/getIngressDetail',
    method: 'post',
    data
  })
}
