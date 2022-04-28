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

// 获取节点信息 in raw
export const getNodeRaw = (data) => {
  return service({
    url: '/node/getNodeRaw',
    method: 'post',
    data
  })
}

// 获取node pods
export const getNodePods = (data) => {
  return service({
    url: '/node/getNodePods',
    method: 'post',
    data
  })
}
