import service from '@/utils/request'

// 获取节点树
export const getNodeList = (data) => {
  return service({
    url: '/node/getNodeList',
    method: 'post',
    data
  })
}

// 获取节点详情
export const getNodeDetail = (data) => {
  return service({
    url: '/node/getNodeDetail',
    method: 'post',
    data
  })
}
